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
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, errors.Join(ErrNoRecord, fmt.Errorf("from %d id", id))
		}
		return Snippet{}, err
	}

	return payload, nil
}
