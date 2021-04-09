package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"token": "fake-jwt-token"})
	})
	log.Println("Identity Service running on :8081")
	r.Run(":8081")
}
