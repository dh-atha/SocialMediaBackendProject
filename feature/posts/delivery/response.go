package delivery

import "time"

type GetAllPost struct {
	ID          uint
	User_ID     uint
	Username    string
	Caption     string
	Created_At  time.Time
	Updated_At  time.Time
	Post_Images []string
}
