package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rsoto003/go-learning/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "working correctly",
		})
	})
	r.Run()
}
