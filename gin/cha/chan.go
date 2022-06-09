package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name    string
	Message string
	Number  int
}

func main() {
	r := gin.Default()
	r.GET("/chan", func(c *gin.Context) {
		query := c.Query("data")
		c.JSON(http.StatusOK, gin.H{"type": "JSON", "status": 200, "data": query})
	})
	r.Run("localhost:8000")
}
