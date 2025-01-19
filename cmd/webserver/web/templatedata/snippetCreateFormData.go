package templatedata

import (
	"github.com/gregidonut/snippetbox/cmd/webserver/internal/validator"
)

type SnippetCreateFormData struct {
	Title   string
	Content string
	Expires int
	validator.Validator
}

func NewSnippetCreateFormData(title, content string, expires int) SnippetCreateFormData {
	form := SnippetCreateFormData{
		Title:   title,
		Content: content,
		Expires: expires,
	}
	if validator.Blank(title) {
		form.AddFieldError("title", "This field cannot be blank")
	} else if validator.MoreThanMaxChars(title, 100) {
		form.AddFieldError("title", "this field cannot be more than 100 chars long")
	}

	if validator.Blank(content) {
		form.AddFieldError("content", "This field cannot be blank")
	}

	if !validator.PermittedValue(expires, []int{1, 7, 365}) {
		form.AddFieldError("expires", "This field must be equal to 1, 7 or 365")
	}

	return form
}
