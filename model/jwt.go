package model

import (
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	jwt.StandardClaims
}
