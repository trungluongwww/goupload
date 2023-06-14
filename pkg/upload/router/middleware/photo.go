package middleware

import (
	"archive/zip"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/internal/constant"
	"github.com/trungluongwww/goupload/internal/response"
	"github.com/trungluongwww/goupload/internal/utils/echocontext"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func generateDatetimeName(name string) string {
	return fmt.Sprintf("%s_%s", time.Now().Format("2006_01_02_15_04_05_0700"), name)
}

func UploadPhoto() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			form, err := c.MultipartForm()
			if err != nil {
				return response.R400(c, nil, err.Error())
			}

			files := form.File["files"]
			if files == nil || len(files) == 0 {
				return response.R400(c, nil, response.CommonFileNotFound)
			}

			var (
				infos        = make([]requestmodel.FileInfoPayload, 0)
				processError error
			)

			for _, file := range files {
				rename := generateDatetimeName(file.Filename)
				f, err := file.Open()
				if err != nil {
					processError = err
					break
				}

				defer f.Close()

				dir, _ := os.Getwd()

				filePath := path.Join(dir, constant.UploadFolderPath, rename)

				dst, err := os.Create(filePath)
				if err != nil {
					processError = err
					break
				}

				defer dst.Close()

				if _, err := io.Copy(dst, f); err != nil {
					processError = err
					break
				}

				p := requestmodel.FileInfoPayload{
					OriginName: file.Filename,
					Name:       rename,
					Size:       file.Size,
					Ext:        strings.ReplaceAll(path.Ext(file.Filename), ".", ""),
					Path:       filePath,
					Type:       constant.TypeFile.Photo,
				}

				infos = append(infos, p)
			}

			// response if has a error
			if processError != nil {
				for _, p := range infos {
					os.Remove(p.Path)
				}

				return response.R400(c, nil, processError.Error())
			}

			echocontext.SetFiles(c, infos)
			return next(c)
		}
	}
}

func UploadZipPhoto() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			file, err := c.FormFile("file")
			if err != nil || file == nil {
				return response.R400(c, nil, response.CommonFileNotFound)
			}

			if ext := path.Ext(file.Filename); ext != constant.ZipExtension {
				return response.R400(c, nil, response.CommonInvalidExtension)
			}

			src, err := file.Open()
			if err != nil {
				return response.R400(c, nil, err.Error())
			}

			defer src.Close()

			var (
				dir, _ = os.Getwd()
			)

			zipPath := path.Join(dir, constant.UploadFolderPath, generateDatetimeName(file.Filename))

			dst, err := os.Create(zipPath)
			if err != nil {
				return response.R400(c, nil, err.Error())
			}
			defer dst.Close()

			// copy to folder upload
			if _, err = io.Copy(dst, src); err != nil {
				return response.R400(c, nil, err.Error())
			}

			defer os.Remove(zipPath)

			archive, err := zip.OpenReader(zipPath)
			if err != nil {
				return response.R400(c, nil, err.Error())
			}

			defer archive.Close()

			var (
				infos        = make([]requestmodel.FileInfoPayload, 0)
				processError error
			)

			for _, unzipFile := range archive.File {
				f, err := unzipFile.Open()
				if err != nil {
					processError = err
					break
				}

				defer f.Close()

				rename := generateDatetimeName(unzipFile.FileInfo().Name())
				filePath := path.Join(dir, constant.UploadFolderPath, rename)

				dst2, err := os.Create(filePath)
				if err != nil {
					processError = err
					break
				}

				defer dst2.Close()

				if _, err := io.Copy(dst2, f); err != nil {
					processError = err
					break
				}

				p := requestmodel.FileInfoPayload{
					OriginName: unzipFile.FileInfo().Name(),
					Name:       rename,
					Size:       unzipFile.FileInfo().Size(),
					Ext:        strings.ReplaceAll(path.Ext(unzipFile.FileInfo().Name()), ".", ""),
					Path:       filePath,
					Type:       constant.TypeFile.Photo,
				}

				infos = append(infos, p)
			}

			// response failed if list has an error
			if processError != nil {
				for _, p := range infos {
					os.Remove(p.Path)
				}

				return response.R400(c, nil, processError.Error())
			}

			echocontext.SetFiles(c, infos)
			return next(c)
		}
	}
}
