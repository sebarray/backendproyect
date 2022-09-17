package cardb

import (
	"fmt"

	"github.com/sebarray/backendproyect/db/mysql"
)

// agregar funcion para borrar de masos
func (c CardDB) DeleteCard(id string) error {

	querysql := fmt.Sprintf("DELETE FROM card WHERE id='%s' ;", id)
	conn := mysql.ConnectionDB()
	defer conn.Close()
	query, err := conn.Prepare(querysql)
	if err != nil {
		return err
	}
	result, err := query.Exec()
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
