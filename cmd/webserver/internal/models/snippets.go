package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Expires time.Time `json:"expires"`
}

type SnippetModel struct {
	*sql.DB
}

func NewSnippetModel(db *sql.DB) *SnippetModel {
	return &SnippetModel{DB: db}
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
