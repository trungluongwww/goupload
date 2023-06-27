package resizeimage

import (
	"fmt"
	"github.com/disintegration/imaging"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	"image"
)

func ProcessResizeImage(data requestmodel.FileInfoPayload, src image.Image) error {
	dstImage := imaging.Resize(src, data.Width, 0, imaging.CatmullRom)

	err := imaging.Save(dstImage, data.Path, imaging.JPEGQuality(90))
	if err != nil {
		fmt.Println("failed to save image:", err)
	}

	return err
}
