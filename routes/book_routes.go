package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group("/api/books/")
	bookRoutes.Use(middleware.AuthMiddleware())
	{
		bookRoutes.GET("/", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
	}
}
