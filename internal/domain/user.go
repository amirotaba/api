package domain

import (
	"github.com/golang-jwt/jwt"
	"login/internal/user/database"
)

var SigningKey = []byte("secret")

type LoginForm struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type JwtCustom struct {
	UID  int    `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type Message struct {
	Msg   string
	UserInfo    database.User
	Token string

}