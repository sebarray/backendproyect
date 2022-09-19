package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/backendproyect/db/mysql/deckdb"
	"github.com/sebarray/backendproyect/model"
	"github.com/sebarray/backendproyect/service"
)

func PostDeck(ctx echo.Context) error {
	claims := service.GetJwtClaims(ctx)
	if claims.Id == "" {
		return echo.ErrUnauthorized

	}
	var deck model.Deck

	err := ctx.Bind(&deck)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}
	deck.IdUser = claims.Id
	typedb := deckdb.GetProvider(os.Getenv("TYPE_DB"))
	err = typedb.CreateDeck(deck)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{"error": ""})
}
