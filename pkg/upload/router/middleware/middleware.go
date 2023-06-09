package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/internal/constant"
	"github.com/trungluongwww/goupload/internal/response"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

func generateDatetimeName(name string) string {
	return fmt.Sprintf("%s_%s", time.Now().Format("2006_01_02_15_04_05_0700"), name)
}

func UploadPhoto() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			form, err := c.MultipartForm()
			if err != nil || form == nil {
				return response.R400(c, nil, err.Error())
			}

			files := form.File["files"]
			if files == nil || len(files) == 0 {
				return response.R400(c, nil, response.CommonFileNotFound)
			}

			var (
				wg       = sync.WaitGroup{}
				payloads = make([]requestmodel.FileInfoPayload, 0)
			)

			wg.Add(len(files))

			for _, file := range files {
				rename := generateDatetimeName(file.Filename)
				f, err := file.Open()
				if err != nil {
					return response.R400(c, nil, err.Error())
				}

				defer f.Close()

				dir, _ := os.Getwd()

				filePath := path.Join(dir, constant.UploadFolderPath, rename)

				dst, err := os.Create(filePath)
				if err != nil {
					return response.R400(c, nil, err.Error())
				}

				defer dst.Close()

				if _, err := io.Copy(dst, f); err != nil {
					return response.R400(c, nil, err.Error())
				}

				payloads = append(payloads, requestmodel.FileInfoPayload{
					OriginName: file.Filename,
					Name:       rename,
					Size:       file.Size,
					Ext:        strings.ReplaceAll(path.Ext(file.Filename), ".", ""),
					Path:       filePath,
				})
			}

			c.Set("payload", payloads)

			return next(c)
		}
	}
}
