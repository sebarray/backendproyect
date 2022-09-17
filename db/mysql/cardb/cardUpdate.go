package cardb

import (
	"fmt"

	"github.com/sebarray/backendproyect/db/mysql"
	"github.com/sebarray/backendproyect/model"
)

func (c CardDB) UpdateCard(card model.Card) error {
	fmt.Println(mysql.QueryCardUpdate(card))
	conn := mysql.ConnectionDB()
	defer conn.Close()
	insert, err := conn.Prepare(mysql.QueryCardUpdate(card))
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
