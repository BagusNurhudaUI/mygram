package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type ResponseGetPhoto struct {
	ID        uint   `json:"id,omitempty"`
	Title     string `json:"title"`
	Caption   string `json:"caption,omitempty"`
	Photo_url string `json:"photo_url"`
	User      struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
	CreatedAt *time.Time `json:"created_at,omitempty" `
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Photo struct {
	GormModel
	Title     string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url" valid:"required~Photo URL is required"`
	User_id   uint
	User      *User `json:"omitempty"  gorm:"foreignKey:User_id "`
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	return
}

func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {

	// if u.Caption == "" {
	// 	err = errors.New("Photo has no caption")
	// }
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
