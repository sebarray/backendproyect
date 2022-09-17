package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sebarray/backendproyect/config"
	"github.com/sebarray/backendproyect/controller"
)

func RouteCard(e *echo.Group) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}), middleware.JWTWithConfig(config.Jwt()))
	e.POST("", controller.PostCard)
	e.GET("", controller.GetCard)
	e.PUT("", controller.PutCard)
	e.DELETE("", controller.DeleteCard)
}
