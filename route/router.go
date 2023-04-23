package route

import (
	"MyGram/controller"
	"MyGram/middlewares"
	"MyGram/service"

	// _ "ShowCase/docs"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := controller.NewServer(app)
	api := r.Group("/user")
	{
		api.POST("/register", server.UserRegister)
		api.POST("/login", server.UserLogin)
	}
	photo := r.Group("/photo")
	{
		photo.Use(middlewares.Authentification())
		photo.POST("/", server.CreatePhoto)
		photo.GET("/", server.GetAllPhoto)
		photo.GET("/:photoId", middlewares.PhotoAuthorization(), server.GetPhotoById)
		photo.PUT("/:photoId", middlewares.PhotoAuthorization(), server.UpdatePhoto)
		photo.DELETE("/:photoId", middlewares.PhotoAuthorization(), server.DeletePhoto)
	}
	comment := r.Group("/comment")
	{
		comment.Use(middlewares.Authentification())
		comment.POST("/:photoId", server.CreateComment)
		comment.GET("/", server.GetAllComment)
		comment.GET("/:commentId/:photoId", middlewares.CommentAuthorization(), server.GetCommentById)
		comment.PUT("/:commentId/:photoId", middlewares.CommentAuthorization(), server.UpdateComment)
		comment.DELETE("/:commentId/:photoId", middlewares.CommentAuthorization(), server.DeleteComment)
	}
	socialMedia := r.Group("/socialMedia")
	{
		socialMedia.Use(middlewares.Authentification())
		socialMedia.POST("/", server.CreateSocialMedia)
		socialMedia.GET("/", server.GetAllSocialMedia)
		socialMedia.GET("/:socialMediaId", middlewares.SocialMediaAuthorization(), server.GetSocialMediaById)
		socialMedia.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), server.UpdateSocialMedia)
		socialMedia.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), server.DeleteSocialMedia)
	}
}
