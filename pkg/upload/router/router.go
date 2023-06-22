package router

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/pkg/upload/handler"
	"github.com/trungluongwww/goupload/pkg/upload/router/middleware"
	routervalidation "github.com/trungluongwww/goupload/pkg/upload/router/validation"
)

func Init(e *echo.Echo) {
	var (
		g = e.Group("")
		h = handler.Init()
		v = routervalidation.Validation()
	)

	g.POST("/photos", h.Photo, v.ClientPayload, middleware.UploadPhotos())

	g.POST("/zip-photos", h.Photo, v.ClientPayload, middleware.UploadZipPhoto())

	g.POST("/pdfs", h.PDF, v.ClientPayload, middleware.UploadFile())

}
