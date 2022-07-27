package delivery

import "socialmediabackendproject/domain"

type Register struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   bool   `json:"gender" form:"gender"`
	Bod      string `json:"bod" form:"bod" validate:"required"`
}

func (r *Register) ToDomain() domain.User {
	return domain.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		Gender:   r.Gender,
		Bod:      r.Bod,
	}
}
