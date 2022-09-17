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
func QueryCreateCard(card model.Card) string {
	query := "INSERT INTO card (id, title, description, image,iduser) VALUES "
	query += fmt.Sprintf("('%s', '%s','%s',  '%s', '%s'); ",
		card.Id, card.Title, card.Description, card.Image, card.IdUser)
	return query

}

func QueryReadCard(where string) string {
	return "select * from card  " + where
}

func QueryCardUpdate(card model.Card) string {
	query := "UPDATE card SET "
	query += fmt.Sprintf("title='%s',description='%s',image='%s'", card.Title, card.Description, card.Image)
	query += fmt.Sprintf(" where id = '%s'", card.Id)
	return query

}
