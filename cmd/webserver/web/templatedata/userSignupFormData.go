package templatedata

import "github.com/gregidonut/snippetbox/cmd/webserver/internal/validator"

type UserSignupFormData struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (f *UserSignupFormData) GetName() string {
	return f.Name
}

func (f *UserSignupFormData) GetEmail() string {
	return f.Email
}

func (f *UserSignupFormData) GetPassword() string {
	return f.Password
}

func (f *UserSignupFormData) GetValidator() validator.Validator {
	return f.Validator
}

func (f *UserSignupFormData) Validate() {
	f.CheckField(
		validator.NotBlank(f.GetName()),
		"name",
		"This field cannot be blank",
	)
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
		validator.NotBlank(f.GetPassword()),
		"password",
		"This field cannot be blank",
	)
	f.CheckField(
		validator.MoreThanMinChars(f.GetPassword(), 8),
		"password",
		"field must be at least 8 characters long",
	)
}
