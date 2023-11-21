package database

import (
	"las_api/models"

	"github.com/jmoiron/sqlx"
)

type PatronStore struct {
	*sqlx.DB
}

func (s *PatronStore) PatronById(id string) (models.Patron, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PatronStore) Patrons() ([]models.Patron, error) {
	panic("not implemented") // TODO: Implement
}

func (s *PatronStore) CreatePatron(p *models.Patron) error {
	panic("not implemented") // TODO: Implement
}

func (s *PatronStore) UpdatePatron(p *models.Patron) error {
	panic("not implemented") // TODO: Implement
}
