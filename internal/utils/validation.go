package utils

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate = validator.New()

func ValidateRequest(data interface{}) []ValidationError {
	var errors []ValidationError

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field:   err.Field(),
				Message: GetErrorMessage(err),
			})
		}
	}
	if len(errors) == 0 {
		return nil
	}

	return errors
}

func GetErrorMessage(e validator.FieldError) (message string) {
	switch e.Tag() {
	case "required":
		message = "Field " + e.Field() + " is required"
	case "email":
		message = "Field " + e.Field() + " must be a valid email address"
	case "min":
		message = "Field " + e.Field() + " must be at least " + e.Param() + " characters long"
	case "max":
		message = "Field " + e.Field() + " must be at most " + e.Param() + " characters long"
	default:
		message = "Field " + e.Field() + " is not valid"
	}

	return message

}
