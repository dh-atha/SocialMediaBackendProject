package usecase

import(
	"socialmediabackendproject/domain"
	// "socialmediabackendproject/feature/common"
)


type postUsecase struct {
	postData domain.PostData
}


func New(up domain.PostData) domain.PostUsecase {
	return &postUsecase{
		postData: up,
	}
}


func (us *postUsecase) GetSpecificPost(id uint) (domain.Post, error) {
	data, err := us.postData.GetSpecific(id)
	return data, err
}

func (us *postUsecase) UpdatePost(id uint ,data domain.Post) (domain.Post, error) {
	// data, err := us.postData.UpdatePost(uint(id))
	return domain.Post{}, nil
}
