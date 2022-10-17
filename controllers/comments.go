package controllers

import "github.com/gin-gonic/gin"

func PostComments(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PostComments!",
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
