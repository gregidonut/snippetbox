package templatedata

import (
	"github.com/gregidonut/snippetbox/cmd/webserver/internal/validator"
)

type SnippetCreateFormData struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}

func (form *SnippetCreateFormData) Validate() {
	form.CheckField(
		validator.NotBlank(form.Title),
		"title",
		"This field cannot be blank")
	form.CheckField(
		validator.LessThanMaxChars(form.Title, 100),
		"title",
		"this field cannot be more than 100 chars long",
	)
	form.CheckField(
		validator.NotBlank(form.Content),
		"content",
		"This field cannot be blank",
	)
	form.CheckField(
		validator.PermittedValue(form.Expires, []int{1, 7, 365}),
		"expires",
		"This field must be equal to 1, 7 or 365",
	)
}
