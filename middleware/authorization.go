package middleware

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AthPhoto() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		if Photo.User_id != userData {
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
		db := database.GetDB()
		params, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(uint)
		Comment := models.Comment{}

		err = db.Debug().First(&Comment, params).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Cannot found data",
				"message": err.Error(),
			})
			return
		}

		if Comment.User_id != userData {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}

func AthSocialMedia() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		params, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(uint)
		SocialMedia := models.SocialMedia{}

		err = db.Debug().First(&SocialMedia, params).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Cannot found data",
				"message": err.Error(),
			})
			return
		}

		if SocialMedia.User_id != userData {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
