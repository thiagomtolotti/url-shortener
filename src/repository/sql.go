package repository

import (
	"database/sql"
	"fmt"
)

type SQLRepository struct {
	conn *sql.DB
}

func NewSQLRepository(c *sql.DB) Repository {
	return &SQLRepository{
		conn: c,
	}
}

func (r *SQLRepository) GetURL(id string) (string, error) {
	row, err := r.conn.Query(`SELECT url FOM shortened WHERE id = $1`, id)
	if err != nil {
		return "", fmt.Errorf("error fetching url: %w", err)
	}
	defer row.Close()

	row.Close()
	var url string

	if err := row.Scan(&url); err != nil {
		return "", fmt.Errorf("error fetching url: %w", err)
	}

	return url, nil
}

func (r *SQLRepository) CreateURL(url string, id string) error {
	_, err := r.conn.Query(`INSERT INTO shortened (id, url) VALUES ($1, $2)`)
	if err != nil {
		return fmt.Errorf("error inserting url: %w", err)
	}

	return nil
}
