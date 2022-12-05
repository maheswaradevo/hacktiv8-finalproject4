package auth

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
)

type UserService interface {
	Register(ctx context.Context, data *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)
}
