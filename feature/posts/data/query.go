package data

import (
	"errors"
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

func (pd *postData) GetAll() ([]domain.Post, []string, [][]string, error) {
	var data []Post
	pd.db.Find(&data)
	if len(data) < 1 {
		return []domain.Post{}, []string{}, [][]string{}, errors.New("no data found")
	}

	var domainData []domain.Post
	var username []string
	var postimages [][]string
	for i := 0; i < len(data); i++ {
		domainData = append(domainData, data[i].ToDomain())
		var tmpusername string
		pd.db.Raw("SELECT name FROM users WHERE id = ?", data[i].User_ID).Scan(&tmpusername)
		username = append(username, tmpusername)
		var tmpimages []string
		pd.db.Raw("SELECT image_path FROM post_images WHERE post_id = ?", data[i].ID).Scan(&tmpimages)
		postimages = append(postimages, tmpimages)
	}

	return domainData, username, postimages, nil
}

func (pd *postData) Insert(data domain.Post) (domain.Post, error) {
	var postData Post = ToEntity(data)
	err := pd.db.Create(&postData).Error
	if err != nil {
		return domain.Post{}, err
	}

	return postData.ToDomain(), nil
}

// func (pd *postData) GetPostsByID(id uint) ([])
