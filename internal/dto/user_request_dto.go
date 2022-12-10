package dto

import "github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"

type (
	UserRegisterRequest struct {
		FullName string `json:"full_name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required,min=6"`
	}

	UserSignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserTopupBalanceRequest struct {
		Balance int `json:"balance"`
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

func (dto *UserTopupBalanceRequest) ToEntity() (u *model.User) {
	u = &model.User{
		Balance: dto.Balance,
	}
	return
}
