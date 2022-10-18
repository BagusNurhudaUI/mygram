package controllers

import (
	"mygram/database"
	"mygram/models"

	"github.com/gin-gonic/gin"
)

func PostComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(uint)

	Comment := models.Comment{}
	Photo := models.Photo{}
	User := models.User{}
	c.ShouldBindJSON(&Comment)
	_, _ = db, userData

	err := db.Debug().Where("id = ?", userData).Find(&User).Error
	errPhoto := db.Debug().Where("id = ?", Comment.Photo_id).Find(&Photo).Error
	_, _ = err, errPhoto
	Comment.Photo = &Photo
	Comment.User = &User
	c.JSON(200, gin.H{
		"data": Comment,
	})

}

func GetComments(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome GetComments!",
	})
}

func PutCommentsId(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PutCommentsId!",
	})
}

func DeleteCommentsId(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome DeleteCommentsId!",
	})
}
