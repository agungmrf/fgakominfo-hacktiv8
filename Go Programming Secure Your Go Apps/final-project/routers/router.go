package routers

import (
	"final-project/controllers"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4/middleware"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/create", controllers.CreatePhoto)
		photoRouter.GET("/getall", controllers.GetallPhoto)
		photoRouter.GET("/get/:photoID", controllers.GetPhoto)
		photoRouter.PUT("/update/:photoID", middlewares
	}
}
