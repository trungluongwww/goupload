package upload

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/internal/response"
	"github.com/trungluongwww/goupload/internal/utils/initfolder"
	"github.com/trungluongwww/goupload/pkg/upload/router"
)

func Server(e *echo.Echo) {
	response.Init()

	initfolder.Init()

	router.Init(e)
}
