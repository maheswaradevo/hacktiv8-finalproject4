package auth

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, data model.User) (userID uint64, err error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindUserByID(ctx context.Context, userId uint64) (*model.User, error)
	UpdateBalance(ctx context.Context, balance int, userId uint64) (int64, error)
}
