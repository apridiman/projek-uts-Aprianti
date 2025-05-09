package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", Register)
		auth.POST("/login", Login)
	}

	book := r.Group("/books")
	{
		book.GET("/", ListBooks)
		book.POST("/", CreateBook)
		book.PUT("/:id", UpdateBook)
		book.DELETE("/:id", DeleteBook)
	}

	borrow := r.Group("/borrowings")
	{
		borrow.POST("/borrow", BorrowBook)
		borrow.POST("/return", ReturnBook)
		borrow.GET("/", ListBorrowings)
	}
}
