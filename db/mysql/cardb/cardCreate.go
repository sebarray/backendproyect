package cardb

import (
	"github.com/google/uuid"
	"github.com/sebarray/backendproyect/db/mysql"
	"github.com/sebarray/backendproyect/model"
)

func (c CardDB) CreateCard(card model.Card) error {
	connection := mysql.ConnectionDB()
	defer connection.Close()
	card.Id = uuid.New().String()
	insert, err := connection.Prepare(mysql.QueryCreateCard(card))
	if err != nil {
		return err
	}
	_, err = insert.Exec()
	if err != nil {
		return err
	}
	return err

}
