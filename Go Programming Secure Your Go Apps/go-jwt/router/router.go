package router

import (
	"github.com/gin-gonic/gin"
	"go-jwt/controllers"
	"go-jwt/midlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(midlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productId", midlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
