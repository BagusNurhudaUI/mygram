package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	User_id  uint
	User     *User  `gorm:"foreignKey:User_id"`
	Photo_id uint   `json:"photo_id" valid:"required~Photo id is required"`
	Photo    *Photo `gorm:"foreignKey:Photo_id "`
	Message  string `json:"message" valid:"required~Message is required"`
}

type ResponseComment struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	Photo_id  uint      `json:"photo_id"`
	User_id   uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	User      struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
	Photo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Caption   string `json:"caption,omitempty"`
		Photo_url string `json:"photo_url"`
		User_id   uint   `json:"user_id"`
	} `json:"photo"`
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	return
}

func (u *Comment) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
