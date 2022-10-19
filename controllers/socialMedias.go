package controllers

import (
	"mygram/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSocialMedias(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)
	SocialMedia := models.SocialMedia{}
	_, _ = db, userData
	c.ShouldBindJSON(&SocialMedia)
	SocialMedia.User_id = userData

	err := db.Debug().Create(&SocialMedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot create social media object, bad request",
			"message": err.Error(),
		})
	} else {
		c.JSON(201, gin.H{
			"id":               SocialMedia.ID,
			"name":             SocialMedia.Name,
			"social_media_url": SocialMedia.Social_media_url,
			"User_id":          SocialMedia.User_id,
			"created_at":       SocialMedia.CreatedAt,
		})
	}

}

func GetSocialMedias(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)
	SocialMedia := []models.SocialMedia{}
	// SocialMedia.User_id = userData
	data := []models.ResponseSocialMedia{}
	_, _ = userData, data
	err := db.Debug().Model(&SocialMedia).Preload("User").Where("user_id = ?", userData).Find(&SocialMedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot getsocial media object, bad request",
			"message": err.Error(),
		})
	} else {
		for _, sm := range SocialMedia {
			temp := models.ResponseSocialMedia{}
			temp.ID = sm.ID
			temp.Name = sm.Name
			temp.Social_media_url = sm.Social_media_url
			temp.User_id = sm.User_id
			temp.CreatedAt = *sm.CreatedAt
			temp.UpdatedAt = *sm.UpdatedAt
			temp.User.ID = sm.User.ID
			temp.User.Username = sm.User.Username
			temp.User.Profile_image_url = "this is links"
			data = append(data, temp)
		}
		c.JSON(200, gin.H{
			"social_medias": data,
		})
	}

}

func PutSocialMediasId(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)
	params := c.Param("socialMediaId")
	SocialMedia := models.SocialMedia{}
	c.ShouldBindJSON(&SocialMedia)
	name := SocialMedia.Name
	url := SocialMedia.Social_media_url
	_, _, _ = userData, name, url
	err := db.Debug().Where("id = ?", params).Find(&SocialMedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot getsocial media object, bad request",
			"message": err.Error(),
		})
		return
	} else {
		SocialMedia.Name = name
		SocialMedia.Social_media_url = url
		err := db.Debug().Save(&SocialMedia).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Cannot getsocial media object, bad request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"id":               SocialMedia.ID,
			"name":             SocialMedia.Name,
			"social_media_url": SocialMedia.Social_media_url,
			"user_id":          SocialMedia.User_id,
			"updated_at":       *SocialMedia.UpdatedAt,
		})
	}
}

func DeleteSocialMediasId(c *gin.Context) {

	db := database.GetDB()
	params := c.Param("socialMediaId")

	err := db.Debug().Where("id = ?", params).Delete(&models.SocialMedia{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot delete the data, bad request",
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Your social media has been successfully deleted",
		})
	}
}
