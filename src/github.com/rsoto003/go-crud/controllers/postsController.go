package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rsoto003/go-learning/initializers"
	"github.com/rsoto003/go-learning/models"
)

func PostsCreate(c *gin.Context) {
	//Get data off request body

	//Create a post
	post := models.Post{Title: "First Post :)", Body: "Learning Golang :)"}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return created post
	c.JSON(200, gin.H{
		"post": post,
	})
}