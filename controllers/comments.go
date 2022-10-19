package controllers

import (
	"mygram/database"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)
	Comment := models.Comment{}

	c.ShouldBindJSON(&Comment)
	Comment.User_id = userData

	err := db.Debug().Create(&Comment).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot create comments",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(201, gin.H{
			"id":         Comment.ID,
			"message":    Comment.Message,
			"photo_id":   Comment.Photo_id,
			"User_id":    Comment.User_id,
			"created_at": Comment.CreatedAt,
		})
	}

}

func GetComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)
	data := []models.ResponseComment{}
	Comment := []models.Comment{}
	User := models.User{}
	_ = data
	err := db.Debug().Where("user_id", userData).Find(&Comment).Error
	errUser := db.Where("id = ? ", userData).Find(&User).Error
	if err != nil || errUser != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot create comments",
			"message": err.Error(),
		})
		return
	}

	if len(Comment) > 0 {
		for i, comment := range Comment {
			_, _ = i, comment
			tempResp := models.ResponseComment{}
			tempResp.User.ID = User.ID
			tempResp.User.Email = User.Email
			tempResp.User.Username = User.Username
			tempResp.ID = comment.ID
			tempResp.Message = comment.Message
			tempResp.Photo_id = comment.Photo_id
			tempResp.User_id = comment.User_id

			tempPhoto := models.Photo{}
			err := db.Debug().Where("id = ?", comment.Photo_id).Find(&tempPhoto).Error
			if err != nil {
				c.JSON(200, gin.H{
					"message": "any errors",
				})
			}
			tempResp.Photo.ID = tempPhoto.ID
			tempResp.Photo.Title = tempPhoto.Title
			tempResp.Photo.Photo_url = tempPhoto.Photo_url
			tempResp.Photo.User_id = tempPhoto.User_id
			data = append(data, tempResp)
		}
		c.JSON(200, data)
	} else {
		c.JSON(404, gin.H{
			"message": "You dont have any comments yet",
		})
	}

}

func PutCommentsId(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)
	params := c.Param("commentId")
	Comment := models.Comment{}

	c.ShouldBindJSON(&Comment)
	_, _, _ = db, params, userData
	message := Comment.Message
	err := db.Debug().First(&Comment, "id = ?", params).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "cannot update Comment, bad request",
			"message": err.Error(),
		})
		return
	} else {
		Comment.Message = message
		err := db.Debug().Save(&Comment).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Cannot update message",
			})
			return
		}
		c.JSON(200, gin.H{
			"id ":        Comment.ID,
			"message":    Comment.Message,
			"photo_id":   Comment.Photo_id,
			"user_id":    Comment.User_id,
			"updated_at": Comment.UpdatedAt,
		})

	}

}

func DeleteCommentsId(c *gin.Context) {
	db := database.GetDB()
	params := c.Param("commentId")

	err := db.Debug().Where("id = ?", params).Delete(&models.Comment{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot delete the data, bad request",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Your comment has been successfully deleted",
		})
	}
}
