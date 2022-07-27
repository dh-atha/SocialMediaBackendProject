package data

import (
	// "errors"

	"socialmediabackendproject/domain"

	"gorm.io/gorm"
)

type postData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.PostData {
	return &postData{
		db: DB,
	}
}

func (pd *postData) GetSpecific(id uint) (domain.Post, error) {
	var data Post
	err := pd.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return domain.Post{}, err
	}
	return data.ToDomain(), nil
}
