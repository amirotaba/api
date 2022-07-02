package httpd

import (
	"github.com/labstack/echo/v4"
	"login/internal/domain"
	"net/http"
)

func Deliver() {
	e := echo.New()

	e.File("/","pkg/HTML/index.html")

	e.POST("/signin", Signin)

	e.GET("/signup", Signup)
	e.GET("/user/account/:username", Account)

	e.Logger.Fatal(e.Start(":4000"))
}

func Account(c echo.Context) error {
	u := &domain.User{
	c.Param("username"),
	//Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

func Signin(c echo.Context) error {
	var loginform domain.LoginForm
	if err := c.JSON(http.StatusOK, loginform); err != nil {
		return c.JSON(200, "not ok")
	}else {
		return  c.JSON(http.StatusOK, "ok")
	}
	return c.JSON(200, "nothing")

}

//func Signup(c echo.Context) error {
//	user := new(domain.User)
//	if err := c.Bind(user); user != nil {
//		return err
//	}
//	if user.Name == "" || user.Password == "" {
//		return c.JSON(http.StatusConflict, "invalid name or password")
//		}
//	}
//	if u := model.
//
//}