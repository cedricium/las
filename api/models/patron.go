package models

import "time"

type Patron struct {
	ID        string     `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
	FirstName string     `db:"first_name" json:"first_name"`
	LastName  string     `db:"last_name" json:"last_name"`
	Address   string     `db:"address" json:"address"`
	ZipCode   string     `db:"zip_code" json:"zip_code"`
	Balance   uint16     `db:"balance" json:"balance"`
}
