package auth

import (
	"github.com/labstack/echo/v4"
	"login/internal/user/database"
	"net/http"
)

func SignUp(c echo.Context) error {
	user := new(database.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if user.Username == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}
	database.CreateUser(user)
	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}