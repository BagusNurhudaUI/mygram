package controllers

import "github.com/gin-gonic/gin"

func PostSocialMedias(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PostSocialMedias!",
	})
}

func GetSocialMedias(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome GetSocialMedias!",
	})
}

func PutSocialMediasId(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PutSocialMediasId!",
	})
}

func DeleteSocialMediasId(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome DeleteSocialMediasId!",
	})
}
