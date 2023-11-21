package api

import "las_api/models"

type AdminStore interface {
	AdminById(id string) (models.Admin, error)
	AdminByUsername(username string) (models.Admin, error)
	Admins() ([]models.Admin, error)
	CreateAdmin(a *models.Admin) error
}

type PatronStore interface {
	PatronById(id string) (models.Patron, error)
	Patrons() ([]models.Patron, error)
	CreatePatron(p *models.Patron) error
	UpdatePatron(p *models.Patron) error
}

// type BookStore interface {
// 	BookById(id string) (models.Book, error)
// 	Books() ([]models.Book, error)
// 	CreateBook(b models.Book) error
// 	UpdateBook(b models.Book) error
// }

type Store interface {
	AdminStore
	PatronStore
	// BookStore
}
