package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ip, _ := net.InterfaceAddrs()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": ip,
		})
	})
	r.Run("0.0.0.0:8080")
}
