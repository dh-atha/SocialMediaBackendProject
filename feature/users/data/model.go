package data

import (
	commentData "socialmediabackendproject/feature/comments/data"
	"socialmediabackendproject/feature/posts/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                 string `json:"name" form:"name"`
	Email                string `json:"email" form:"email"`
	Password             string `json:"password" form:"password"`
	Profile_picture_path string
	Address              string                `json:"address" form:"address"`
	Gender               bool                  `json:"gender" form:"gender"`
	Bod                  string                `json:"bod" form:"bod"`
	Post                 []data.Post           `gorm:"foreignKey:User_ID"`
	Comment              []commentData.Comment `gorm:"foreignKey:User_ID"`
}
