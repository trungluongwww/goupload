package compresspdf

import (
	"fmt"
	"os/exec"
)

func Run(input, output string) (err error) {
	cmd := exec.Command("gs", buildCompressionArguments(input, output)...)

	if _, err = cmd.Output(); err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func buildCompressionArguments(input, output string) []string {
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
