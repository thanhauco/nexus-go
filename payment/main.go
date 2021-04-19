package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.POST("/pay", func(c *gin.Context) {
        // Mock payment processing
        c.JSON(http.StatusOK, gin.H{"status": "paid", "transactionId": "tx_12345"})
    })
    log.Println("Payment Service running on :8083")
    r.Run(":8083")
}
