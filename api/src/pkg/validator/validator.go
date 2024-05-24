package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		validate.RegisterValidation("status", validateStatus)
	}
	return validate
}

func validateStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	for _, validStatus := range ValidStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}

var ValidStatuses = []string{
	"Pending",
	"InProgress",
	"Completed",
}
