package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/track", func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
			return
		}
		log.Printf("Email opened: id=%s", id)

		data, err := os.ReadFile("pixel.png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot load pixel"})
			return
		}
		c.Data(http.StatusOK, "image/png", data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
