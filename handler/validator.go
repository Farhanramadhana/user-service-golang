package handler

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func NewValidator(translator ut.Translator) *validator.Validate {
	validate := validator.New()

	// Register custom translations for better error messages
	_ = validate.RegisterTranslation("hasupper", translator, func(ut ut.Translator) error {
		return ut.Add("hasupper", "{0} must contain uppercase letter", true)
	}, TranslationFn)
	_ = validate.RegisterTranslation("hasnumber", translator, func(ut ut.Translator) error {
		return ut.Add("hasnumber", "{0} must contain number", true)
	}, TranslationFn)
	_ = validate.RegisterTranslation("hasspecial", translator, func(ut ut.Translator) error {
		return ut.Add("hasspecial", "{0} must contain special character", true)
	}, TranslationFn)

	// Register custom validation functions
	_ = validate.RegisterValidation("hasupper", hasUpper)
	_ = validate.RegisterValidation("hasnumber", hasNumber)
	_ = validate.RegisterValidation("hasspecial", hasSpecial)

	return validate
}

func TranslationFn(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		// Fallback to default message if translation not found
		return fe.Error()
	}

	return t
}

func hasUpper(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func hasNumber(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func hasSpecial(fl validator.FieldLevel) bool {
	specialChars := "!@#$%^&*()-_=+[]{}|;:'\"<>,.?/~`"
	for _, char := range fl.Field().String() {
		if strings.ContainsRune(specialChars, char) {
			return true
		}
	}
	return false
}
