package models

import (
	"net/http"
	"strconv"

	"github.com/BookStore/pkg/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PublisherModel struct {
	DB *gorm.DB
}

func (h PublisherModel) CreatePublisher(ctx *gin.Context) {
	publisher := entities.Publisher{
		Name:        "Unknow",
		Description: "Unknow",
	}

	if err := ctx.BindJSON(&publisher); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "can not bindJSON ! Check your JSON request",
		})
		return
	}
	if err := h.DB.Create(&publisher); err.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":            true,
			"message":          "can not CreateNewPublisher! Check your JSON request",
			"ErrorDescription": err.Error,
		})
		return
	}
	ctx.JSON(http.StatusOK, &publisher)
}

func (h PublisherModel) FindPublisher(ctx *gin.Context) {

	id, err /*error*/ := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "id: " + ctx.Param("id") + " khong hop le",
		})
		return
	}
	//"bookstore", "root", "example", "127.0.0.1:3306"
	publisher := entities.Publisher{}
	h.DB.Where("id = ?", id).First(&publisher)
	if (entities.Publisher{}) == publisher {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "id: " + ctx.Param("id") + " khong ton tai",
		})
		return
	}
	h.DB.Preload(clause.Associations).First(&publisher)
	ctx.JSON(http.StatusOK, &publisher)
}
