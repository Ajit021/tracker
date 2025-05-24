package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
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

		// Log the open event or store it in DB
		log.Printf("Email opened: id=%s", id)

		// Create 1x1 transparent PNG dynamically
		img := image.NewRGBA(image.Rect(0, 0, 1, 1))
		img.Set(0, 0, color.RGBA{0, 0, 0, 0}) // Transparent pixel

		var buf bytes.Buffer
		err := png.Encode(&buf, img)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to encode pixel"})
			return
		}

		c.Data(http.StatusOK, "image/png", buf.Bytes())
	})

	// Run on environment PORT (for Render) or default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
