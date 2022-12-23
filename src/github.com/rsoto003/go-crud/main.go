package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rsoto003/go-learning/controllers"
	"github.com/rsoto003/go-learning/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run()
}
