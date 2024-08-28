package forum

import (
	"fmt"
	"forum-api/internal/repository"
)

func All() ([]*Forum, error) {
	res, err := repository.All("forums", "1000")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	forums := []*Forum{}

	for res.Next() {
		var forum Forum

		err = res.Scan(&forum.ID, &forum.Name, &forum.Description, &forum.Image, &forum.CreatedAt, &forum.UpdatedAt)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		forums = append(forums, &forum)
	}

	return forums, nil
}
