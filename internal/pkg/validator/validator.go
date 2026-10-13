package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func ValidateStruct(s any) map[string]string {
	err := validate.Struct(s)

	if err == nil {
		return nil
	}

	errorsMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {

			fieldName := fieldErr.Field()
			tag := fieldErr.Tag()

			switch tag {
			case "required":
				errorsMap[fieldName] = fieldName + " wajib diisi"
			default:
				errorsMap[fieldName] = fieldName + " tidak valid (" + tag + ")"
			}
		}
	}

	return errorsMap
}
