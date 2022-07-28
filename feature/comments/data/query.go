package data

import (
	"errors"
	"socialmediabackendproject/domain"

	"gorm.io/gorm"
)

type commentData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.CommentData {
	return &commentData{
		db: DB,
	}
}

func (cd *commentData) Insert(data domain.Comment) (domain.Comment, error) {
	var commentData Comment = ToEntity(data)
	err := cd.db.Create(&commentData).Error
	if err != nil {
		return domain.Comment{}, err
	}

	return commentData.ToDomain(), nil
}

func (cd *commentData) Delete(data domain.Comment) error {
	var commentData Comment
	err := cd.db.Where("id = ?", data.ID).First(&commentData).Error
	if err != nil {
		return errors.New("comments not found")
	}

	if data.User_ID != commentData.User_ID {
		return errors.New("cant delete, comment not yours")
	}

	err = cd.db.Delete(&commentData).Error
	if err != nil {
		return err
	}

	return nil
}
