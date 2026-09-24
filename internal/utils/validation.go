package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(value any) error {
	if err := validate.Struct(value); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			messages := make([]string, 0, len(validationErrors))
			for _, fieldErr := range validationErrors {
				messages = append(messages, fmt.Sprintf("%s failed on %s validation", fieldErr.Field(), fieldErr.Tag()))
			}
			return fmt.Errorf("validation failed: %s", strings.Join(messages, ", "))
		}
		return err
	}
	return nil
}
