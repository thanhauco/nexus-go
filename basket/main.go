package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func main() {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

	r := gin.Default()
	r.GET("/basket/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"items": []string{}})
	})
	log.Println("Basket Service running on :8082")
	r.Run(":8082")
}
