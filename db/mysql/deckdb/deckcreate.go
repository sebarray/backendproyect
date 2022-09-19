package deckdb

import (
	"github.com/google/uuid"
	"github.com/sebarray/backendproyect/db/mysql"
	"github.com/sebarray/backendproyect/model"
)

func (d DeckDB) CreateDeck(deck model.Deck) error {
	connection := mysql.ConnectionDB()
	defer connection.Close()
	deck.Id = uuid.New().String()
	insert, err := connection.Prepare(mysql.QueryCreateDeck(deck))
	if err != nil {
		return err
	}
	_, err = insert.Exec()
	if err != nil {
		return err
	}
	return err

}
