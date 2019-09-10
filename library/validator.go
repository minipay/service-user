package library

import (
	"errors"
	"gopkg.in/go-playground/validator.v9"
)

func IsRequestValid(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return false, errors.New(err.Field() + " is " + err.Tag())
		}
	}
	return true, nil
}