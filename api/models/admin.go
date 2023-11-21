package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	ID        string     `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
	Username  string     `db:"username" json:"username"`
	Password  string     `db:"password" json:"-"`
}

func (a *Admin) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
}
