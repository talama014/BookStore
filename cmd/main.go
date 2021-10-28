package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
