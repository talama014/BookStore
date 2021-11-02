package models

import (
	"net/http"
	"time"

	"github.com/BookStore/pkg/db"
	"github.com/BookStore/pkg/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookModel struct {
	DB *gorm.DB
}

func (h BookModel) Book_Create(ctx *gin.Context) {
	book := entities.Book{
		Title:          "Unknown",
		Total_pages:    0,
		Isbn:           "chua cap nhat",
		Publisher_date: time.Now(),
	}

	if err := ctx.BindJSON(&book); err != nil {
		panic("CAN NOT BIND JSON INTO BOOK!!!")
	}
	h.DB.Create(&book)
	ctx.JSON(http.StatusOK, &book)
}

func (bookModel BookModel) FindBook(ctx *gin.Context) {
	db := db.GetDB()
	var book entities.Book
	db.Where("id = ?", 2).First(&book)
	ctx.JSON(http.StatusOK, &book)
}
