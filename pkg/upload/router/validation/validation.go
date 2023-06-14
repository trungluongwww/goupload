package routervalidation

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/trungluongwww/goupload/internal/response"
	"github.com/trungluongwww/goupload/internal/utils/echocontext"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
)

type ValidationInterface interface {
	ClientPayload(next echo.HandlerFunc) echo.HandlerFunc
}

type validationImpl struct {
}

func Validation() ValidationInterface {
	return validationImpl{}
}

func (validationImpl) ClientPayload(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload requestmodel.ClientPayload
		)

		strPayload := c.FormValue("payload")

		if err := json.Unmarshal([]byte(strPayload), &payload); err != nil {

			return response.R400(c, nil, err.Error())
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}
