package models

type Patron struct {
	Base
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Address   string `gorm:"not null" json:"address"`
	ZipCode   string `gorm:"type:varchar(5);not null" json:"zip_code"`

	Transactions []Transaction
}
