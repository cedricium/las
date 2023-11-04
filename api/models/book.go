package models

import (
	"las_api/database"
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

	Transactions []*Transaction

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

type Inventory []*Book

func (i *Inventory) Save() (bool, error) {
	err := database.DB.Create(i).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// TODO: add/handle pagination using Limit & Offset
func FindAllBooks() (Inventory, error) {
	var bb Inventory
	err := database.DB.Find(&bb).Error
	if err != nil {
		return Inventory{}, err
	}
	return bb, nil
}

func FindBookById(id uint) (Book, error) {
	var b = Book{Base: Base{ID: id}}
	err := database.DB.Preload("Transactions").Find(&b).Error
	if err != nil {
		return Book{}, err
	}
	return b, nil
}

func (b *Book) Update() (bool, error) {
	err := database.DB.Save(&b).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
