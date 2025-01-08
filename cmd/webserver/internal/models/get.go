package models

import (
	"database/sql"
	"errors"
	"fmt"
)

func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires
FROM snippets
WHERE expires > now() AND id=$1`

	payload := Snippet{}
	err := m.QueryRow(
		stmt, id,
	).Scan(
		&payload.ID,
		&payload.Title,
		&payload.Content,
		&payload.Created,
		&payload.Expires,
	)

	payload.Created.Local().Weekday()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, fmt.Errorf("%w on id: %d", ErrNoRecord, id)
		}
		return Snippet{}, err
	}

	return payload, nil
}
