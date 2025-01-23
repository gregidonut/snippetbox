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

func NewSnippetCreateFormData() *SnippetCreateFormData {
	return &SnippetCreateFormData{
		Expires: 7,
	}
}

func (f *SnippetCreateFormData) GetTitle() string {
	return f.Title
}

func (f *SnippetCreateFormData) GetContent() string {
	return f.Content
}

func (f *SnippetCreateFormData) GetExpires() int {
	return f.Expires
}

func (f *SnippetCreateFormData) GetValidator() validator.Validator {
	return f.Validator
}

func (f *SnippetCreateFormData) Validate() {
	f.CheckField(
		validator.NotBlank(f.GetTitle()),
		"title",
		"This field cannot be blank")
	f.CheckField(
		validator.LessThanMaxChars(f.GetTitle(), 100),
		"title",
		"this field cannot be more than 100 chars long",
	)
	f.CheckField(
		validator.NotBlank(f.GetContent()),
		"content",
		"This field cannot be blank",
	)
	f.CheckField(
		validator.PermittedValue(f.GetExpires(), []int{1, 7, 365}),
		"expires",
		"This field must be equal to 1, 7 or 365",
	)
}
