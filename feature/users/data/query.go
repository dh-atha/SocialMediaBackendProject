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

func (ud *userData) Login(data domain.User) (domain.User, string, error) {
	var loginData User
	err := ud.db.Where("email = ?", data.Email).First(&loginData).Error
	if err != nil {
		return domain.User{}, "", errors.New(data.Email + " not registered")
	}

	return loginData.ToDomain(), loginData.Password, nil
}

func (ud *userData) GetAll() ([]domain.User, error) {
	var allData []User
	ud.db.Find(&allData)
	if len(allData) < 1 {
		return []domain.User{}, errors.New("no data found")
	}

	var convertToDomain []domain.User
	for i := 0; i < len(allData); i++ {
		convertToDomain = append(convertToDomain, allData[i].ToDomain())
	}

	return convertToDomain, nil
}

func (ud *userData) GetSpecific(id uint) (domain.User, error) {
	var data User
	err := ud.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return domain.User{}, err
	}

	return data.ToDomain(), nil
}

func (ud *userData) Update(updateData domain.User, id uint) (domain.User, error) {
	var currentData User
	err := ud.db.Model(&User{}).Where("id = ?", id).Updates(updateData).Error
	if err != nil {
		return domain.User{}, err
	}

	ud.db.Where("id = ?", id).First(&currentData)
	return currentData.ToDomain(), nil
}

func (ud *userData) Delete(id uint) error {
	var userData User
	err := ud.db.Where("id = ?", id).First(&userData).Error
	if err != nil {
		return err
	}

	ud.db.Exec("update posts set deleted_at = now() where user_id = ?;", id)
	ud.db.Exec("update comments set deleted_at = now() where user_id = ?", id)

	err = ud.db.Delete(&userData).Error
	if err != nil {
		return err
	}

	return nil
}

func (ud *userData) UpdateProfilePic(data domain.User) error {
	err := ud.db.Save(&data).Error
	if err != nil {
		return err
	}
	return nil
}
