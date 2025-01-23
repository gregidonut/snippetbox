package models

type snippetCreateFormData interface {
	GetTitle() string
	GetContent() string
	GetExpires() int
}

func (m *SnippetModel) Insert(f snippetCreateFormData) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
VALUES ($1, $2, now(), now() + $3 * INTERVAL '1 days')
RETURNING id`

	id := 0
	err := m.QueryRow(
		stmt, f.GetTitle(), f.GetContent(), f.GetExpires(),
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
