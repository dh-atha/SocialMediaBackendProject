package usecase

import "socialmediabackendproject/domain"

type postUsecase struct {
	postData domain.PostData
}

func New(pd domain.PostData) domain.PostUsecase {
	return &postUsecase{
		postData: pd,
	}
}

func (ps *postUsecase) GetAllPosts() ([]domain.Post, []string, [][]string, error) {
	data, username, post_images, err := ps.postData.GetAll()
	return data, username, post_images, err
}

func (ps *postUsecase) AddPost(id uint, data domain.Post) (domain.Post, error) {
	data.User_ID = id
	data, err := ps.postData.Insert(data)
	return data, err
}
