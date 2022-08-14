package controller

import (
	"net/http"

	"github.com/Praveenkusuluri08/helpers"
	"github.com/Praveenkusuluri08/models"
	"github.com/gin-gonic/gin"
)

func TestRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Messagee": "Hello world!",
	})
}

func CreatePost(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	post := models.Posts{Title: body.Title, Body: body.Body}
	result := helpers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	var posts []models.Posts
	helpers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetSinglePost(c *gin.Context) {
	queryString := c.Query("id")

	var post []models.Posts
	helpers.DB.Where("id = ?", queryString).Find(&post)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetSinglePostParams(c *gin.Context) {
	id := c.Param("id")
	var post models.Posts

	helpers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdateSinglePost(c *gin.Context) {
	id := c.Query("id")

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	var post models.Posts

	helpers.DB.First(&post, id)

	helpers.DB.Model(&post).Updates(models.Posts{Title: body.Title, Body: body.Body})

	c.JSON(http.StatusOK, gin.H{"post": post, "message": "Post updated successfully"})

}

func DeleteSinglePost(c *gin.Context) {
	id := c.Query("id")
	var posts models.Posts
	helpers.DB.Where("id = ?", id).Delete(&posts)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

//TODO:This type of endpoint is mainly for permanently delete data from db
func DeletePostPermanently(c *gin.Context) {
	id := c.Query("id")
	helpers.DB.Unscoped().Delete(&models.Posts{}, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted permanently",
	})
}
