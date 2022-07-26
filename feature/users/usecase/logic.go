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

// masih ada yg eror keburu abis billing wkwkwkw
// ===============================
func (us *userUsecase) UpdateUser(id uint, updateUser domain.User) (domain.User, error) {
	new_id := uint(id)
	if id != new_id {
		return domain.User{}, errors.New("invalid user")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
	// err = bcrypt.CompareHashAndPassword([]byte(upassword), []byte(data.Password))

	if err != nil {
		log.Println("error encrpt password", err)
		return domain.User{}, err
	}

	UpdateUser.Password = string(hashed)
	res := us.userData.UpdateUser(id,domain.User{})

	if res.ID == 0 {
		return domain.User{}, errors.New("error update user")
	}

	return res, nil
}
