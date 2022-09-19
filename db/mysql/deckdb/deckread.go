package deckdb

import (
	"fmt"

	"github.com/sebarray/backendproyect/db/mysql"
	"github.com/sebarray/backendproyect/model"
)

func (d DeckDB) ReadDeck(def, id string) ([]model.Deck, error) {

	condition := "where iduser = "
	if def == "default" {
		condition += fmt.Sprintf("'%s'", def)
	} else {
		condition += fmt.Sprintf("'%s'", id)
	}
	fmt.Println(mysql.QueryReadCard(condition))

	var decks []model.Deck
	var deck model.Deck
	conn := mysql.ConnectionDB()
	defer conn.Close()
	registry, err := conn.Query(mysql.QueryReaDeck(condition))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for registry.Next() {
		err := registry.Scan(&deck.Id, &deck.Name, &deck.Image, &deck.IdUser)
		if err != nil {
			fmt.Println(err.Error())
		}
		decks = append(decks, deck)
	}

	return decks, nil
}
