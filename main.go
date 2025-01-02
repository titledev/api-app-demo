package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

	r.POST("/add", func(c *gin.Context) {
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
