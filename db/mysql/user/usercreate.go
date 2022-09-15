package userdb

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sebarray/backendproyect/db/mysql"
	"github.com/sebarray/backendproyect/model"
)

func (u Userdb) CreateUser(user model.User) (model.User, error) {
	connection := mysql.ConnectionDB()
	defer connection.Close()
	if user.Email == "" || user.Password == "" {

		return user, fmt.Errorf("ocurrio un erro revise un request")

	}
	user.Id = uuid.New().String()

	insert, err := connection.Prepare(mysql.QueryCreateUser(user))
	if err != nil {
		return user, err
	}
	_, err = insert.Exec()
	if err != nil {
		return user, err
	}
	return user, nil

}
