package models

func (m *SnippetModel) Insert(
	title string,
	content string,
	expires int,
) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
VALUES ($1, $2, now(), now() + $3 * INTERVAL '1 days')
RETURNING id`

	id := 0
	err := m.QueryRow(
		stmt, title, content, expires,
	).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}
