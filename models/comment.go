package models

type Comment struct {
	GormModel
	User_id  uint
	User     *User
	Photo_id uint `json:"photo_id"`
	Photo    *Photo
	Message  string `json:"message"`
}
