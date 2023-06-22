package service

import (
	"context"
	"fmt"
	"github.com/trungluongwww/goupload/internal/constant"
	"github.com/trungluongwww/goupload/internal/model"
	"github.com/trungluongwww/goupload/internal/utils/pfile"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	responsemodel "github.com/trungluongwww/goupload/pkg/upload/model/response"
	"github.com/trungluongwww/goupload/pkg/upload/module/resizeimage"
	"github.com/trungluongwww/goupload/pkg/upload/module/s3"
	"math"
	"os"
	"path"
	"sync"
	"time"
)

type PhotoInterface interface {
	Upload(c context.Context, files []requestmodel.FileInfoPayload, payload requestmodel.ClientPayload) ([]*responsemodel.PhotoResponse, error)
}

type photoImpl struct {
}

func Photo() PhotoInterface {
	return photoImpl{}
}

func (s photoImpl) Upload(ctx context.Context, files []requestmodel.FileInfoPayload, payload requestmodel.ClientPayload) (result []*responsemodel.PhotoResponse, err error) {
	var (
		//bucketOpt = modules3.GetBucketOption(payload.BucketName)
		wg sync.WaitGroup
	)

	wg.Add(len(files))
	chanErr := make(chan error, len(files))
	result = make([]*responsemodel.PhotoResponse, len(files))

	for i, file := range files {
		go func(index int, f requestmodel.FileInfoPayload) {
			defer wg.Done()
			var (
				err error
			)

			result[index], err = s.processingPhoto(ctx, f, payload)
			chanErr <- err
		}(i, file)
	}
	wg.Wait()
	close(chanErr)

	for e := range chanErr {
		if e != nil {
			err = e
			go s.Rollback(result)
			return
		}
	}

	return
}
func (s photoImpl) Rollback(result []*responsemodel.PhotoResponse) {
	var (
		dir, _ = os.Getwd()
	)

	for _, item := range result {
		if item != nil {
			sm := fmt.Sprintf("%s_%s", constant.PrefixDimension.Small, item.Name)
			md := fmt.Sprintf("%s_%s", constant.PrefixDimension.Medium, item.Name)

			// TODO: remove all file stored in s3
			os.Remove(path.Join(dir, "static", sm))
			fmt.Println("remove sm", path.Join(dir, "static", sm))
			os.Remove(path.Join(dir, "static", md))
			fmt.Println("remove md", path.Join(dir, "static", md))
		}
	}
}

func (s photoImpl) convertToRawAndInsertDoc(ctx context.Context, origin requestmodel.FileInfoPayload, small requestmodel.FileInfoPayload, medium requestmodel.FileInfoPayload) *model.FileRaw {
	// TODO insert db
	return &model.FileRaw{
		ID:           model.NewAppID(),
		Name:         origin.Name,
		Ext:          origin.Ext,
		OriginalName: origin.OriginName,
		Dimension: model.DimensionFilePhoto{
			Original: model.FileSize{
				Width:  origin.Width,
				Height: origin.Height,
			},
			Small: model.FileSize{
				Width:  small.Width,
				Height: small.Height,
			},
			Medium: model.FileSize{
				Width:  medium.Width,
				Height: medium.Height,
			},
		},
		CreatedAt: time.Now(),
		Type:      origin.Type,
		UpdatedAt: time.Now(),
	}
}

func (s photoImpl) convertToResponse(raw model.FileRaw) *responsemodel.PhotoResponse {
	var (
		// TODO: fix
		url = os.Getenv("HOST")
	)

	return &responsemodel.PhotoResponse{
		ID:   raw.ID.Hex(),
		Name: raw.Name,
		Dimensions: responsemodel.DimensionPhoto{
			Small: responsemodel.SizePhoto{
				Width:  raw.Dimension.Small.Width,
				Height: raw.Dimension.Small.Height,
				Url:    fmt.Sprintf("%s/%s/%s_%s", url, "static", constant.PrefixDimension.Small, raw.Name),
			},
			Medium: responsemodel.SizePhoto{
				Width:  raw.Dimension.Medium.Width,
				Height: raw.Dimension.Medium.Height,
				Url:    fmt.Sprintf("%s/%s/%s_%s", url, "static", constant.PrefixDimension.Medium, raw.Name),
			},
		},
	}
}

// processingPhoto ...
func (s photoImpl) processingPhoto(ctx context.Context, f requestmodel.FileInfoPayload, payload requestmodel.ClientPayload) (result *responsemodel.PhotoResponse, err error) {
	var (
		bucketOpt = s3.GetBucketOption(payload.BucketName)
	)
	defer os.Remove(f.Path)

	if err = f.ValidateSize(bucketOpt.MaxSizePhoto); err != nil {
		return
	}

	if err = f.ValidateExtensionPhoto(); err != nil {
		return
	}

	small, medium, err := s.resizePhotos(f, payload.Resize)

	if err != nil {
		return
	}

	// TODO: store s3 and remove file

	fileRaw := s.convertToRawAndInsertDoc(ctx, f, small, medium)

	result = s.convertToResponse(*fileRaw)
	return
}

// resizePhoto ...
func (s photoImpl) resizePhotos(f requestmodel.FileInfoPayload, option string) (sm requestmodel.FileInfoPayload, md requestmodel.FileInfoPayload, err error) {
	var (
		width, height, image = pfile.GetImageDimension(f.Path)
		resizeDimension      = s.getResizeDimension(option)
	)

	f.Width = width
	f.Height = height

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		md = s.getDimension(f, constant.PrefixDimension.Medium, resizeDimension)
		err = resizeimage.ProcessResizeImage(md, image)
	}()

	go func() {
		defer wg.Done()
		sm = s.getDimension(f, constant.PrefixDimension.Small, resizeDimension)
		err = resizeimage.ProcessResizeImage(sm, image)
	}()

	wg.Wait()

	return
}

func (s photoImpl) getDimension(origin requestmodel.FileInfoPayload, prefixSize string, option model.ResizeDimension) (result requestmodel.FileInfoPayload) {
	var (
		maxWidth  int
		maxHeight int
		dir, _    = os.Getwd()
		rename    = fmt.Sprintf("%s_%s", prefixSize, origin.Name)
	)

	result = requestmodel.FileInfoPayload{
		Name:       rename,
		OriginName: origin.OriginName,
		// TODO : change folder storage
		Path:   path.Join(dir, "static", rename),
		Ext:    origin.Ext,
		Type:   origin.Type,
		Width:  origin.Width,
		Height: origin.Height,
		Size:   origin.Size,
	}

	switch prefixSize {
	case constant.PrefixDimension.Small:
		maxWidth = option.SmallWidth
		maxHeight = option.SmallHeight
	case constant.PrefixDimension.Medium:
		maxWidth = option.MediumWidth
		maxHeight = option.MediumHeight
	}

	if origin.Width <= maxWidth && origin.Height <= maxHeight {
		return
	}

	rw := float64(maxWidth) / float64(origin.Width)
	rh := float64(maxHeight) / float64(origin.Height)

	if rw > rh {
		result.Width = int(math.Round(float64(origin.Width) * rh))
		result.Height = maxHeight
	} else {
		result.Width = maxWidth
		result.Height = int(math.Round(float64(origin.Height) * rw))
	}
	return
}

func (s photoImpl) getResizeDimension(name string) model.ResizeDimension {
	switch name {
	case constant.ResizeName.Size200x200:
		return model.ResizeDimension{
			MediumWidth:  200,
			MediumHeight: 200,
			SmallWidth:   100,
			SmallHeight:  100,
		}
	case constant.ResizeName.Size1280x1124:
		return model.ResizeDimension{
			MediumWidth:  1280,
			MediumHeight: 1124,
			SmallWidth:   640,
			SmallHeight:  562,
		}
	case constant.ResizeName.Size430x600:
		return model.ResizeDimension{
			MediumWidth:  430,
			MediumHeight: 600,
			SmallWidth:   215,
			SmallHeight:  300,
		}
	case constant.ResizeName.Size1280x1280:
		return model.ResizeDimension{
			MediumWidth:  1280,
			MediumHeight: 1280,
			SmallWidth:   640,
			SmallHeight:  640,
		}
	default:
		return model.ResizeDimension{
			MediumWidth:  720,
			MediumHeight: 720,
			SmallWidth:   480,
			SmallHeight:  480,
		}
	}
}
