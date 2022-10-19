package controllers

import (
	"mygram/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostPhotos(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData")
	Photo := models.Photo{}
	_ = db
	c.ShouldBindJSON(&Photo)
	Photo.User_id = userData.(uint)

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "cannot create photo, bad request",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(201, gin.H{
			"id":         Photo.User_id,
			"title":      Photo.Title,
			"caption":    Photo.Caption,
			"photo_url":  Photo.Photo_url,
			"user_id":    Photo.User_id,
			"created_at": Photo.CreatedAt,
		})
	}
}

func GetPhotos(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)
	Photo := []models.Photo{}
	User := models.User{}
	data := []models.ResponseGetPhoto{}
	_ = db

	errPhoto := db.Where("user_id = ?", userData).Find(&Photo).Error
	errUser := db.Where("id = ? ", userData).Find(&User).Error

	if errPhoto != nil || errUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "cannot create photo, bad request",
			"message": errPhoto.Error(),
		})
		return
	} else {
		if len(Photo) > 0 {
			for _, photo := range Photo {
				tempResp := models.ResponseGetPhoto{}
				tempResp.ID = photo.ID
				tempResp.Title = photo.Title
				tempResp.Caption = photo.Caption
				tempResp.Photo_url = photo.Photo_url
				tempResp.CreatedAt = photo.CreatedAt
				tempResp.UpdatedAt = photo.UpdatedAt
				tempResp.User.Username = User.Username
				tempResp.User.Email = User.Email
				data = append(data, tempResp)
			}
			c.JSON(200, data)
		} else {
			c.JSON(404, gin.H{
				"message": "You dont have any photos available",
			})
		}

	}

}

func PutPhotosId(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)
	params := c.Param("photoId")
	Photo := models.Photo{}
	UpdatedPhoto := models.Photo{}
	c.ShouldBindJSON(&Photo)
	Photo.User_id = userData

	err := db.Debug().First(&UpdatedPhoto, "id = ?", params).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "cannot update photo, bad request",
			"message": err.Error(),
		})
	} else {
		UpdatedPhoto.Title = Photo.Title
		UpdatedPhoto.Caption = Photo.Caption
		UpdatedPhoto.Photo_url = Photo.Photo_url
		err := db.Debug().Save(&UpdatedPhoto).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "cannot update photo, bad request",
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"id":         UpdatedPhoto.ID,
				"title":      UpdatedPhoto.Title,
				"caption":    UpdatedPhoto.Caption,
				"photo_url":  UpdatedPhoto.Photo_url,
				"user_id":    UpdatedPhoto.User_id,
				"updated_at": UpdatedPhoto.UpdatedAt,
			})
		}

	}

}

func DeletePhotosId(c *gin.Context) {
	db := database.GetDB()
	params := c.Param("photoId")

	err := db.Debug().Where("id = ?", params).Delete(&models.Photo{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot delete the data, bad request",
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Your photo has been successfully deleted",
		})
	}

}
