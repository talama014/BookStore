package models

import (
	"net/http"
	"strconv"

	"github.com/BookStore/pkg/db"
	"github.com/BookStore/pkg/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthorModel struct {
	DB *gorm.DB
}

func (h AuthorModel) Author_Create(ctx *gin.Context) {
	author := entities.Author{
		Name:        "Unknow",
		Description: "",
	}

	if err := ctx.BindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "can not bindJSON ! Check your JSON request",
		})
		return
	}
	if err := h.DB.Create(&author); err.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":            true,
			"message":          "can not bindJSON ! Check your JSON request",
			"ErrorDescription": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, &author)
}

func (authorModel AuthorModel) FindAuthor(ctx *gin.Context) {

	id, err /*error*/ := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "id: " + ctx.Param("id") + " khong hop le",
		})
		return
	}

	db := db.GetDB()
	var author entities.Author
	db.Where("id = ?", id).Find(&author)
	db.Preload(clause.Associations).Find(&author)
	ctx.JSON(http.StatusOK, &author)
}
