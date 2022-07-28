package usecase

import (
	"socialmediabackendproject/domain"
)

type CommentUsecase struct {
	commentData domain.CommentData
}

func New(cd domain.CommentData) domain.CommentUsecase {
	return &CommentUsecase{
		commentData: cd,
	}
}

func (cu *CommentUsecase) AddComment(data domain.Comment) (domain.Comment, error) {
	data, err := cu.commentData.Insert(data)
	return data, err
}

func (cu *CommentUsecase) DeleteComment(data domain.Comment) error {
	err := cu.commentData.Delete(data)
	return err
}
