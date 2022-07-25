package usecase

import (
	"socialmediabackendproject/domain"
)

type userUsecase struct {
	userData domain.UserData
}

func New(ud domain.UserData) domain.UserUsecase {
	return &userUsecase{
		userData: ud,
	}
}

func (us *userUsecase) Register(newUser domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (us *userUsecase) Login(data domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (us *userUsecase) GetAllUser() ([]domain.User, error) {
	return []domain.User{}, nil
}

func (us *userUsecase) GetSpecificUser(id uint) (domain.User, error) {
	return domain.User{}, nil
}
