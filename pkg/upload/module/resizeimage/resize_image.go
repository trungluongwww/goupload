package resizeimage

import (
	"fmt"
	"github.com/disintegration/imaging"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	"image"
)

func ProcessResizeImage(originPath string, data requestmodel.FileInfoPayload, src image.Image) error {
	dstImage := imaging.Resize(src, data.Width, 0, imaging.CatmullRom)
	//dstImage := resize.Resize(uint(data.Width), 0, src, resize.Lanczos2)
	// brightness upto 5%
	//adjustImg := imaging.AdjustFunc(dstImage, func(c color.NRGBA) color.NRGBA {
	//	brightness := (int(c.R) + int(c.G) + int(c.B)) / 3
	//
	//	if brightness > 128 {
	//		rgba := color.RGBA{
	//			R: c.R,
	//			G: c.G,
	//			B: c.B,
	//			A: c.A,
	//		}
	//
	//		img := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	//		img.Set(0, 0, rgba)
	//
	//		adjustedImg := imaging.AdjustContrast(img, 5)
	//
	//		adjustedRGBA := adjustedImg.At(0, 0).(color.NRGBA)
	//
	//		return color.NRGBA{
	//			R: adjustedRGBA.R,
	//			G: adjustedRGBA.G,
	//			B: adjustedRGBA.B,
	//			A: adjustedRGBA.A,
	//		}
	//	}
	//
	//	return c
	//})

	err := imaging.Save(dstImage, data.Path)
	if err != nil {
		fmt.Println("failed to save image:", err)
	}

	return err
}
