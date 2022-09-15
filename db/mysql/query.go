package mysql

import (
	"fmt"

	"github.com/sebarray/backendproyect/model"
	"github.com/sebarray/backendproyect/service"
)

func QueryCreateUser(user model.User) string {
	user.Password, _ = service.Encryptpswd(user.Password)
	query := "INSERT INTO user (id, name, email, password) VALUES "
	query += fmt.Sprintf("('%s', '%s','%s',  '%s'); ",
		user.Id, user.Name, user.Email, user.Password)
	return query

}
