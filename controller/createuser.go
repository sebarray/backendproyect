package controller

import (
	"fmt"
	"log"
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
		log.Println("error parsing user model")
		return fmt.Errorf("error parsing user model")
	}
	typedb := userdb.GetProvider(os.Getenv("TYPE_DB"))
	user, err = typedb.CreateUser(user)
	if err != nil {
		log.Println("error creating user")
		err = fmt.Errorf("error creating user")
		http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	user.Password = ""
	log.Println("user create id: ", user.Id)
	return c.JSON(http.StatusOK, user)

}
