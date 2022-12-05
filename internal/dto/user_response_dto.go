package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type UserRegisterResponse struct {
	UserID    uint64    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserRegisterResponse(u model.User) *UserRegisterResponse {
	return &UserRegisterResponse{
		UserID:    u.UserID,
		FullName:  u.FullName,
		Email:     u.Email,
		Balance:   u.Balance,
		CreatedAt: time.Now(),
	}
}
