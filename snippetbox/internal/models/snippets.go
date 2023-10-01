package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Updated time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title string, content string, expires time.Time) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created_at, updated_at, expires_at) VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $3)`

	result, err := m.DB.Exec(context.Background(), stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	fmt.Printf("result::: %v", result)
	return 0, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created_at, updated_at, expires_at FROM snippets WHERE expires_at > CURRENT_TIMESTAMP AND id = $1`

	row := m.DB.QueryRow(context.Background(), stmt, id)

	var s Snippet

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Updated, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
