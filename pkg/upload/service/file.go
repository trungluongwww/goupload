package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/trungluongwww/goupload/internal/constant"
	"github.com/trungluongwww/goupload/internal/model"
	"github.com/trungluongwww/goupload/internal/response"
	"github.com/trungluongwww/goupload/internal/utils/prandom"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	responsemodel "github.com/trungluongwww/goupload/pkg/upload/model/response"
	"github.com/trungluongwww/goupload/pkg/upload/module/compresspdf"
	"io"
	"os"
	"path"
	"time"
)

type FileInterface interface {
	UploadCompressionPDF(ctx context.Context, file requestmodel.FileInfoPayload, payload requestmodel.ClientPayload) (*responsemodel.FileResponse, error)
}

type fileImpl struct {
}

func File() FileInterface {
	return fileImpl{}
}

func (s fileImpl) UploadCompressionPDF(ctx context.Context, file requestmodel.FileInfoPayload, payload requestmodel.ClientPayload) (*responsemodel.FileResponse, error) {
	var (
		dir, _ = os.Getwd()
	)

	defer os.Remove(file.Path)

	// validate
	if file.Ext != constant.PDFExtension {
		return nil, errors.New(response.CommonInvalidExtension)
	}

	reName := prandom.RandomNameFileFromExtension(file.Ext)

	rePath := path.Join(dir, constant.UploadFolderPath, reName)
	defer os.Remove(rePath)

	// compress
	c := compresspdf.PDF{
		Input:  file.Path,
		Output: rePath,
	}

	if err := c.Compress(); err != nil {
		return nil, err
	}

	file.Path = rePath
	file.Name = reName

	// TODO: store s3
	dst, err := os.Create(path.Join(dir, "static", reName))
	defer dst.Close()
	if err != nil {
		return nil, err
	}

	src, err := os.Open(file.Path)
	defer src.Close()
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	raw := s.convertToRawAndInsertDoc(ctx, file)

	return s.convertToResponse(*raw), nil
}

func (s fileImpl) convertToRawAndInsertDoc(ctx context.Context, origin requestmodel.FileInfoPayload) *model.FileRaw {
	// TODO insert db
	return &model.FileRaw{
		ID:           model.NewAppID(),
		Name:         origin.Name,
		Ext:          origin.Ext,
		OriginalName: origin.OriginName,
		CreatedAt:    time.Now(),
		Type:         origin.Type,
		UpdatedAt:    time.Now(),
	}
}

func (s fileImpl) convertToResponse(raw model.FileRaw) *responsemodel.FileResponse {
	var (
		// TODO: fix
		host = os.Getenv("HOST")
	)

	return &responsemodel.FileResponse{
		ID:         raw.ID.Hex(),
		Name:       raw.Name,
		NameOrigin: raw.OriginalName,
		Url:        fmt.Sprintf("%s/%s/%s", host, "static", raw.Name),
	}
}
