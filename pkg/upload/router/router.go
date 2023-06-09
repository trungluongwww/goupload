package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/pkg/upload/handler"
	"github.com/trungluongwww/goupload/pkg/upload/router/middleware"
)

func Init(e *echo.Echo) {
	var (
		g = e.Group("")
		h = handler.Init()
	)

	g.POST("/photos", h.Photo, middleware.UploadPhoto())
}
