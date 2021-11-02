package models

import (
	"net/http"
	"strconv"
	"time"

	"github.com/BookStore/pkg/db"
	"github.com/BookStore/pkg/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		Quarantine:     1,
		PublisherID:    1,
		AuthorID:       1,
	}

	if err := ctx.BindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "can not bindJSON ! Check your JSON request",
		})
		return
	}
	if err := h.DB.Create(&book); err.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":            true,
			"message":          "can not bindJSON ! Check your JSON request",
			"ErrorDescription": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, &book)
}

func (bookModel BookModel) FindBook(ctx *gin.Context) {

	id, err /*error*/ := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "id: " + ctx.Param("id") + " khong hop le",
		})
		return
	}

	db := db.GetDB()
	var book entities.Book
	db.Where("id = ?", id).Find(&book)
	db.Preload(clause.Associations).Find(&book)
	ctx.JSON(http.StatusOK, &book)
}
