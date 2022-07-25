package domain

type User struct {
	ID                   uint
	Name                 string `json:"name" form:"name"`
	Email                string `json:"email" form:"email"`
	Password             string `json:"password" form:"password"`
	Profile_picture_path string
	Address              string `json:"address" form:"address"`
	Gender               bool   `json:"gender" form:"gender"`
	Bod                  string `json:"bod" form:"bod"`
}
