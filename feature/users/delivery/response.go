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
