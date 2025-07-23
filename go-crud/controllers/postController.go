package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get Data from Body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a Post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get Data from Body
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Get Data from Body
	var post models.Post
	initializers.DB.First(&post, id)

	// Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Get the data off URL
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find post
	var post models.Post
	initializers.DB.First(&post, id)

	// Update
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	//Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Return
	c.Status(200)
}
