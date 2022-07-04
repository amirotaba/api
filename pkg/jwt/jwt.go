package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"login/internal/domain"
	"login/internal/user/database"
	"time"
)

func GetToken(user database.User) string {
	claims := &domain.JwtCustom{
		UID:  user.ID,
		Name: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(domain.SigningKey)
	if err != nil{
		fmt.Println(err)
	}
	//t = token code
	return t
}