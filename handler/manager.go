package hanlder

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/backendproyect/router"
)

func Manager() {
	e := echo.New()

	cardRouter := e.Group("/card")
	userRouter := e.Group("/user")
	deckRouter := e.Group("/deck")

	router.RouteDeck(deckRouter)
	router.RouteCard(cardRouter)
	router.RouteUser(userRouter)
	router.RouteLogin(e)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8086"

	}

	e.Start(":" + PORT)
}
