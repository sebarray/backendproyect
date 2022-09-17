package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sebarray/backendproyect/db/mysql/cardb"
	"github.com/sebarray/backendproyect/service"
)

func GetCard(ctx echo.Context) error {

	claims := service.GetJwtClaims(ctx)
	def := ctx.QueryParam("maso")
	if claims.Id == "" {
		return echo.ErrUnauthorized
	}
	typedb := cardb.GetProvider(os.Getenv("TYPE_DB"))
	Cards, err := typedb.ReadCard(def, claims.Id)

	if err != nil {
		ctx.Error(err)
		http.Error(ctx.Response().Writer, err.Error(), http.StatusConflict)
	}
	return ctx.JSON(http.StatusOK, Cards)

}
