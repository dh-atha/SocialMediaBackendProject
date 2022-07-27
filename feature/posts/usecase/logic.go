package usecase

import(
	"socialmediabackendproject/domain"
	// "socialmediabackendproject/feature/common"
)


type postUsecase struct {
	postData domain.PostData
}


func (us *postUsecase) GetSpecificPost(id uint) (domain.Post, error) {
	data, err := us.postData.GetSpecific(id)
	return data, err
}
