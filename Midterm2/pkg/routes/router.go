package routes

import (
	"Test2/middleware"
	"Test2/pkg/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router = func(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Posts",
		})
	})
	// book actions
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", middleware.RequireAuth, middleware.UserIsSellerOrAdmin, controllers.CreateBook)
	router.POST("/give_a_rating", middleware.RequireAuth, middleware.UserIsClientOrAdmin, controllers.GiveRating)
	router.GET("/books/:id", controllers.GetBook)
	router.PUT("/books/:id", middleware.RequireAuth, middleware.UserIsSellerOrAdmin, controllers.EditBook)
	router.DELETE("/books/:id", middleware.RequireAuth, middleware.UserIsSellerOrAdmin, controllers.DeleteBook)
	router.GET("/filter_by_prices", controllers.PriceFiltering)
	router.GET("/filter_by_rating", controllers.RatingFiltering)

	//comments
	router.POST("/comment", middleware.RequireAuth, middleware.UserIsClientOrAdmin, controllers.CreateComment)
	router.GET("/comment", middleware.RequireAuth, controllers.GetAllComments)
	//router.GET("/comment/:id", middleware.RequireAuth, controllers.GetCommentsForBook)
	router.DELETE("/comment/:id", middleware.RequireAuth, middleware.UserIsClientOrAdmin, controllers.DeleteComment)

	// authentication
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/logout", middleware.RequireAuth, controllers.SignOut)

	router.POST("/addtoorder", middleware.RequireAuth, middleware.UserIsClientOrAdmin, controllers.Addtoorder)
	router.GET("/myorders", middleware.RequireAuth, middleware.UserIsClientOrAdmin, controllers.Viewmyorders)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
