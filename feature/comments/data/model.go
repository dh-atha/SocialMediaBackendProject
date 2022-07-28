package data

import (
	"socialmediabackendproject/domain"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User_ID uint   `json:"user_id" form:"user_id"`
	Post_ID uint   `json:"post_id" form:"post_id"`
	Caption string `json:"caption" form:"caption"`
}

func ToEntity(data domain.Comment) Comment {
	return Comment{
		Post_ID: data.Post_ID,
		User_ID: data.User_ID,
		Caption: data.Caption,
	}
}

func (c *Comment) ToDomain() domain.Comment {
	return domain.Comment{
		ID:         c.ID,
		User_ID:    c.User_ID,
		Post_ID:    c.Post_ID,
		Caption:    c.Caption,
		Created_At: c.CreatedAt,
	}
}
