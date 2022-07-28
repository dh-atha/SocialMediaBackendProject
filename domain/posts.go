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
	AddPostImages(post []string, postID uint) error
	GetMyPosts(id uint) ([]Post, User, [][]string, error)
	GetSpecificPost(id uint) (Post, User, []string, []Comment, []User, error)
	UpdatePost(id uint, updateData Post) (Post, error)
	DeletePost(id uint, userID uint) error
}

type PostData interface {
	GetAll() ([]Post, []User, [][]string, error)
	Insert(data Post) (Post, error)
	InsertPostImages(post []string, postID uint) error
	GetAllPostsByID(id uint) ([]Post, User, [][]string, error)
	GetPostByID(id uint) (Post, User, []string, []Comment, []User, error)
	Update(id uint, updatePost Post) (Post, error)
	Delete(id uint, userID uint) error
}
