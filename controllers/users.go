package controllers

import (
	"encoding/json"
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// var appJSON = "application/json"

func PostUserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := c.Request.Header.Get("Content-Type")
	User := models.User{}
	_, _ = db, contentType
	c.ShouldBindJSON(&User)
	fmt.Println(json.Marshal(&User))
	err := db.Debug().Create(&User).Error

	if err != nil {
		if strings.Contains(fmt.Sprint(err), "duplicate") {
			err = fmt.Errorf("username or email is already used")
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot create user, bad request",
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(201, gin.H{
			"age":      User.Age,
			"email":    User.Email,
			"id":       User.ID,
			"username": User.Username,
		})
	}
}

func PostUserLogin(c *gin.Context) {
	db := database.GetDB()
	User := models.User{}

	c.ShouldBindJSON(&User)
	password := User.Password
	_, _ = db, password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Email or Password is invalid",
		})
		return
	}

	comparePassword := helpers.ComparePassword(password, User.Password)
	if !comparePassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Email or Password is invalid",
		})
		return
	}

	token := helpers.CreateToken(User.ID)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func PutUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData")
	contentType := helpers.GetContentType(c)
	User := models.User{}
	UserDB := models.User{}
	_, _ = db, contentType
	c.ShouldBindJSON(&User)
	User.ID = userData.(uint)
	err1 := db.Debug().Where("id = ?", User.ID).Take(&UserDB).Error
	if err1 != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Cannot find user",
		})
		return
	}

	if User.Password != "" {
		UserDB.Password = helpers.HashPassword(User.Password)
	}

	if User.Age != 0 {
		UserDB.Age = User.Age
	}

	UserDB.Email = User.Email
	UserDB.Username = User.Username

	err := db.Model(&UserDB).Where("id = ?", User.ID).Updates(&UserDB).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot update the data, bad request",
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"id":         UserDB.ID,
			"email":      UserDB.Email,
			"username":   UserDB.Username,
			"age":        UserDB.Age,
			"updated_at": UserDB.UpdatedAt,
		})
	}

}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData")
	contentType := helpers.GetContentType(c)
	User := models.User{}
	_, _ = db, contentType
	userId := userData.(uint)
	User.ID = userId

	err1 := db.Debug().Where("id = ?", userId).Take(&User).Error
	if err1 != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Cannot find data, maybe has been deleted",
		})
		return
	}

	db.Debug().Where("User_id = ?", userId).Delete(&models.Photo{})
	db.Debug().Where("User_id = ?", userId).Delete(&models.Comment{})
	db.Debug().Where("User_id = ?", userId).Delete(&models.SocialMedia{})
	err := db.Debug().Select("Comment", "Photo", "SocialMedia").Delete(&User).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Cannot delete the data, bad request",
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Your account has been successfully deleted",
			"id":      userId,
		})
	}

}
