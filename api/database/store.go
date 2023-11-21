package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	*AdminStore
	*PatronStore
}

func NewStore(dsn string) (*Store, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		AdminStore:  &AdminStore{DB: db},
		PatronStore: &PatronStore{DB: db},
	}, nil
}
