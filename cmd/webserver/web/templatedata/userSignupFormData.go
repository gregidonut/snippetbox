package templatedata

import "github.com/gregidonut/snippetbox/cmd/webserver/internal/validator"

type UserSignupFormData struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (f *UserSignupFormData) GetValidator() validator.Validator {
	return f.Validator
}
