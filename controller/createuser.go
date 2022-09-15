package controller

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	userdb "github.com/sebarray/backendproyect/db/mysql/user"
	"github.com/sebarray/backendproyect/model"
)

func CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
	}
	typedb := userdb.GetProvider(os.Getenv("TYPE_DB"))
	user, err = typedb.CreateUser(user)
	if err != nil {
		http.Error(c.Response().Writer, err.Error(), http.StatusConflict)
	}
	user.Password = ""
	return c.JSON(http.StatusOK, user)

}
