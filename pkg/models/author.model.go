package models

import (
	"net/http"
	"strconv"

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
			"message":          "can not CreateNewAuthor! Check your JSON request",
			"ErrorDescription": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, &author)
}

func (h AuthorModel) FindAuthor(ctx *gin.Context) {

	id, err /*error*/ := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "id: " + ctx.Param("id") + " khong hop le",
		})
		return
	}
	//"bookstore", "root", "example", "127.0.0.1:3306"
	author := entities.Author{}
	h.DB.Where("id = ?", id).First(&author)
	if (entities.Author{}) == author {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "id: " + ctx.Param("id") + " khong ton tai",
		})
		return
	}
	h.DB.Preload(clause.Associations).First(&author)
	ctx.JSON(http.StatusOK, &author)
}
