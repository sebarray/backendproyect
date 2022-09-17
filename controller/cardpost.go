package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/backendproyect/db/mysql/cardb"
	"github.com/sebarray/backendproyect/model"
	"github.com/sebarray/backendproyect/service"
)

func PostCard(ctx echo.Context) error {
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
	card.IdUser = claims.Id
	typedb := cardb.GetProvider(os.Getenv("TYPE_DB"))
	err = typedb.CreateCard(card)
	if err != nil {
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{"error": ""})
}
