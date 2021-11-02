package main

import (
	"net/http"

	"github.com/BookStore/pkg/db"
	"github.com/BookStore/pkg/entities"
	"github.com/BookStore/pkg/models"
	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()

	db := db.GetDB()

	if err := db.AutoMigrate(&entities.Book{}, &entities.Author{}, &entities.Publisher{}); err != nil {
		panic("can not auto migrate database")
	}
	// Ping
	r.GET("/ping", pingHandler)
	bHandler := models.BookModel{
		DB: db,
	}

	// tao sach moi
	r.POST("/createbook", bHandler.Book_Create)
	//Tim sach qua ID
	r.GET("/getbook/:id", bHandler.FindBook)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
