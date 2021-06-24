package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sayuen0/echo-book-api/lib/presentation/api"
	validator "gopkg.in/go-playground/validator.v9"
)

func main() {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.Use(api.HandleCustomContext)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users/:id", api.C(api.GetUser))
	e.POST("/users", api.C(api.CreateUser))
	log.Fatal(e.Start(":8888"))
}

// validator
type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
