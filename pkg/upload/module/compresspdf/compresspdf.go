package compresspdf

import (
	"errors"
	"fmt"
	"github.com/trungluongwww/goupload/internal/response"
	"os/exec"
)

type PDF struct {
	Input  string
	Output string
}

func (c PDF) Compress() (err error) {
	if c.Input == "" {
		return errors.New(response.CommonFileNotFound)
	}

	if c.Output == "" {
		c.Output = c.Input
	}

	cmd := exec.Command("gs", c.buildCompressionArguments(c.Input, c.Output)...)

	if _, err = cmd.Output(); err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func (c PDF) buildCompressionArguments(input, output string) []string {
	if output == "" {
		output = input
	}

	args := []string{
		"-sDEVICE=pdfwrite",
		"-dPDFSETTINGS=/screen",
		"-dNOPAUSE", "-dQUIET", "-dBATCH",
		"-dCompatibilityLevel=1.4",

		"-dSubsetFonts=true",
		"-dCompressFonts=true",
		"-dEmbedAllFonts=true",

		"-sProcessColorModel=DeviceRGB",
		"-sColorConversionStrategy=RGB",
		"-sColorConversionStrategyForImages=RGB",
		"-dConvertCMYKImagesToRGB=true",

		"-dDetectDuplicateImages=true",
		"-dColorImageDownsampleType=/Bicubic",
		"-dColorImageResolution=300",
		"-dGrayImageDownsampleType=/Bicubic",
		"-dGrayImageResolution=300",
		"-dMonoImageDownsampleType=/Bicubic",
		"-dMonoImageResolution=300",
		"-dDownsampleColorImages=true",

		"-dDoThumbnails=false",
		"-dOptimize=true",
		"-dCreateJobTicket=false",
		"-dPreserveEPSInfo=false",
		"-dPreserveOPIComments=false",
		"-dPreserveOverprintSettings=false",
		"-dUCRandBGInfo=/Remove",
		fmt.Sprintf("-sOutputFile=%s", output),
		input,
	}

	return args
}
