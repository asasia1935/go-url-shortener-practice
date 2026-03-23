package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type URLMappingData struct {
	store map[string]string
}

func main() {
	data := &URLMappingData{
		store: make(map[string]string),
	}

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
	r.POST("/api/shorten", data.shortenURL)
	r.GET("/s/:shortCode", func(c *gin.Context) {
		shortCode := c.Param("shortCode")
		originalURL, exists := data.store[shortCode]
		if !exists {
			c.JSON(404, gin.H{
				"error": "Short code not found",
			})
			return
		} else {
			c.Redirect(302, originalURL)
		}
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

func (d *URLMappingData) shortenURL(c *gin.Context) {
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

	// 기존 shortCode가 존재하는 경우 새로운 shortCode 생성
	for {
		_, exists := d.store[shortCode]
		if !exists {
			break
		}
		shortCode = stringWithCharset(6, charset)
	}

	d.store[shortCode] = json.URL
	c.JSON(200, gin.H{
		"shortCode": shortCode,
		"shortUrl":  "http://localhost:8080/s/" + shortCode,
	})
}
