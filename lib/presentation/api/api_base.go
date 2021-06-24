package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type callFunc func(c *CustomContext) error

type CustomContext struct {
	echo.Context
}

func C(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*CustomContext))
	}
}

func HandleCustomContext(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(&CustomContext{c})
	}
}
func (c *CustomContext) BindValidate(i interface{}) error {
	if err := c.Bind(i); err != nil {
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	if err := c.Validate(i); err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
	}
	return nil
}
