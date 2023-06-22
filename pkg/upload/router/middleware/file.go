package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/internal/constant"
	"github.com/trungluongwww/goupload/internal/response"
	"github.com/trungluongwww/goupload/internal/utils/echocontext"
	"github.com/trungluongwww/goupload/internal/utils/prandom"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	"io"
	"os"
	"path"
	"strings"
)

func UploadFile() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			file, err := c.FormFile("file")
			if err != nil || file == nil {
				return response.R400(c, nil, response.CommonFileNotFound)
			}

			src, err := file.Open()
			if err != nil {
				return response.R400(c, nil, err.Error())
			}

			defer src.Close()

			var (
				dir, _   = os.Getwd()
				ext      = strings.ReplaceAll(path.Ext(file.Filename), ".", "")
				rename   = prandom.RandomNameFileFromExtension(ext)
				pathFile = path.Join(dir, constant.UploadFolderPath, rename)
			)

			dst, err := os.Create(pathFile)
			if err != nil {
				return response.R400(c, nil, err.Error())
			}
			defer dst.Close()

			// copy to folder upload
			if _, err = io.Copy(dst, src); err != nil {
				return response.R400(c, nil, err.Error())
			}

			infos := requestmodel.FileInfoPayload{
				OriginName: file.Filename,
				Path:       pathFile,
				Name:       rename,
				Size:       file.Size,
				Ext:        ext,
				Type:       constant.TypeFile.File,
			}

			echocontext.SetSingleFile(c, infos)

			return next(c)
		}
	}
}
