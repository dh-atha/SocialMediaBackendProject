package data

import (
	"socialmediabackendproject/domain"
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
	Gender               bool                  `json:"gender" form:"gender"`
	Bod                  string                `json:"bod" form:"bod"`
	Post                 []data.Post           `gorm:"foreignKey:User_ID"`
	Comment              []commentData.Comment `gorm:"foreignKey:User_ID"`
}

func ToEntity(data domain.User) User {
	return User{
		Name:                 data.Name,
		Email:                data.Email,
		Password:             data.Password,
		Profile_picture_path: data.Profile_picture_path,
		Gender:               data.Gender,
		Bod:                  data.Bod,
	}
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		ID:                   u.ID,
		Name:                 u.Name,
		Email:                u.Email,
		Password:             u.Password,
		Profile_picture_path: u.Profile_picture_path,
		Gender:               u.Gender,
		Bod:                  u.Bod,
	}
}
