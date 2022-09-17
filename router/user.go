package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sebarray/backendproyect/controller"
)

func RouteUser(e *echo.Group) {

	e.POST("", controller.CreateUser)
}
