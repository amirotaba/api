package auth

import (
	"github.com/labstack/echo/v4"
	"login/internal/domain"
	"login/internal/user/database"
	"login/pkg/jwt"
)

func SignUp(c echo.Context) error {
	user := new(database.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if user.Username == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    400,
			Message: "invalid username or password",
		}
	}
	db := database.DbConn()
	if u := user.Find(db); u.ID != 0{
		return &echo.HTTPError{
			Code:    409,
			Message: "this username isn't available",
		}
	}
	user.CreateUser(db)
	user.Password = ""
	return c.JSON(201, user)
}

func SignIn(c echo.Context) error {
	user := new(database.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if user.Username == "" || user.Password == "" {
		return &echo.HTTPError{
			Code: 400,
			Message: "invalid username or password",
		}
	}
	db := database.DbConn()
	u := user.Find(db)
	if u.ID == 0 || u.Password != user.Password {
		return &echo.HTTPError{
			Code: 401,
			Message: "invalid username or password",
		}
	}
	token := jwt.GetToken(u)

	m := &domain.Message{
		Msg:   "you logged in as, ",
		UserInfo: u,
		Token: token,
	}
	return c.JSON(200, m)
}