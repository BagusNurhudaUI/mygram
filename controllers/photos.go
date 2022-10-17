package controllers

import "github.com/gin-gonic/gin"

func PostPhotos(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PostPhotos!",
	})
}

func GetPhotos(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome GetPhotos!",
	})
}

func PutPhotosId(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PutPhotosId!",
	})
}

func DeletePhotosId(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome DeletePhotosId!",
	})
}
