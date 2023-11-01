package controllers

import (
	"las_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BookImportInput struct {
	ISBN            string `json:"isbn" binding:"required"`
	ISBN13          string `json:"isbn13" binding:"required"`
	Title           string `json:"title" binding:"required"`
	Authors         string `json:"authors" binding:"required"`
	Publisher       string `json:"publisher" binding:"required"`
	PublicationDate string `json:"publication_date" binding:"required"`
	FrappeBookID    string `json:"frappe_book_id" binding:"required"`
	Quantity        uint   `json:"quantity" binding:"required"`
}

type ImportPayload struct {
	Books []BookImportInput `binding:"dive"`
}

func Import(ctx *gin.Context) {
	var inventory ImportPayload
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var books models.Inventory
	for _, b := range inventory.Books {
		for i := 1; i <= int(b.Quantity); i++ {
			date, _ := time.Parse("1/2/2006", b.PublicationDate)
			book := models.Book{
				ISBN:            b.ISBN,
				ISBN13:          b.ISBN13,
				Title:           b.Title,
				Authors:         b.Authors,
				Publisher:       b.Publisher,
				PublicationDate: date,
				FrappeBookID:    b.FrappeBookID,
			}
			books = append(books, &book)
		}
	}

	ok, err := books.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "success": ok})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": ok})
}
