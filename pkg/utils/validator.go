package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Valid *validator.Validate

func SetupValidate() {
	Valid = validator.New()
}

func ValidateThis(data interface{}) map[string]string {
	errors := make(map[string]string)
	if err := Valid.Struct(data); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationErr := range validationErrors {
			field := strings.ToLower(validationErr.Field())
			errorMsg := buildErrorMessage(field, validationErr.Tag(), validationErr.Param())
			errors[field] = errorMsg
		}
	}
	return errors
}

func buildErrorMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("The '%s' field is required.", field)
	case "min":
		return fmt.Sprintf("The '%s' field must have a minimum value of %s.", field, param)
	default:
		return fmt.Sprintf("Validation failed on '%s' field with '%s' tag.", field, tag)
	}
}
