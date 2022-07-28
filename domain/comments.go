package domain

import "time"

type Comment struct {
	ID         uint
	User_ID    uint
	Post_ID    uint
	Caption    string `json:"caption" form:"caption"`
	Created_At time.Time
}

type CommentUsecase interface {
	AddComment(data Comment) (Comment, error)
	DeleteComment(data Comment) error
}

type CommentData interface {
	Insert(data Comment) (Comment, error)
	Delete(data Comment) error
}
