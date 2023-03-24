package routes

import (
	"challenge-dua/controllers"
	"challenge-dua/repository"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	bookRepo := repository.NewBookRepository()
	bookCtrl := controllers.NewBookController(bookRepo)

	v1 := r.Group("/books")
	{
		v1.GET("/", bookCtrl.GetAllBooks)
		v1.GET("/:id", bookCtrl.GetBookByID)
		v1.POST("/", bookCtrl.AddBook)
		v1.PUT("/:id", bookCtrl.UpdateBook)
		v1.DELETE("/:id", bookCtrl.DeleteBook)
	}

	return r
}
