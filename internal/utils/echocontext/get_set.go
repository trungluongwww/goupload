package echocontext

import "github.com/labstack/echo/v4"

func SetPayload(c echo.Context, value interface{}) {
	c.Set(keyPayload, value)
}

func GetPayload(c echo.Context) interface{} {
	return c.Get(keyPayload)
}

func SetFiles(c echo.Context, value interface{}) {
	c.Set(keyFiles, value)
}

func GetFiles(c echo.Context) interface{} {
	return c.Get(keyFiles)
}
