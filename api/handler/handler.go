package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateError(err error) []string {
	if ve, ok := err.(validator.ValidationErrors); ok {
		errors := make([]string, len(ve))
		for i, fe := range ve {
			errors[i] = fmt.Sprintf("%s: %s", fe.Field(), fe.Tag())
		}

		return errors
	}

	return []string{err.Error()}
}
