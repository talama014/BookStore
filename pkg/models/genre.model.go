package models

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/BookStore/pkg/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GenreModel struct {
	DB *gorm.DB
}

func (h GenreModel) CreateGenre(ctx *gin.Context) {
	genre := entities.Genre{
		Name:        "Unknow",
		Description: "Unknow",
	}

	if err := ctx.BindJSON(&genre); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "can not bindJSON ! Check your JSON request",
		})
		return
	}
	if err := h.DB.Create(&genre); err.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":            true,
			"message":          "can not CreateNewGenre! Check your JSON request",
			"ErrorDescription": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, &genre)
}

func (h GenreModel) FindGenre(ctx *gin.Context) {

	id, err /*error*/ := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "id: " + ctx.Param("id") + " khong hop le",
		})
		return
	}
	//"bookstore", "root", "example", "127.0.0.1:3306"
	genre := entities.Genre{}
	h.DB.Where("id = ?", id).First(&genre)
	if reflect.DeepEqual((entities.Genre{}), genre) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "id: " + ctx.Param("id") + " khong ton tai",
		})
		return
	}
	h.DB.Preload(clause.Associations).First(&genre)
	ctx.JSON(http.StatusOK, &genre)
}
