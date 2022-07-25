package data

import (
	commentData "socialmediabackendproject/feature/comments/data"
	"socialmediabackendproject/feature/post_images/data"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	User_ID    uint                  `json:"user_id" form:"user_id"`
	Caption    string                `json:"caption" form:"caption"`
	Post_Image []data.Post_Image     `gorm:"foreignKey:Post_ID"`
	Comment    []commentData.Comment `gorm:"foreignKey:Post_ID"`
}
