package templatedata

import (
	"github.com/gregidonut/snippetbox/cmd/webserver/internal/validator"
)

type UserLoginFormData struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (f *UserLoginFormData) GetEmail() string {
	return f.Email
}

func (f *UserLoginFormData) GetPassword() string {
	return f.Password
}

func (f *UserLoginFormData) GetValidator() validator.Validator {
	return f.Validator
}

func (f *UserLoginFormData) Validate() {
	f.CheckField(
		validator.NotBlank(f.GetEmail()),
		"email",
		"This field cannot be blank",
	)
	f.CheckField(
		validator.Matches(f.GetEmail(), validator.EmailRX),
		"email",
		"This field must be a vaild email address",
	)
	f.CheckField(
		validator.NotBlank(f.Password),
		"password",
		"This field cannot be blank",
	)
}
