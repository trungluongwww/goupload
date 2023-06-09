package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/internal/response"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	"github.com/trungluongwww/goupload/pkg/upload/service"
)

type Handler interface {
	Photo(c echo.Context) error
}

type handler struct {
}

func Init() Handler {
	return handler{}
}

func (handler) Photo(c echo.Context) error {
	var (
		ctx     = c.Request().Context()
		payload = c.Get("payload").([]requestmodel.FileInfoPayload)
		s       = service.Photo()
	)

	res, err := s.Upload(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, res, "")
}
