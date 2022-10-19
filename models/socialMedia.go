package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name             string `json:"name" gorm:"not null" valid:"required~Name is required"`
	Social_media_url string `json:"social_media_url" gorm:"not null" valid:"required~URL is required"`
	User_id          uint
	User             *User `gorm:"foreignKey:User_id"`
}

type ResponseSocialMedia struct {
	ID               uint      `json:"id"`
	CreatedAt        time.Time `json:"created_at,omitempty" `
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	Name             string    `json:"name" `
	Social_media_url string    `json:"social_media_url" `
	User_id          uint      `json:"user_id"`
	User             struct {
		ID                uint   `json:"id"`
		Username          string `json:"username" `
		Profile_image_url string `json:"profile_image_url"`
	} `json:"user"`
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	return
}

func (u *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
