package models

func (m *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > now()
    ORDER BY id DESC lIMIT 10`

	rows, err := m.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := []Snippet{}

	for rows.Next() {
		s := Snippet{}
		err = rows.Scan(
			&s.ID,
			&s.Title,
			&s.Content,
			&s.Created,
			&s.Expires,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, s)
	}

	return payload, nil
}
