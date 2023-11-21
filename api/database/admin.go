package database

import (
	"fmt"
	"las_api/models"

	"github.com/jmoiron/sqlx"
)

type AdminStore struct {
	*sqlx.DB
}

func (s *AdminStore) AdminById(id string) (models.Admin, error) {
	var a models.Admin
	if err := s.Get(&a, `SELECT * FROM admins WHERE id = $1`, id); err != nil {
		return models.Admin{}, fmt.Errorf("error getting admin: %w", err)
	}
	return a, nil
}

func (s *AdminStore) AdminByUsername(username string) (models.Admin, error) {
	var a models.Admin
	if err := s.Get(&a, `SELECT * FROM admins WHERE username = $1`, username); err != nil {
		return models.Admin{}, fmt.Errorf("error getting admin: %w", err)
	}
	return a, nil
}

func (s *AdminStore) Admins() ([]models.Admin, error) {
	var aa []models.Admin
	if err := s.Select(&aa, `SELECT * FROM admins`); err != nil {
		return []models.Admin{}, fmt.Errorf("error getting admins: %w", err)
	}
	return aa, nil
}

func (s *AdminStore) CreateAdmin(a *models.Admin) error {
	if err := s.Get(a,
		`INSERT INTO admins (id, created_at, updated_at, username, password) VALUES ($1, DEFAULT, DEFAULT, $2, $3) RETURNING *`,
		a.ID,
		a.Username,
		a.Password,
	); err != nil {
		return fmt.Errorf("error creating admin: %w", err)
	}
	return nil
}
