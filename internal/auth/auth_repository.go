package auth

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type User interface {
	InsertUser(ctx context.Context, data model.User) (userID uint64, err error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
