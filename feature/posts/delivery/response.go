package delivery

import "socialmediabackendproject/domain"

type GetSpecificPost struct {
	ID         uint
	User_ID    uint
	Caption    string
	Created_At string
	Updated_At string
}

func ToGetSpecificPost(data domain.Post) GetSpecificPost {
	return GetSpecificPost{
		ID:      data.ID,
		User_ID: data.User_ID,
		Caption: data.Caption,
	}
}
