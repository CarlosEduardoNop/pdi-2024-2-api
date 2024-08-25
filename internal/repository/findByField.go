package repository

import (
	"database/sql"
	"fmt"

	"forum-api/pkg/db"
)

func FindByField(table string, field string, value string) (*sql.Row, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer conn.Close()

	sql := fmt.Sprintf("SELECT * FROM %s WHERE %s = ? LIMIT 1", table, field)

	res := conn.QueryRow(sql, value)

	return res, nil
}
