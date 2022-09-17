package config

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
	"github.com/sebarray/backendproyect/model"
)

func Jwt() middleware.JWTConfig {
	cofingJwt := middleware.JWTConfig{
		Claims:     &model.JwtClaims{},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}
	return cofingJwt
}
