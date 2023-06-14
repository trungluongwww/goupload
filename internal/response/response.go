package response

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
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

func RouteValidation(c echo.Context, err error) error {
	key := getMessage(err)

	// Return
	return R400(c, nil, key)
}

func getMessage(err error) string {
	err1, ok := err.(validation.Errors)
	if !ok {
		err2, ok := err.(validation.ErrorObject)
		if ok {
			return err2.Message()
		}
		return err.Error()
	}
	for _, item := range err1 {
		if item == nil {
			continue
		}
		return getMessage(item)
	}
	return err.Error()
}

func R200(c echo.Context, data interface{}, key string) error {
	code := getCodeByKey(key)

	if code.Code == -1 {
		code.Message = CommonSuccess
	}
	return sendResponse(c, http.StatusOK, data, true, key, code.Code)
}

func R400(c echo.Context, data interface{}, key string) error {
	code := getCodeByKey(key)

	if code.Code == -1 {
		code.Message = key
	}
	return sendResponse(c, http.StatusBadRequest, data, false, code.Message, code.Code)
}

func R401(c echo.Context, data interface{}, key string) error {
	code := getCodeByKey(key)

	if code.Code == -1 {
		code.Message = key
	}
	return sendResponse(c, http.StatusUnauthorized, data, true, code.Message, code.Code)
}
