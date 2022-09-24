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
func QueryCreateDeck(deck model.Deck) string {
	query := "INSERT INTO deck (id, name, image, iduser) VALUES "
	query += fmt.Sprintf("('%s', '%s','%s',  '%s'); ",
		deck.Id, deck.Name, deck.Image, deck.IdUser)
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
func QueryReaDeck(where string) string {
	return "select * from deck  " + where
}

func QueryCardUpdate(card model.Card) string {
	query := "UPDATE card SET "
	query += fmt.Sprintf("title='%s',description='%s',image='%s'", card.Title, card.Description, card.Image)
	query += fmt.Sprintf(" where id = '%s'", card.Id)
	return query

}

func QueryDeckUpdate(deck model.Deck) string {
	query := "UPDATE card SET "
	query += fmt.Sprintf("name='%s',image='%s'", deck.Name, deck.Image)
	query += fmt.Sprintf(" where id = '%s'", deck.Id)
	return query

}
func QueryInsertDeckCard(id, idcard, iddeck, iduser string) string {
	query := "INSERT INTO deckcard (id, idcard, iddeck,iduser)VALUES"
	query += fmt.Sprintf("'%s','%s', '%s', '%s');", id, idcard, iddeck, iduser)
	return query
}

func QueryReadDeckCard(iduser, iddeck string) string {
	query := "SELECT idcard, card.title, card.description, card.image, card.iduser FROM deckcard JOIN card ON deckcard.idcard=card.id JOIN deck ON deckcard.iddeck= deck.id"
	query += fmt.Sprintf("deckcard.iduser= '%s' AND  deckcard.iddeck= '%s';", iduser, iddeck)
	return query
}
