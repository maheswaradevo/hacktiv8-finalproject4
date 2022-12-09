package transaction

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type TransactionRepository interface {
	FindCategoryByID(ctx context.Context, categoryId uint64) (*model.Category, error)
	UpdateStockBalanceSoldProduct(ctx context.Context, newStock uint64, newBalance int, newSoldProductAmount uint64, productID uint64, userID uint64, categoryID uint64) error
	InsertTransactionHistory(ctx context.Context, data model.TransactionHistory) error
}
