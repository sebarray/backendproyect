package cardb

import (
	"fmt"

	"github.com/sebarray/backendproyect/db/mysql"
	"github.com/sebarray/backendproyect/model"
)

func (c CardDB) ReadCard(def, id string) ([]model.Card, error) {

	condition := "where iduser = "
	if def == "default" {
		condition += fmt.Sprintf("'%s'", def)
	} else {
		condition += fmt.Sprintf("'%s'", id)
	}
	fmt.Println(mysql.QueryReadCard(condition))

	var cards []model.Card
	var card model.Card
	conn := mysql.ConnectionDB()
	defer conn.Close()
	registry, err := conn.Query(mysql.QueryReadCard(condition))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for registry.Next() {
		err := registry.Scan(&card.Id, &card.Title, &card.Description, &card.Image, &card.IdUser)
		if err != nil {
			fmt.Println(err.Error())
		}
		cards = append(cards, card)
	}

	return cards, nil
}
