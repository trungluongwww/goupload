package main

import (
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/pkg/upload"
)

func main() {
	e := echo.New()

	upload.Server(e)

	e.Logger.Fatal(e.Start(":5000"))
}
