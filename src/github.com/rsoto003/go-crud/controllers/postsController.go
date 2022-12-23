package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rsoto003/go-learning/initializers"
	"github.com/rsoto003/go-learning/models"
)

func PostsCreate(c *gin.Context) {
	//Get data off request body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body	}
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

func PostsIndex(c *gin.Context){
	//Get Posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//Respond with Posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context){
	//Get ID off url
	id := c.Param("id")

	//Get specific post
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}