package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

type user struct {
	Name    string
	Message string
	Number  int
}

func main() {
	r := gin.Default()
	r.GET("/chan", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"type": "JSON", "status": 200})
	})
	r.Run("localhost:8000")
}
