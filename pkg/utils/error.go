package utils

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var error ErrorResponse
			error.FailedField = err.StructNamespace()
			error.Tag = err.Tag()
			error.Value = err.Param()
			errors = append(errors, &error)
		}
	}

	return errors
}
