package templatedata

import "github.com/gregidonut/snippetbox/cmd/webserver/internal/validator"

type BlankFormData struct {
	validator.Validator
}

func (f *BlankFormData) GetValidator() validator.Validator {
	return f.Validator
}
