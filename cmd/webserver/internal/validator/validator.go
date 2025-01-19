package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = map[string]string{}
	}

	if _, ok := v.FieldErrors[key]; !ok {
		v.FieldErrors[key] = message
	}
}

func Blank(value string) bool {
	return strings.TrimSpace(value) == ""
}

func MoreThanMaxChars(value string, maxN int) bool {
	return utf8.RuneCountInString(value) > maxN
}

func PermittedValue[T comparable](value T, permittedValues []T) bool {
	if permittedValues == nil {
		return false
	}

	return slices.Contains(permittedValues, value)
}
