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

type UserUsecase interface {
	Register(newUser User) (User, error)
	Login(data User) (User, error)
	GetAllUser() ([]User, error)
	GetSpecificUser(id uint) (User, error)
}

type UserData interface {
	Insert(newUser User) (User, error)
	Login(data User) (User, error)
	GetAll() ([]User, error)
	GetSpecific(id uint) (User, error)
}
