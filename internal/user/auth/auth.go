package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"login/internal/domain"
	"login/internal/user/database"
	"net/http"
	"time"
)

func SignUp(c echo.Context) error {
	user := new(database.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if user.Username == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid username or password",
		}
	}
	db := database.DbConn()
	user.CreateUser(db)
	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

func SignIn(c echo.Context) error {
	user := new(database.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if user.Username == "" || user.Password == "" {
		return &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: "invalid username or password",
		}
	}
	db := database.DbConn()
	u := user.Find(db)
	if u.ID == 0 || u.Password != user.Password {
		return &echo.HTTPError{
			Code: http.StatusUnauthorized,
			Message: "invalid username or password",
		}
	}
	claims := &domain.JwtCustom{
		UID:  u.ID,
		Name: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(domain.SigningKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}