package domain

type User struct {
	ID                   uint
	Name                 string `json:"name" form:"name"`
	Email                string `json:"email" form:"email"`
	Password             string `json:"password" form:"password"`
	Profile_picture_path string
	Gender               bool   `json:"gender" form:"gender"`
	Bod                  string `json:"bod" form:"bod"`
}

type UserUsecase interface {
	Register(newUser User) (User, error)
	Login(data User) (User, string, error)
	GetAllUser() ([]User, error)
	GetSpecificUser(id uint) (User, error)
	UpdateUser(data User, id uint) (User, error)
	DeleteUser(id uint) error
}

type UserData interface {
	Insert(newUser User) (User, error)
	Login(data User) (User, string, error)
	GetAll() ([]User, error)
	GetSpecific(id uint) (User, error)
	Update(data User, id uint) (User, error)
	Delete(id uint) error
}
