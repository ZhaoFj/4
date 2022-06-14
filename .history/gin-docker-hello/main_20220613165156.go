package main

import (
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	name, _ := os.Hostname()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, I am " + name,
		})
	})
	r.Run("0.0.0.0:8080")
}
