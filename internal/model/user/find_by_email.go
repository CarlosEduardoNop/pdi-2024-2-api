package user

import (
	"fmt"

	"forum-api/internal/repository"
)

func FindByEmail(email string) (*User, error) {
	res, err := repository.FindByField("users", "email", email)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user User

	err = res.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}
