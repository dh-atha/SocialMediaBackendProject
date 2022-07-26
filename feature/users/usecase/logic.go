package usecase

import (
	"errors"
	"log"
	"socialmediabackendproject/domain"
	"socialmediabackendproject/feature/common"

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
		log.Println("error encrypt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	data, err := us.userData.Insert(newUser)
	return data, err
}

func (us *userUsecase) Login(data domain.User) (domain.User, string, error) {
	userData, password, err := us.userData.Login(data)
	if err != nil {
		return domain.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(data.Password))
	if err != nil {
		return domain.User{}, "", errors.New("wrong password")
	}

	return userData, common.GenerateToken(int(userData.ID)), nil
}

func (us *userUsecase) GetAllUser() ([]domain.User, error) {
	data, err := us.userData.GetAll()
	return data, err
}

func (us *userUsecase) GetSpecificUser(id uint) (domain.User, error) {
	data, err := us.userData.GetSpecific(id)
	return data, err
}

func (us *userUsecase) UpdateUser(data domain.User, id uint) (domain.User, error) {
	if data.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("error encrypt password", err)
			return domain.User{}, err
		}
		data.Password = string(hashed)
	}
	updateData, err := us.userData.Update(data, id)
	return updateData, err
}

func (us *userUsecase) DeleteUser(id uint) error {
	err := us.userData.Delete(id)
	return err
}

func (us *userUsecase) UpdateProfilePic(data domain.User) error {
	err := us.userData.UpdateProfilePic(data)
	return err
}
