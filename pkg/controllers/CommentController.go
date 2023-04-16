package controllers

import (
	"Test2/initializers"
	"Test2/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var body struct {
		Title   string
		Comment string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title or comment is wrong",
		})
		return
	}

	comment := models.Comment{Author: GetUserEmail(c), Title: body.Comment, Book: body.Title}
	result := initializers.GetDB().Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create new comment",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"New comment": comment,
	})
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var targetComment models.Comment
	initializers.GetDB().Find(&targetComment, "id=?", id)
	var admin models.User
	initializers.GetDB().Find(&admin, "email=?", GetUserEmail(c))

	if GetUserEmail(c) != targetComment.Author && admin.Type != "Admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Only owners can delete his comments",
		})
		return
	}

	initializers.GetDB().Delete(&targetComment, "id=?", id)
	c.JSON(http.StatusOK, gin.H{
		"Comment": "Successfully removed",
	})
}

func GetAllComments(c *gin.Context) {
	var comments []models.Comment
	initializers.GetDB().Find(&comments)
	c.JSON(http.StatusOK, gin.H{
		"Comments": comments,
	})
}
