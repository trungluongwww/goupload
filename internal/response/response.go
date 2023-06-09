package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func sendResponse(c echo.Context, httpCode int, data interface{}, success bool, message string, code int) error {
	if data == nil {
		data = echo.Map{}
	}

	return c.JSON(httpCode, echo.Map{
		"data":    data,
		"success": success,
		"message": message,
		"code":    code,
	})
}

func R200(c echo.Context, data interface{}, key string) error {
	code := getCodeByKey(key)

	if code.Code == -1 {
		code.Message = CommonSuccess
	}
	return sendResponse(c, http.StatusOK, data, true, code.Message, code.Code)
}

func R400(c echo.Context, data interface{}, key string) error {
	code := getCodeByKey(key)

	if code.Code == -1 {
		code.Message = CommonBadRequest
	}
	return sendResponse(c, http.StatusBadRequest, data, false, code.Message, code.Code)
}

func R401(c echo.Context, data interface{}, key string) error {
	code := getCodeByKey(key)

	if code.Code == -1 {
		code.Message = CommonUnAuthorization
	}
	return sendResponse(c, http.StatusUnauthorized, data, true, code.Message, code.Code)
}
