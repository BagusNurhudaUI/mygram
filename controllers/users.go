package controllers

import "github.com/gin-gonic/gin"

func PostUserRegister(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PostUserResgister!",
	})
}

func PostUserLogin(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PostUserLogin!",
	})
}

func PutUser(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Welcome PutUser!",
	})
}
