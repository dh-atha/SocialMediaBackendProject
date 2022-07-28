package domain

import "time"

type Post struct {
	ID          uint
	User_ID     uint
	Caption     string `json:"caption" form:"caption"`
	Created_At  time.Time
	Updated_At  time.Time
	Post_images []string `json:"post_images" form:"post_images"`
}

type PostUsecase interface {
	GetAllPosts() ([]Post, []User, [][]string, error)
	AddPost(id uint, data Post) (Post, error)
	GetMyPosts(id uint) ([]Post, User, [][]string, error)
	GetSpecificPost(id uint) (Post, User, []string, []Comment, []User, error)
	UpdatePost(id uint, updateData Post) (Post, error)
}

type PostData interface {
	GetAll() ([]Post, []User, [][]string, error)
	Insert(data Post) (Post, error)
	GetAllPostsByID(id uint) ([]Post, User, [][]string, error)
	GetPostByID(id uint) (Post, User, []string, []Comment, []User, error)
	Update(id uint, updatePost Post) (Post, error)
}
