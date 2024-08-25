package user

import (
	"database/sql"
	"fmt"

	"forum-api/internal/repository"
)

func Create(data map[string]interface{}) (sql.Result, error) {
	res, err := repository.Create("users", data)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}
