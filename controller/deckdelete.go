package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/backendproyect/db/mysql/deckdb"
	"github.com/sebarray/backendproyect/model"
	"github.com/sebarray/backendproyect/service"
)

func DeleteDeck(ctx echo.Context) error {
	claims := service.GetJwtClaims(ctx)
	if claims.Id == "" {
		return echo.ErrUnauthorized
	}
	var card model.Card
	err := ctx.Bind(&card)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}

	typedb := deckdb.GetProvider(os.Getenv("TYPE_DB"))
	err = typedb.DeleteDeck(card.Id)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}

	return ctx.JSON(http.StatusAccepted, echo.Map{"error": ""})

}
