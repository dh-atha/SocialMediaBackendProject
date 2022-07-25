package data

import (
	"socialmediabackendproject/domain"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.UserData {
	return &userData{
		db: DB,
	}
}

func (ud *userData) Insert(newUser domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (us *userData) Login(data domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (ud *userData) GetSpecific(id uint) (domain.User, error) {
	return domain.User{}, nil
}

func (ud *userData) GetAll() ([]domain.User, error) {
	return []domain.User{}, nil
}
