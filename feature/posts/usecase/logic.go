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

func (ps *postUsecase) GetAllPosts() ([]domain.Post, []domain.User, [][]string, error) {
	data, userdata, post_images, err := ps.postData.GetAll()
	return data, userdata, post_images, err
}

func (ps *postUsecase) AddPost(id uint, data domain.Post) (domain.Post, error) {
	data.User_ID = id
	data, err := ps.postData.Insert(data)
	return data, err
}

func (ps *postUsecase) GetMyPosts(id uint) ([]domain.Post, domain.User, [][]string, error) {
	posts, userdata, postimages, err := ps.postData.GetAllPostsByID(id)
	return posts, userdata, postimages, err
}

func (ps *postUsecase) GetSpecificPost(id uint) (domain.Post, domain.User, []string, []domain.Comment, []domain.User, error) {
	post, userdata, postimages, comments, commentUserData, err := ps.postData.GetPostByID(id)
	return post, userdata, postimages, comments, commentUserData, err
}

func (ps *postUsecase) UpdatePost(id uint, updateData domain.Post) (domain.Post, error) {
	data, err := ps.postData.Update(id, updateData)
	return data, err
}
