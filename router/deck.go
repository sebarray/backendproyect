package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sebarray/backendproyect/config"
	"github.com/sebarray/backendproyect/controller"
)

func RouteDeck(e *echo.Group) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}), middleware.JWTWithConfig(config.Jwt()))
	e.POST("", controller.PostCard)
	e.GET("", controller.GetDeck)
	e.PUT("", controller.PutDeck)
	e.DELETE("", controller.DeleteDeck)
}
