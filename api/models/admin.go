package models

import (
	"html"
	"las_api/database"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	Base
	FirstName string `gorm:"not null" json:"first_name"`
	Username  string `gorm:"not null;unique" json:"username"`
	Password  string `gorm:"not null" json:"-"`
}

func (a *Admin) Save() (*Admin, error) {
	err := database.DB.Create(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}

func (a *Admin) BeforeSave(*gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.FirstName = html.EscapeString(strings.TrimSpace(a.FirstName))
	a.Username = html.EscapeString(strings.TrimSpace(a.Username))
	a.Password = string(hash)
	return nil
}

func (a *Admin) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
}

func FindAdminByUsername(username string) (Admin, error) {
	var a Admin
	err := database.DB.Where("username=?", username).Find(&a).Error
	if err != nil {
		return Admin{}, err
	}
	return a, nil
}

func FindAdminById(id uint) (Admin, error) {
	var a Admin
	err := database.DB.Where("id=?", id).Find(&a).Error
	if err != nil {
		return Admin{}, err
	}
	return a, nil
}
