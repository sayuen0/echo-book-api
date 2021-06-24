package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sayuen0/echo-book-api/lib/presentation/api"
	validator "gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)



func main() {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.Use(api.HandleCustomContext)

	e.GET("/users/:id", api.C(api.GetUser))

	e.POST("/users", api.C(func(c *api.CustomContext) error {
		u := new(UserCreateRequest)
		if err := c.BindValidate(u); err != nil {
			return err
		}
		log.Printf("%+v", u)
		return c.JSON(http.StatusCreated, u)
	}))
	e.Start(":8888")
}

// validator
type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

