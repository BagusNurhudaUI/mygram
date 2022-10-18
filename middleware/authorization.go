package middleware

import (
	"fmt"
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AthPhoto() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hello this is from the middleware!")

		db := database.GetDB()
		params, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(uint)
		Photo := models.Photo{}

		err = db.Debug().First(&Photo, params).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Cannot found data",
				"message": err.Error(),
			})
			return
		}

		if Photo.UserID != userData {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func AthComment() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func AthSocialMedia() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
