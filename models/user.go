package models

import (
	"errors"
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null; uniqueIndex" json:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null; uniqueIndex" json:"email" valid:"required~Email is required, email~Email is not validate as email"`
	Password string `gorm:"not null" json:"password" valid:"required~Password is required, minstringlength(6)~Minimum password length is 6 digits"`
	Age      uint   `gorm:"not null json:"age" valid:"required~Age is required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	if u.Age <= 8 {
		err = errors.New("Age must be greater than or equal to 8")
		return
	}

	u.Password = helpers.HashPassword(u.Password)

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errUpdate := govalidator.ValidateStruct(u)
	if errUpdate != nil {
		err = errUpdate
		return
	}
	if u.Age <= 8 {
		err = errors.New("Age must be greater than or equal to 8")
		return
	}
	// u.Password = helpers.HashPassword(u.Password)

	return
}
