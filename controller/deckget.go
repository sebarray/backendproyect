package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/backendproyect/db/mysql/deckdb"
	"github.com/sebarray/backendproyect/service"
)

func GetDeck(ctx echo.Context) error {

	claims := service.GetJwtClaims(ctx)
	def := ctx.QueryParam("maso")
	if claims.Id == "" {
		return echo.ErrUnauthorized
	}
	typedb := deckdb.GetProvider(os.Getenv("TYPE_DB"))
	Cards, err := typedb.ReadDeck(def, claims.Id)

	if err != nil {
		ctx.Error(err)
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}
	return ctx.JSON(http.StatusOK, Cards)
}
