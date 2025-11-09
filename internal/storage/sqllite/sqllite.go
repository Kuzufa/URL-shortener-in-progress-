package sqllite

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqllite.New"

	db, err := sql.Open("sqllite3", "./url-shortner.db")
	if err != nil {
		return nil, fmt.Errorf("%s, %w", op, err)
	}
}
