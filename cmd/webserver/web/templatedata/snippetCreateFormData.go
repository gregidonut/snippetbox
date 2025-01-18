package templatedata

import (
	"strings"
	"unicode/utf8"
)

type SnippetCreateFormData struct {
	Title       string
	Content     string
	Expires     int
	FieldErrors map[string]string
}

func NewSnippetCreateFormData(title, content string, expires int) *SnippetCreateFormData {
	form := SnippetCreateFormData{
		Title:       title,
		Content:     content,
		Expires:     expires,
		FieldErrors: map[string]string{},
	}
	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(form.Title) > 100 {
		form.FieldErrors["title"] = "this field cannot be more than 100 chars long"
	}

	if strings.TrimSpace(form.Content) == "" {
		form.FieldErrors["content"] = "This field cannot be blank"
	}

	if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
		form.FieldErrors["content"] = "This field must be equal to 1, 7 or 365"
	}

	return &form
}
