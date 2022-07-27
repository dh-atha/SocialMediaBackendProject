package delivery

import "socialmediabackendproject/domain"

type GetUser struct {
	ID                   uint
	Name                 string
	Email                string
	Profile_picture_path string
}

func ToGetUser(data domain.User) GetUser {
	return GetUser{
		ID:                   data.ID,
		Name:                 data.Name,
		Email:                data.Email,
		Profile_picture_path: data.Profile_picture_path,
	}
}

type GetSpecificUser struct {
	ID                   uint
	Name                 string
	Email                string
	Profile_picture_path string
	Gender               bool
	Bod                  string
}

func ToGetSpecificUser(data domain.User) GetSpecificUser {
	return GetSpecificUser{
		ID:                   data.ID,
		Name:                 data.Name,
		Email:                data.Email,
		Profile_picture_path: data.Profile_picture_path,
		Gender:               data.Gender,
		Bod:                  data.Bod,
	}
}

type DeletedUser struct {
	ID                   uint
	Name                 string
	Email                string
	Profile_picture_path string
	Gender               bool
	Bod                  string
}

func ToDeletedUser(data domain.User) DeletedUser {
	return DeletedUser{
		ID:                   data.ID,
		Name:                 data.Name,
		Email:                data.Email,
		Profile_picture_path: data.Profile_picture_path,
		Gender:               data.Gender,
		Bod:                  data.Bod,
	}
}

