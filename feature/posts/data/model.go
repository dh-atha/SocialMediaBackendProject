package data

import (
	"socialmediabackendproject/domain"
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

func ToEntity(data domain.Post) Post {
	return Post{
		User_ID:              data.User_ID,
		Caption:              data.Caption,
	}
}

func (u *Post) ToDomain() domain.Post{
	return domain.Post{
		User_ID:              u.User_ID,
		Caption:              u.Caption,
	}
}

func FromModel(data domain.Post) Post {
	var res Post
	res.ID = uint(data.ID)
	res.User_ID = data.User_ID
	res.Caption = data.Caption
	return res
}
