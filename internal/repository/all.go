package repository

import (
	"database/sql"
	"fmt"
	"forum-api/pkg/db"
)

func All(table string) (*sql.Rows, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer conn.Close()

	sql := fmt.Sprintf("SELECT * FROM %s", table)

	res, err := conn.Query(sql)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}
