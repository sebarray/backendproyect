package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sebarray/backendproyect/model"
)

func CreateToken(user model.User) (string, error) {

	claims := &model.JwtClaims{
		Name: user.Name,
		Id:   user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return t, err
	}
	return t, nil
}
