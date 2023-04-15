package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
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
		photoRouter.GET("/getall", controllers.GetAllPhoto)
		photoRouter.GET("/get/:photoID", controllers.GetPhoto)
		photoRouter.PUT("/update/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/delete/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/create/:photoID", controllers.CreateComment)
		commentRouter.GET("/getall", controllers.GetAllComment)
		commentRouter.GET("/getall/:photoID", controllers.GetAllCommentPhoto)
		commentRouter.GET("/get/:commentID", controllers.GetComment)
		commentRouter.PUT("/update/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/delete/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	sosialMediaRouter := r.Group("/sosialmedia")
	{
		sosialMediaRouter.Use(middlewares.Authentication())
		sosialMediaRouter.POST("/create", controllers.CreateSocialMedia)
		sosialMediaRouter.GET("/getall", controllers.GetAllSocialMedia)
		sosialMediaRouter.GET("/get/:socialMediaID", controllers.GetSocialMedia)
		sosialMediaRouter.PUT("/update/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		sosialMediaRouter.DELETE("/delete/:socialMediaID", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
