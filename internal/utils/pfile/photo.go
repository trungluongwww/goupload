package pfile

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

// GetImageDimension ...
func GetImageDimension(imagePath string) (int, int, image.Image) {
	file, err := os.Open(imagePath)
	defer file.Close()
	rdr := bufio.NewReader(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.Decode(rdr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
		return 0, 0, nil
	}

	b := image.Bounds()
	return b.Max.X, b.Max.Y, image
}
