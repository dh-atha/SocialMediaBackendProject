package domain

type Post struct {
	ID         uint
	User_ID    uint
	Caption    string
	Created_At string
	Updated_At string
}

type PostUsecase interface {
	GetSpecificPost(id uint) (Post, error)
	UpdatePost(id uint, updatePost Post) (Post, error)
}

type PostData interface {
	GetSpecific(id uint) (Post, error)
	UpdatePost(id uint, updatePost Post) (Post, error)
}
