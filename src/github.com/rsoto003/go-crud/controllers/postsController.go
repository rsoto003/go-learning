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

func PostsUpdate(c *gin.Context){
	//get the id off the url
	id := c.Param("id")

	//get the data off req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	//find post we are updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})
	//respond with updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context){
	//get id off the url
	id := c.Param("id")

	//find post we want to delete
	var post models.Post
	initializers.DB.First(&post, id)

	//delete post
	initializers.DB.Model(&post).Delete(&post)

	//respond with message stating post was deleted
	c.JSON(200, gin.H{
		"message": "ID of Post deleted: %id",
	})
}