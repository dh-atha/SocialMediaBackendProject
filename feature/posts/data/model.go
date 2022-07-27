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

func (p *Post) ToDomain() domain.Post {
	return domain.Post{
		ID:         p.ID,
		User_ID:    p.User_ID,
		Caption:    p.Caption,
		Created_At: p.CreatedAt,
		Updated_At: p.UpdatedAt,
	}
}

func ToEntity(data domain.Post) Post {
	return Post{
		User_ID: data.User_ID,
		Caption: data.Caption,
	}
}
