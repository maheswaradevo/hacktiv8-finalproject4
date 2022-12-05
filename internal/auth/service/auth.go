package service

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo   auth.User
	logger *zap.Logger
}

func NewAuthService(repo auth.User, logger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (auth *Service) Register(ctx context.Context, data *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
	userInfo := data.ToEntity()

	exists, err := auth.repo.FindByEmail(ctx, userInfo.Email)
	if err != nil && err != errors.ErrInvalidResources {
		auth.logger.Sugar().Errorf("[Register] failed to fetch the data, err: %v", err)
		return nil, err
	}
	if exists != nil {
		err = errors.ErrUserExists
		auth.logger.Sugar().Errorf("[Register] user with email %v already exist", userInfo.Email)
		return nil, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		auth.logger.Sugar().Errorf("[Register] failed to create hashed password, err: %v", err)
		return nil, err
	}
	userInfo.Password = string(hashed)

	userID, err := auth.repo.InsertUser(ctx, *userInfo)
	if err != nil {
		auth.logger.Sugar().Errorf("[Register] failed to save user to the db, err: %v", err)
		return nil, err
	}
	userInfo.UserID = userID
	return dto.NewUserRegisterResponse(*userInfo), nil
}
