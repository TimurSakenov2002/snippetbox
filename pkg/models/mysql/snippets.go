package mysql

import (
	"alexedwards.net/snippetbox/pkg/models"
	"database/sql"
	"errors" // New import
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}
