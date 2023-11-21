package models

import "time"

type Book struct {
	ID              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
	ISBN            string
	ISBN13          string
	Title           string
	Authors         string
	Publisher       string
	PublicationDate time.Time
	FrappeBookID    string
}
