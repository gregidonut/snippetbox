package templatedata

import (
	"net/http"
	"time"

	"github.com/gregidonut/snippetbox/cmd/webserver/internal/models"
	"github.com/gregidonut/snippetbox/cmd/webserver/internal/validator"
	"github.com/gregidonut/snippetbox/cmd/webserver/web/appinterface"
)

type FormData interface {
	*BlankFormData |
		*SnippetCreateFormData |
		*UserSignupFormData |
		*UserLoginFormData
	GetValidator() validator.Validator
}

type TemplateData[T FormData] struct {
	models.Snippet
	Snippets    []models.Snippet
	CurrentYear int
	Flash       string
	Form        T
}

func New[T FormData](
	r *http.Request, app appinterface.App,
) TemplateData[T] {
	return TemplateData[T]{
		CurrentYear: time.Now().Year(),
		Flash:       app.PopString(r.Context(), "flash"),
	}
}
