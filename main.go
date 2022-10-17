package main

import (
	"fmt"
	"mygram/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting...")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to index",
		})
	})
	users := r.Group("/users")
	users.POST("/register", controllers.PostUserRegister)
	users.POST("/login", controllers.PostUserLogin)
	users.PUT("/", controllers.PutUser)
	users.DELETE("/", controllers.PutUser)

	photos := r.Group("/photos")
	photos.POST("/", controllers.PostPhotos)
	photos.GET("/", controllers.GetPhotos)
	photos.PUT("/:photoId", controllers.PutPhotosId)
	photos.DELETE("/:photoId", controllers.DeletePhotosId)

	comments := r.Group("/comments")
	comments.POST("/", controllers.PostComments)
	comments.GET("/", controllers.GetComments)
	comments.PUT("/:commentId", controllers.PutCommentsId)
	comments.DELETE("/:commentId", controllers.DeleteCommentsId)

	SocialMedias := r.Group("/socialmedias")
	SocialMedias.POST("/", controllers.PostSocialMedias)
	SocialMedias.GET("/", controllers.GetSocialMedias)
	SocialMedias.PUT("/:socialMediasId", controllers.PutSocialMediasId)
	SocialMedias.DELETE("/:socialMediasId", controllers.DeleteSocialMediasId)

	r.Run("127.0.0.1:3000")
}
