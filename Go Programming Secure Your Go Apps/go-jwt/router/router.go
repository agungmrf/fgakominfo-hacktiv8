package router

import (
	"go-jwt/controllers"
	"go-jwt/midlewares"

	"github.com/gin-gonic/gin"
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
		productRouter.DELETE("/:productId", midlewares.ProductAuthorization(), controllers.DeleteProduct)
		productRouter.GET("/:productId", midlewares.ProductAuthorization(), controllers.GetProduct)
		productRouter.GET("/", midlewares.ProductAuthorization(), controllers.GetProducts)
	}

	return r
}
