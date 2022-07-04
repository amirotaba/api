package httpd

import (
	"github.com/labstack/echo/v4"
	"login/internal/user/auth"
)

func Deliver() {
	e := echo.New()

	e.File("/", "pkg/HTML/index.html")

	e.POST("/signin", auth.SignIn)

	e.POST("/signup", auth.SignUp)
	//e.GET("/user/account/:username", Account)

	e.Logger.Fatal(e.Start(":4000"))
}