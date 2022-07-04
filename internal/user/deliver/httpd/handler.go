package httpd

import (
	"github.com/labstack/echo/v4"
	"login/internal/user/auth"
)

func Deliver() {
	e := echo.New()

	e.POST("/signin", auth.SignIn)
	e.POST("/signup", auth.SignUp)

	e.Logger.Fatal(e.Start(":4000"))
}