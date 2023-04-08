package routes

import (
	"Test2/middleware"
	"Test2/pkg/controllers"
	gin "github.com/gin-gonic/gin"
)

var Router = func(router *gin.Engine) {
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.POST("/give_a_rating", controllers.GiveRating)
	router.GET("/books/:id", controllers.GetBook)
	router.PUT("/books/:id", controllers.EditBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	router.GET("/filter_by_prices", controllers.PriceFiltering)
	router.GET("/filter_by_rating", controllers.RatingFiltering)
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.SignOut)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	err := router.Run()
	if err != nil {
		return
	}
}
