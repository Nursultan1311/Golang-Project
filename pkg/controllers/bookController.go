package controllers

import (
	"Test2/initializers"
	"Test2/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	initializers.GetDB().Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"Books": books,
	})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var target models.Book
	initializers.GetDB().Find(&target, "id=?", id)
	c.JSON(http.StatusOK, gin.H{
		"Book": target,
	})
}

func CreateBook(c *gin.Context) {
	book := models.Book{}
	if c.Bind(&book) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	result := initializers.GetDB().Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a new book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func EditBook(c *gin.Context) {
	id := c.Param("id")
	book := models.Book{}
	if c.Bind(&book) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var target models.Book
	initializers.GetDB().Find(&target, "id=?", id)
	if book.Title != target.Title {
		initializers.GetDB().Model(&target).Update("Title", book.Title)
	}
	if book.Description != target.Description {
		initializers.GetDB().Model(&target).Update("Description", book.Description)
	}
	if book.Author != target.Author {
		initializers.GetDB().Model(&target).Update("Author", book.Author)
	}
	if book.Rating != target.Rating {
		initializers.GetDB().Model(&target).Update("Rating", book.Rating)
	}
	if book.Price != target.Price {
		initializers.GetDB().Model(&target).Update("Price", book.Price)
	}
	if book.Language != target.Language {
		initializers.GetDB().Model(&target).Update("Language", book.Language)
	}
	if book.Image != target.Image {
		initializers.GetDB().Model(&target).Update("Image", book.Image)
	}
	if book.Category != target.Category {
		initializers.GetDB().Model(&target).Update("Category", book.Category)
	}
	c.JSON(http.StatusOK, gin.H{
		"Edit": "Book successfully edited",
	})

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var target models.Book
	initializers.GetDB().Delete(&target, "id=?", id)
	c.JSON(http.StatusOK, gin.H{
		"Book": "removed",
	})
}

func PriceFiltering(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	var books []models.Book
	result := initializers.GetDB().Where("price >= ?", from).Where("price <= ?", to).Order("price").Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "failed to get books between of prices",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": books,
	})
}

func RatingFiltering(c *gin.Context) {
	rating := c.Query("rating")
	var books []models.Book
	result := initializers.GetDB().Where("rating >= ?", rating).Order("rating desc").Find(&books)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "failed to get books by rating",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": books,
	})
}

func GiveRating(c *gin.Context) {
	var body struct {
		Title  string
		Rating float32
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var target models.Book
	initializers.GetDB().Find(&target, "title=?", body.Title)
	avg := (body.Rating + target.Rating) / 2
	if body.Rating <= 5.0 {
		initializers.GetDB().Model(&target).Update("Rating", avg)
	}
	c.JSON(http.StatusOK, gin.H{
		"Change book rating": avg,
	})
}
