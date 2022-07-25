package data

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	User_ID uint   `json:"user_id" form:"user_id"`
	Post_ID uint   `json:"post_id" form:"post_id"`
	Caption string `json:"caption" form:"caption"`
}
