package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	mapURL := make(map[string]string)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/api/shorten", func(c *gin.Context) {
		var json struct {
			URL string `json:"url"`
		}

		if err := c.BindJSON(&json); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		shortCode := stringWithCharset(6, charset)
		mapURL[shortCode] = json.URL

		shortURL := "http://localhost:8080/s/" + shortCode

		c.JSON(200, gin.H{
			"shortCode": shortCode,
			"shortUrl":  shortURL,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
