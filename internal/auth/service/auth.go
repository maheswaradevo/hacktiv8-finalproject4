package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo   auth.UserRepository
	logger *zap.Logger
}

func NewAuthService(repo auth.UserRepository, logger *zap.Logger) *service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (auth *service) Register(ctx context.Context, data *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
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

func (auth *service) Login(ctx context.Context, data *dto.UserSignInRequest) (*dto.UserSignInResponse, error) {
	userLogin := data.ToEntity()

	userCred, errFindByEmail := auth.repo.FindByEmail(ctx, userLogin.Email)
	if errFindByEmail != nil {
		auth.logger.Sugar().Errorf("[Login] failed to fetch data by email, err: %v", zap.Error(errFindByEmail))
		return nil, errFindByEmail
	}
	errMismatchPassword := bcrypt.CompareHashAndPassword([]byte(userCred.Password), []byte(userLogin.Password))
	if errMismatchPassword != nil {
		errMismatchPassword = errors.ErrMismatchedHashAndPassword
		auth.logger.Sugar().Errorf("[Login] wrong password")
		return nil, errMismatchPassword
	}
	token, errCreateAccessToken := auth.createAccessToken(userCred)
	if errCreateAccessToken != nil {
		auth.logger.Sugar().Errorf("[Login] failed to create access token, err: %v", zap.Error(errCreateAccessToken))
		return nil, errCreateAccessToken
	}
	return dto.NewUserSignInResponse(token), nil
}

func (auth *service) TopupBalance(ctx context.Context, data *dto.UserTopupBalanceRequest, userId uint64) (*dto.UserTopupBalanceResponse, error) {
	userBalance := data.ToEntity()

	userData, errFindUserByID := auth.repo.FindUserByID(ctx, userId)
	if errFindUserByID != nil && errFindUserByID != errors.ErrInvalidResources {
		auth.logger.Sugar().Errorf("[TopupBalance] failed to fetch data by id, err: %v", zap.Error(errFindUserByID))
		return nil, errFindUserByID
	}
	if userData == nil {
		errUserNotFound := errors.ErrDataNotFound
		auth.logger.Sugar().Errorf("[TopupBalance] data user with id %v not found", userId)
		return nil, errUserNotFound
	}
	newBalance := userData.Balance + userBalance.Balance

	rowsAffect, errUpdateBalance := auth.repo.UpdateBalance(ctx, newBalance, userId)
	if errUpdateBalance != nil {
		auth.logger.Sugar().Errorf("[TopupBalance] failed to update balance, err: %v", zap.Error(errUpdateBalance))
		return nil, errUpdateBalance
	}
	if rowsAffect < 0 {
		errUpdateBalance = errors.ErrTopupBalance
		auth.logger.Sugar().Errorf("[TopupBalance] there's no rows to be updated")
		return nil, errUpdateBalance
	}
	msg := fmt.Sprintf("Your balance has been successfully updated to %v", newBalance)
	return dto.NewUserTopupBalanceResponse(msg), nil
}

func (auth *service) createAccessToken(user *model.User) (string, error) {
	cfg := config.GetConfig()

	claims := jwt.MapClaims{
		"authorized":  true,
		"exp":         time.Now().Add(time.Hour * 8).Unix(),
		"userId":      user.UserID,
		"userRole":    user.Role,
		"userBalance": user.Balance,
	}
	token := jwt.NewWithClaims(cfg.JWTSigningMethod, claims)
	signedToken, errSignedString := token.SignedString([]byte(cfg.ApiSecretKey))
	if errSignedString != nil {
		auth.logger.Sugar().Errorf("failed to create new token, err: %v", zap.Error(errSignedString))
		return "", errSignedString
	}
	return signedToken, nil
}
