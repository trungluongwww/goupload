package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/internal/response"
	"github.com/trungluongwww/goupload/internal/utils/echocontext"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	"github.com/trungluongwww/goupload/pkg/upload/service"
)

type Handler interface {
	Photo(c echo.Context) error
	PDF(c echo.Context) error
}

type handler struct {
}

func Init() Handler {
	return handler{}
}

func (handler) Photo(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		files   = echocontext.GetFiles(c).([]requestmodel.FileInfoPayload)
		payload = echocontext.GetPayload(c).(requestmodel.ClientPayload)
		s       = service.Photo()
	)

	res, err := s.Upload(ctx, files, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}

func (handler) PDF(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		files   = echocontext.GetSingleFile(c).(requestmodel.FileInfoPayload)
		payload = echocontext.GetPayload(c).(requestmodel.ClientPayload)
		s       = service.File()
	)

	res, err := s.UploadCompressionPDF(ctx, files, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}
