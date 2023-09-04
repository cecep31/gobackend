package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Valid *validator.Validate

func SetupValidate() {
	Valid = validator.New()
}

func ValidateThis(data interface{}) map[string]string {
	if err := Valid.Struct(data); err != nil {
		errMsgs := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()

			var errMsg string
			switch tag {
			case "required":
				errMsg = fmt.Sprintf("The field '%s' is required.", field)
			case "min":
				errMsg = fmt.Sprintf("The value of '%s' must be greater than or equal to %s.", field, err.Param())
			default:
				errMsg = fmt.Sprintf("Validation failed on field '%s' with tag '%s'.", field, tag)
			}

			errMsgs[field] = errMsg
		}
		return errMsgs
	}
	return nil
}
