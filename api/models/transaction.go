package models

import (
	"time"
)

type Transaction struct {
	Base
	PatronID   uint      `gorm:"not null" json:"patron_id"`
	BookID     uint      `gorm:"not null" json:"book_id"`
	DueAt      time.Time `gorm:"not null;default:now() + '21 days'::interval" json:"due_at"`
	ReturnedAt time.Time `gorm:"null" json:"returned_at"`
}
