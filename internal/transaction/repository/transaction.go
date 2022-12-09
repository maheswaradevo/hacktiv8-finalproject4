package repository

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
)

type transactionRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewTransactionRepository(db *sql.DB, logger *zap.Logger) *transactionRepository {
	return &transactionRepository{
		db:     db,
		logger: logger,
	}
}

var (
	FIND_CATEGORY_BY_ID        = "SELECT id, type, sold_product_amount FROM categories WHERE id=?;"
	UPDATE_PRODUCT             = "UPDATE products SET stock=? WHERE id=?;"
	UPDATE_BALANCE             = "UPDATE users SET balance = ? WHERE id = ?;"
	UPDATE_SOLD_PRODUCT        = "UPDATE categories SET sold_product_amount=? WHERE id=?;"
	INSERT_TRANSACTION_HISTORY = "INSERT INTO transaction_histories(product_id, user_id, quantity, total_price) VALUES (?, ?, ?, ?)"
)

func (tr transactionRepository) FindCategoryByID(ctx context.Context, categoryId uint64) (*model.Category, error) {
	query := FIND_CATEGORY_BY_ID

	rows := tr.db.QueryRowContext(ctx, query, categoryId)

	category := &model.Category{}

	errScanData := rows.Scan(&category.CategoryID, &category.Type, &category.SoldProductAmount)
	if errScanData != nil && errScanData != sql.ErrNoRows {
		tr.logger.Sugar().Errorf("[FindCategoryByID] failed to scan data", zap.Error(errScanData))
		return nil, errScanData
	} else if errScanData == sql.ErrNoRows {
		tr.logger.Sugar().Errorf("[FindCategoryByID] there's no data with id %v", categoryId)
		return nil, errors.ErrInvalidResources
	}
	return category, nil
}

func (tr transactionRepository) UpdateStockBalanceSoldProduct(ctx context.Context, newStock uint64,
	newBalance int, newSoldProductAmount uint64, productID uint64,
	userID uint64, categoryID uint64) error {
	tx, errTransaction := tr.db.BeginTx(ctx, nil)
	if errTransaction != nil {
		tr.logger.Sugar().Errorf("[UpdateStockBalanceSoldProduct] failed to begin transaction", zap.Error(errTransaction))
		return errTransaction
	}
	defer tx.Rollback()

	updateProduct := UPDATE_PRODUCT
	updateBalance := UPDATE_BALANCE
	updateSoldProduct := UPDATE_SOLD_PRODUCT

	_, errExecProduct := tx.ExecContext(ctx, updateProduct, newStock, productID)
	if errExecProduct != nil {
		tr.logger.Sugar().Errorf("[UpdateStockBalanceSoldProduct] failed to update product stock", zap.Error(errExecProduct))
		return errExecProduct
	}

	_, errUpdateBalance := tx.ExecContext(ctx, updateBalance, newBalance, userID)
	if errUpdateBalance != nil {
		tr.logger.Sugar().Errorf("[UpdateStockBalanceSoldProduct] failed to update user balance", zap.Error(errUpdateBalance))
		return errUpdateBalance
	}
	_, errUpdateSoldProduct := tx.ExecContext(ctx, updateSoldProduct, newSoldProductAmount, categoryID)
	if errUpdateSoldProduct != nil {
		tr.logger.Sugar().Errorf("[UpdateStockBalanceSoldProduct] failed to update stock product amount", zap.Error(errUpdateSoldProduct))
		return errUpdateSoldProduct
	}
	if errCommit := tx.Commit(); errCommit != nil {
		tr.logger.Sugar().Errorf("[UpdateStockBalanceSoldProduct] failed to commit transaction", zap.Error(errCommit))
		return errCommit
	}
	return nil
}

func (tr transactionRepository) InsertTransactionHistory(ctx context.Context, data model.TransactionHistory) error {
	query := INSERT_TRANSACTION_HISTORY

	_, errExec := tr.db.ExecContext(ctx, query, data.ProductID, data.UserID, data.Quantity, data.TotalPrice)
	if errExec != nil {
		tr.logger.Sugar().Errorf("[InsertTransactionHistory] failed to insert the data", zap.Error(errExec))
		return errExec
	}
	return nil
}
