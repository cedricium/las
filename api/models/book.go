package models

import (
	"time"
)

type Book struct {
	Base
	ISBN            string    `gorm:"type:varchar(10);not null;" json:"isbn"`
	ISBN13          string    `gorm:"type:varchar(13);not null;" json:"isbn13"`
	Title           string    `gorm:"not null" json:"title"`
	Authors         string    `gorm:"not null"`
	Publisher       string    `gorm:"not null" json:"publisher"`
	PublicationDate time.Time `gorm:"not null" json:"publication_date"`
	FrappeBookID    string    `gorm:"not null" json:"frappe_book_id"`

	Transactions []Transaction

	// === Frappe Book Schema ===
	// "bookID":"39763",
	// "title":"The Mystical Poems of Rumi 1: First Selection  Poems 1-200",
	// "authors":"Rumi/A.J. Arberry",
	// "average_rating":"4.28",
	// "isbn":"0226731510",
	// "isbn13":"9780226731513",
	// "language_code":"eng",
	// "num_pages":"208",
	// "ratings_count":"114",
	// "text_reviews_count":"8",
	// "publication_date":"3/15/1974",
	// "publisher":"University Of Chicago Press"
}
