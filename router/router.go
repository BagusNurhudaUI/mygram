package router

import (
	"mygram/controllers"
	"mygram/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to index",
		})
	})
	users := r.Group("/users")
	users.POST("/register", controllers.PostUserRegister)
	users.POST("/login", controllers.PostUserLogin)
	users.PUT("/", middleware.Authentication(), controllers.PutUser)
	users.DELETE("/", middleware.Authentication(), controllers.DeleteUser)

	photos := r.Group("/photos")
	photos.Use(middleware.Authentication())
	photos.POST("/", controllers.PostPhotos)
	photos.GET("/", controllers.GetPhotos)
	photos.PUT("/:photoId", middleware.AthPhoto(), controllers.PutPhotosId)
	photos.DELETE("/:photoId", middleware.AthPhoto(), controllers.DeletePhotosId)

	comments := r.Group("/comments")
	comments.Use(middleware.Authentication())
	comments.POST("/", controllers.PostComments)
	comments.GET("/", controllers.GetComments)
	comments.PUT("/:commentId", middleware.AthComment(), controllers.PutCommentsId)
	comments.DELETE("/:commentId", middleware.AthComment(), controllers.DeleteCommentsId)

	SocialMedias := r.Group("/socialmedias")
	SocialMedias.Use(middleware.Authentication())
	SocialMedias.POST("/", controllers.PostSocialMedias)
	SocialMedias.GET("/", controllers.GetSocialMedias)
	SocialMedias.PUT("/:socialMediaId", middleware.AthSocialMedia(), controllers.PutSocialMediasId)
	SocialMedias.DELETE("/:socialMediaId", middleware.AthSocialMedia(), controllers.DeleteSocialMediasId)
	return r
}
