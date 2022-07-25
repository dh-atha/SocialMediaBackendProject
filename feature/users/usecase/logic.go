package usecase

import (
	"log"
	"socialmediabackendproject/domain"

	"golang.org/x/crypto/bcrypt"
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
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrpt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	data, err := us.userData.Insert(newUser)
	return data, err
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
