package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/heise/myproject/Desktop/firstapp/internal/book/handler"
)

func RegisterBookRoutes(r *gin.Engine, bookHandler *handler.BookHandler) {
	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("/", bookHandler.AddBook)
		bookRoutes.GET("/", bookHandler.GetBooks)
		bookRoutes.GET("/:id", bookHandler.GetBook)
		bookRoutes.PUT("/:id", bookHandler.UpdateBook)
		bookRoutes.DELETE("/:id", bookHandler.DeleteBook)
	}
}
