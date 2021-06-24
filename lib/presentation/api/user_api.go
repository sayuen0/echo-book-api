package api

import (
	"github.com/sayuen0/echo-book-api/lib/domain/model/request_model"
	"log"
	"net/http"
)

func GetUser(c *CustomContext) error {
	u := new(request_model.UserGetRequest)
	if err := c.BindValidate(u); err != nil {
		return err
	}
	// TODO: useCase以下に渡してみる
	log.Printf("%+v", u)
	return c.JSON(http.StatusOK, u)
}
