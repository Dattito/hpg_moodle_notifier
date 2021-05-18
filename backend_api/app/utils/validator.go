package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var Validate *validator.Validate

func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

func ValidationErrorsToText(err error) string {
	keyToV := map[string]string{
		"PhoneNumber":      "Telefon-Nummer",
		"MoodleToken":      "Moodle-Token",
		"ID":               "id",
		"VerificationCode": "Verifizierungs-Code",
	}
	msg := ""

	for _, err := range err.(validator.ValidationErrors) {
		if msg != "" {
			msg += " "
		}

		msg += keyToV[err.Field()]

		if strings.Contains(err.Error(), "required") {
			msg += " fehlt."
		} else {
			msg += " ist ungültig."
		}
	}
	return msg
}
