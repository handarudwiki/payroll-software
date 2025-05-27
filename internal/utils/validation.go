package utils

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func isDate(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil
}

var validate = validator.New()

func ValidateRequest(data interface{}) []ValidationError {
	var errors []ValidationError
	validate.RegisterValidation("date", isDate)

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
	case "date":
		message = "Field " + e.Field() + " must be a valid date in the format YYYY-MM-DD"
	default:
		message = "Field " + e.Field() + " is not valid"
	}

	return message

}
