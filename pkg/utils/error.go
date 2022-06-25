package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field string
	Tag   string
	Value string
}

var validate = validator.New()

func ValidateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var error ErrorResponse
			error.Field = err.Field()
			error.Tag = err.Tag()
			error.Value = strings.Split(err.Error(), "Error:")[1]
			errors = append(errors, &error)
		}
	}

	return errors
}
