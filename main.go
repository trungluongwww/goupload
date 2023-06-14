package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/trungluongwww/goupload/pkg/upload"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))

	upload.Server(e)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":5000"))
}
