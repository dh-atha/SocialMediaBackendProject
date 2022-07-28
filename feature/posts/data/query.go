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

func (pd *postData) GetAll() ([]domain.Post, []domain.User, [][]string, error) {
	var data []Post
	pd.db.Find(&data)
	if len(data) < 1 {
		return []domain.Post{}, []domain.User{}, [][]string{}, errors.New("no data found")
	}

	var domainData []domain.Post
	var userdata []domain.User
	var postimages [][]string
	for i := 0; i < len(data); i++ {
		domainData = append(domainData, data[i].ToDomain())
		var tmpuserdata domain.User
		pd.db.Raw("SELECT name, profile_picture_path FROM users WHERE id = ?", data[i].User_ID).Scan(&tmpuserdata)
		userdata = append(userdata, tmpuserdata)
		var tmpimages []string
		pd.db.Raw("SELECT image_path FROM post_images WHERE post_id = ?", data[i].ID).Scan(&tmpimages)
		postimages = append(postimages, tmpimages)
	}

	return domainData, userdata, postimages, nil
}

func (pd *postData) Insert(data domain.Post) (domain.Post, error) {
	var postData Post = ToEntity(data)
	err := pd.db.Create(&postData).Error
	if err != nil {
		return domain.Post{}, err
	}

	return postData.ToDomain(), nil
}

func (pd *postData) GetAllPostsByID(id uint) ([]domain.Post, domain.User, [][]string, error) {
	var postData []Post
	pd.db.Where("user_id", id).Find(&postData)
	if len(postData) < 1 {
		return []domain.Post{}, domain.User{}, [][]string{}, errors.New("no postData found")
	}

	var userData domain.User
	pd.db.Raw("SELECT name, profile_picture_path FROM users WHERE id = ?", id).Scan(&userData)

	var postConvertToDomain []domain.Post
	var postimages [][]string
	for i := 0; i < len(postData); i++ {
		postConvertToDomain = append(postConvertToDomain, postData[i].ToDomain())
		var tmpimages []string
		pd.db.Raw("SELECT image_path FROM post_images WHERE post_id = ?", postData[i].ID).Scan(&tmpimages)
		postimages = append(postimages, tmpimages)
	}

	return postConvertToDomain, userData, postimages, nil
}

func (pd *postData) GetPostByID(id uint) (domain.Post, domain.User, []string, []domain.Comment, []domain.User, error) {
	var postData Post
	err := pd.db.Where("id = ?", id).First(&postData).Error
	if err != nil {
		return domain.Post{}, domain.User{}, []string{}, []domain.Comment{}, []domain.User{}, err
	}

	var userData domain.User
	pd.db.Raw("SELECT name, profile_picture_path FROM users WHERE id = ?", postData.User_ID).Scan(&userData)

	var postimages []string
	pd.db.Raw("SELECT image_path FROM post_images WHERE post_id = ?", id).Scan(&postimages)

	var comments []domain.Comment
	pd.db.Raw("SELECT id, user_id, caption, created_at FROM comments WHERE post_id = ? AND deleted_at is NULL", id).Scan(&comments)

	var commentUserData []domain.User
	for i := 0; i < len(comments); i++ {
		var tmpCommentUserData domain.User
		pd.db.Raw("SELECT name, profile_picture_path FROM users WHERE id = ?", comments[i].User_ID).Scan(&tmpCommentUserData)
		commentUserData = append(commentUserData, tmpCommentUserData)
	}

	return postData.ToDomain(), userData, postimages, comments, commentUserData, nil
}

func (pd *postData) Update(id uint, updateData domain.Post) (domain.Post, error) {
	var currentData Post
	err := pd.db.Where("id = ?", id).First(&currentData).Error
	if err != nil {
		return domain.Post{}, err
	}

	if currentData.User_ID != updateData.User_ID {
		return domain.Post{}, errors.New("post not yours")
	}

	currentData.Caption = updateData.Caption
	err = pd.db.Save(&currentData).Error
	if err != nil {
		return domain.Post{}, errors.New("error update data")
	}

	return currentData.ToDomain(), nil
}

func (pd *postData) Delete(id uint, userID uint) error {
	var postData Post
	err := pd.db.Where("id = ?", id).First(&postData).Error
	if err != nil {
		return err
	}

	if postData.User_ID != userID {
		return errors.New("post not yours")
	}

	err = pd.db.Delete(&postData).Error
	if err != nil {
		return err
	}

	pd.db.Exec("UPDATE comments SET deleted_at = now() WHERE post_id = ?", id)
	return nil
}
