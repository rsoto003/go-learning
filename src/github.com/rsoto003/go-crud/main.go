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
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}
