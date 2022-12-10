package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/product"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/transaction"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/constants"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
)

type service struct {
	transactionRepo transaction.TransactionRepository
	productRepo     product.ProductRepository
	userRepo        auth.UserRepository
	logger          *zap.Logger
}

func NewTransactionService(transactionRepo transaction.TransactionRepository, productRepo product.ProductRepository, userRepo auth.UserRepository, logger *zap.Logger) *service {
	return &service{
		transactionRepo: transactionRepo,
		productRepo:     productRepo,
		userRepo:        userRepo,
		logger:          logger,
	}
}

func (tr *service) DoTransaction(ctx context.Context, data *dto.DoTransactionRequest, userID uint64) (dto.DoTransactionResponse, error) {
	transactionInfo := data.ToEntity()
	exists, errCheckProduct := tr.productRepo.CheckProduct(ctx, transactionInfo.ProductID)
	if errCheckProduct != nil {
		tr.logger.Sugar().Errorf("[DoTransaction] failed to check product", zap.Error(errCheckProduct))
		return dto.DoTransactionResponse{}, errCheckProduct
	}
	if !exists {
		errDataNotFound := errors.ErrDataNotFound
		tr.logger.Sugar().Errorf("[DoTransaction] product with id %v not found", transactionInfo.ProductID)
		return dto.DoTransactionResponse{}, errDataNotFound
	}

	productInfo, errFindProduct := tr.productRepo.FindProductByID(ctx, transactionInfo.ProductID)
	if errFindProduct != nil && errFindProduct == errors.ErrInvalidResources {
		tr.logger.Sugar().Errorf("[DoTransaction] failed to fetch product data", zap.Error(errFindProduct))
		return dto.DoTransactionResponse{}, errFindProduct
	}

	if productInfo.Stock < transactionInfo.Quantity {
		errStockNotFound := errors.ErrStockNotFound
		tr.logger.Sugar().Errorf("[DoTransaction] there's not enough stock", zap.Error(errStockNotFound))
		return dto.DoTransactionResponse{}, errStockNotFound
	}
	userInfo, errFindUser := tr.userRepo.FindUserByID(ctx, userID)
	if errFindUser != nil && errFindUser == errors.ErrInvalidResources {
		errUserNotFound := errors.ErrDataNotFound
		tr.logger.Sugar().Errorf("[DoTransaction] there's no user with id %v", userID)
		return dto.DoTransactionResponse{}, errUserNotFound
	}
	totalPrice := transactionInfo.Quantity * productInfo.Price
	if userInfo.Balance < int(totalPrice) {
		errBalance := errors.ErrBalance
		tr.logger.Sugar().Errorf("[DoTransaction] not enough balance on account", zap.Error(errBalance))
		return dto.DoTransactionResponse{}, errBalance
	}
	categoryInfo, errFindCategory := tr.transactionRepo.FindCategoryByID(ctx, productInfo.CategoryID)
	if errFindCategory != nil && errFindCategory == errors.ErrInvalidResources {
		errCategoryNotFound := errors.ErrDataNotFound
		tr.logger.Sugar().Errorf("[DoTransaction] there's no data with id %v", productInfo.CategoryID)
		return dto.DoTransactionResponse{}, errCategoryNotFound
	}

	newStock := productInfo.Stock - transactionInfo.Quantity
	newBalance := userInfo.Balance - int(totalPrice)
	newSoldProductAmount := categoryInfo.SoldProductAmount + transactionInfo.Quantity

	errUpdate := tr.transactionRepo.UpdateStockBalanceSoldProduct(ctx, newStock, newBalance, newSoldProductAmount, transactionInfo.ProductID, userID, productInfo.CategoryID)
	if errUpdate != nil {
		tr.logger.Sugar().Errorf("[DoTransaction] failed to update changes", zap.Error(errUpdate))
		return dto.DoTransactionResponse{}, errUpdate
	}

	errInsertHistory := tr.transactionRepo.InsertTransactionHistory(ctx, model.TransactionHistory{
		ProductID:  productInfo.ProductID,
		UserID:     userID,
		Quantity:   transactionInfo.Quantity,
		TotalPrice: totalPrice,
	})
	if errInsertHistory != nil {
		tr.logger.Sugar().Errorf("[DoTransaction] failed to insert transaction history", zap.Error(errInsertHistory))
		return dto.DoTransactionResponse{}, errInsertHistory
	}
	msg := "You have successfully purchased the product"
	var transactionResponse = dto.NewDoTransactionResponse(msg, model.TransactionProductJoined{
		TransactionHistory: model.TransactionHistory{
			TotalPrice: totalPrice,
			Quantity:   transactionInfo.Quantity,
		},
		Product: model.Product{
			Title: productInfo.Title,
		},
	})
	return transactionResponse, nil
}

func (tr *service) ViewMyTransaction(ctx context.Context, userID uint64) ([]dto.ViewMyTransactionResponse, error) {
	countTransaction, errCount := tr.transactionRepo.CountMyTransaction(ctx, userID)
	if errCount != nil {
		tr.logger.Sugar().Errorf("[ViewMyTransaction] failed to count transaction", zap.Error(errCount))
		return nil, errCount
	}
	if countTransaction == 0 {
		errDataNotFound := errors.ErrDataNotFound
		tr.logger.Sugar().Errorf("[ViewMyTransaction] no transaction history data", zap.Error(errDataNotFound))
		return nil, errDataNotFound
	}
	userInfo, errGetRole := tr.transactionRepo.FindRoleByUserID(ctx, userID)
	if errGetRole != nil {
		tr.logger.Sugar().Errorf("[ViewUsersTransaction] failed to find role", zap.Error(errGetRole))
		return nil, errGetRole
	}
	fmt.Printf("userInfo.Role: %v\n", userInfo.Role)
	if strings.ToLower(userInfo.Role) != constants.CustomerRole {
		errOnlyAdmin := errors.ErrOnlyAdmin
		tr.logger.Sugar().Errorf("[ViewUsersTransaction] only admin can acces", zap.Error(errOnlyAdmin))
		return nil, errOnlyAdmin
	}
	transactions, errViewTransaction := tr.transactionRepo.ViewMyTransaction(ctx, userID)
	if errViewTransaction != nil {
		tr.logger.Sugar().Errorf("[ViewMyTransaction] failed to view user's transaction", zap.Error(errViewTransaction))
		return nil, errViewTransaction
	}
	transactionResponse := dto.NewViewMyTransactionsResponse(transactions)
	return transactionResponse, nil
}

func (tr *service) ViewUserTransaction(ctx context.Context, userID uint64) ([]dto.ViewUserTransactionResponse, error) {
	countTransaction, errCount := tr.transactionRepo.CountTransaction(ctx)
	if errCount != nil {
		tr.logger.Sugar().Errorf("[ViewUserTransaction] failed to count the transaction", zap.Error(errCount))
		return nil, errCount
	}
	if countTransaction == 0 {
		errDataNotFound := errors.ErrDataNotFound
		tr.logger.Sugar().Errorf("[ViewUserTransaction] no transaction history data", zap.Error(errDataNotFound))
		return nil, errDataNotFound
	}
	userInfo, errGetRole := tr.transactionRepo.FindRoleByUserID(ctx, userID)
	if errGetRole != nil {
		tr.logger.Sugar().Errorf("[ViewUsersTransaction] failed to find role", zap.Error(errGetRole))
		return nil, errGetRole
	}
	fmt.Printf("userInfo.Role: %v\n", userInfo.Role)
	if strings.ToLower(userInfo.Role) != constants.AdminRole {
		errOnlyAdmin := errors.ErrOnlyAdmin
		tr.logger.Sugar().Errorf("[ViewUsersTransaction] only admin can acces", zap.Error(errOnlyAdmin))
		return nil, errOnlyAdmin
	}
	transactions, errViewTransaction := tr.transactionRepo.ViewUsersTransaction(ctx)
	if errViewTransaction != nil {
		tr.logger.Sugar().Errorf("[ViewUserTransaction] failed to view users transaction", zap.Error(errViewTransaction))
		return nil, errViewTransaction
	}
	transactionUsersResponse := dto.NewViewUsersTransactionsResponse(transactions)
	return transactionUsersResponse, nil
}
