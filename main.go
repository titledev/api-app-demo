package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Add validateAPIKey middleware function
func validateAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		expectedAPIKey := os.Getenv("API_KEY") // Get from environment variable

		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API key is missing",
			})
			return
		}

		if apiKey != expectedAPIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API key",
			})
			return
		}

		c.Next()
	}
}

func main() {
	fmt.Println("version", os.Getenv("VERSION"))
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/version", func(c *gin.Context) {
		version := os.Getenv("VERSION")
		c.JSON(200, gin.H{
			"version": version,
		})
	})

	r.POST("/add", validateAPIKey(), func(c *gin.Context) {
		var json struct {
			A int `json:"a"`
			B int `json:"b"`
		}

		if c.BindJSON(&json) == nil {
			result := add(json.A, json.B)
			c.JSON(200, gin.H{
				"result": result,
			})
		}
	})

	fmt.Println("Starting server on port 8080")
	r.Run(":8080")
}

func add(a int, b int) int {
	return a + b
}
