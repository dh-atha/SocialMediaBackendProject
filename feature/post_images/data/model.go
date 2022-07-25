package data

import "gorm.io/gorm"

type Post_Image struct {
	gorm.Model
	Post_ID    uint   `json:"post_id" form:"post_id"`
	Image_Path string `json:"image_path" form:"image_path"`
}
