package deckdb

import (
	"fmt"

	"github.com/sebarray/backendproyect/db/mysql"
	"github.com/sebarray/backendproyect/model"
)

func (d DeckDB) UpdateDeck(deck model.Deck) error {

	conn := mysql.ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(mysql.QueryDeckUpdate(deck))
	if err != nil {
		return err
	}
	result, err := insert.Exec()
	if err != nil {
		return err
	}
	rowAf, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAf == 0 {
		return fmt.Errorf("se ocasiono un error resvise su request")
	}

	return nil
}
