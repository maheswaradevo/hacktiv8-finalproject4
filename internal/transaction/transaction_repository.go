package transaction

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type TransactionRepository interface {
	FindCategoryByID(ctx context.Context, categoryId uint64) (*model.Category, error)
	UpdateStockBalanceSoldProduct(ctx context.Context, newStock uint64, newBalance int, newSoldProductAmount uint64, productID uint64, userID uint64, categoryID uint64) error
	InsertTransactionHistory(ctx context.Context, data model.TransactionHistory) error
	ViewMyTransaction(ctx context.Context, userID uint64) (model.TransactionHistories, error)
	CountMyTransaction(ctx context.Context, userID uint64) (int, error)
	ViewUsersTransaction(ctx context.Context) (model.TransactionUsersHistories, error)
	CountTransaction(ctx context.Context) (int, error)
	FindRoleByUserID(ctx context.Context, userID uint64) (*model.User, error)
}
