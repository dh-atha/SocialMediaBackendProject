package data

import (
	"errors"
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
	var newData User = ToEntity(newUser)
	res := ud.db.Where("email = ?", newData.Email).Find(&User{})
	if res.RowsAffected > 0 {
		return domain.User{}, errors.New("email registered")
	}

	err := ud.db.Create(&newData).Error
	if err != nil {
		return domain.User{}, err
	}

	return newData.ToDomain(), nil
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
