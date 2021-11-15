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

	db := db.GetDB(db.AdminConn)

	if err := db.AutoMigrate(&entities.Book{}, &entities.Author{}, &entities.Publisher{}, &entities.Genre{}); err != nil {
		panic("can not auto migrate database")
	}
	// Ping
	r.GET("/ping", pingHandler)

	// Book----------------------
	bookHandler := models.BookModel{
		DB: db,
	}
	// tao sach moi
	r.POST("/createbook", bookHandler.Book_Create)
	//Tim sach qua ID
	r.GET("/getbook/:id", bookHandler.FindBook)
	//Author---------------------
	authorHandler := models.AuthorModel{
		DB: db,
	}
	// tao author moi
	r.POST("/createauthor", authorHandler.Author_Create)
	//Tim author qua ID
	r.GET("/getauthor/:id", authorHandler.FindAuthor)

	publisherHandler := models.PublisherModel{
		DB: db,
	}
	// tao publisher moi
	r.POST("/createpublisher", publisherHandler.CreatePublisher)
	//Tim publisher qua ID
	r.GET("/getpublisher/:id", publisherHandler.FindPublisher)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
