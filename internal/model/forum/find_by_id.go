package forum

import (
	"fmt"

	"forum-api/internal/repository"
)

func FindById(id int) (*Forum, error) {
	res, err := repository.FindById("forums", id)

	if err != nil {
		return nil, err
	}

	var forum Forum

	err = res.Scan(&forum.ID, &forum.Name, &forum.Description, &forum.Image, &forum.CreatedAt, &forum.UpdatedAt)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &forum, nil
}
