package models

import "github.com/gregidonut/snippetbox/cmd/webserver/web/templatedata"

func (m *SnippetModel) Insert(f templatedata.SnippetCreateFormData) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
VALUES ($1, $2, now(), now() + $3 * INTERVAL '1 days')
RETURNING id`

	id := 0
	err := m.QueryRow(
		stmt, f.Title, f.Content, f.Expires,
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
