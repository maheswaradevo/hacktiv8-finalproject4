package dto

import "github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"

type (
	UserRegisterRequest struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserSignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (dto *UserRegisterRequest) ToEntity() (u *model.User) {
	u = &model.User{
		FullName: dto.FullName,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     "Customer",
		Balance:  0,
	}
	return
}

func (dto *UserSignInRequest) ToEntity() (u *model.User) {
	u = &model.User{
		Email:    dto.Email,
		Password: dto.Password,
	}
	return
}
