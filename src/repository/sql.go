package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
	var url string

	err := r.conn.QueryRow(`SELECT url FROM shortened WHERE id = $1`, id).Scan(&url)
	if err != nil {
		return "", fmt.Errorf("error fetching url: %w", err)
	}

	return url, nil
}

func (r *SQLRepository) CreateURL(url string, id string) error {
	_, err := r.conn.Query(`INSERT INTO shortened (id, url) VALUES ($1, $2)`, id, url)
	if err != nil {
		return fmt.Errorf("error inserting url: %w", err)
	}

	return nil
}

func (r *SQLRepository) DeleteURL(id string) error {
	_, err := r.conn.Query(`DELETE FROM shortened WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting url: %w", err)
	}

	return nil
}

func (r *SQLRepository) Exists(id string) (bool, error) {
	var exists bool

	err := r.conn.QueryRow(`SELECT COUNT(*) > 0 FROM shortened WHERE id = $1`, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if id exists: %w", err)
	}

	return exists, nil
}
