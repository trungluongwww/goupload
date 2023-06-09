package initfolder

import (
	"os"
	"path"
)

func Init() {
	dir, _ := os.Getwd()
	p := path.Join(dir, "uploads")

	if _, err := os.Stat(p); os.IsNotExist(err) {
		_ = os.MkdirAll(p, os.ModePerm)
	}
}
