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
	VIEW_MY_TRANSACTION        = "SELECT th.id, th.product_id, th.user_id, th.quantity, th.total_price, p.id, p.title, p.price, p.stock, p.category_id, p.created_at, p.updated_at FROM transaction_histories th INNER JOIN products p ON th.product_id = p.id WHERE th.user_id = ? ORDER BY th.id ASC;"
	COUNT_MY_TRANSACTION       = "SELECT COUNT(*) FROM transaction_histories th WHERE user_id  = ?;"
	VIEW_USER_TRANSACTION      = "SELECT th.id, th.product_id, th.user_id, th.quantity, th.total_price, p.id, p.title, p.price, p.stock, p.category_id, p.created_at, p.updated_at, u.id, u.full_name, u.email, u.balance, u.created_at, u.updated_at FROM transaction_histories th INNER JOIN products p ON th.product_id = p.id INNER JOIN users u ON th.user_id = u.id  ORDER BY th.id ASC;"
	COUNT_TRANSACTION          = "SELECT COUNT(*) FROM transaction_histories th;"
	FIND_USER_BY_ROLE          = "SELECT full_name, role FROM users WHERE id=?;"
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

func (tr transactionRepository) ViewMyTransaction(ctx context.Context, userID uint64) (model.TransactionHistories, error) {
	query := VIEW_MY_TRANSACTION

	rows, errQuery := tr.db.QueryContext(ctx, query, userID)
	if errQuery != nil {
		tr.logger.Sugar().Errorf("[ViewMyTransaction] failed to query to the database", zap.Error(errQuery))
		return nil, errQuery
	}
	transactionHistories := model.TransactionHistories{}

	for rows.Next() {
		var transactionProduct model.TransactionProductJoined
		errScan := rows.Scan(
			&transactionProduct.TransactionHistory.TransactionID,
			&transactionProduct.TransactionHistory.ProductID,
			&transactionProduct.TransactionHistory.UserID,
			&transactionProduct.TransactionHistory.Quantity,
			&transactionProduct.TransactionHistory.TotalPrice,
			&transactionProduct.Product.ProductID,
			&transactionProduct.Product.Title,
			&transactionProduct.Product.Price,
			&transactionProduct.Product.Stock,
			&transactionProduct.Product.CategoryID,
			&transactionProduct.Product.CreatedAt,
			&transactionProduct.Product.UpdatedAt,
		)
		if errScan != nil {
			tr.logger.Sugar().Errorf("[ViewMyTransaction] failed to scan from the database", zap.Error(errScan))
			return nil, errScan
		}
		transactionHistories = append(transactionHistories, &transactionProduct)
	}
	return transactionHistories, nil
}

func (tr transactionRepository) CountMyTransaction(ctx context.Context, userID uint64) (int, error) {
	query := COUNT_MY_TRANSACTION

	rows, errQuery := tr.db.QueryContext(ctx, query, userID)
	if errQuery != nil {
		tr.logger.Sugar().Errorf("[CountMyTransaction] failed to query to the database", zap.Error(errQuery))
		return 0, errQuery
	}
	var countTransaction int
	for rows.Next() {
		errScan := rows.Scan(&countTransaction)
		if errScan != nil {
			tr.logger.Sugar().Errorf("[CountMyTransaction] failed to scan from the database", zap.Error(errScan))
			return 0, errScan
		}
	}
	return countTransaction, nil
}

func (tr transactionRepository) ViewUsersTransaction(ctx context.Context) (model.TransactionUsersHistories, error) {
	query := VIEW_USER_TRANSACTION

	rows, errQuery := tr.db.QueryContext(ctx, query)
	if errQuery != nil {
		tr.logger.Sugar().Errorf("[ViewUserTransaction] failed to query to the database", zap.Error(errQuery))
		return nil, errQuery
	}

	usersTransactionHistories := model.TransactionUsersHistories{}

	for rows.Next() {
		var transactionProductUser model.TransactionProductUserJoined
		errScan := rows.Scan(
			&transactionProductUser.TransactionHistory.TransactionID,
			&transactionProductUser.TransactionHistory.ProductID,
			&transactionProductUser.TransactionHistory.UserID,
			&transactionProductUser.TransactionHistory.Quantity,
			&transactionProductUser.TransactionHistory.TotalPrice,
			&transactionProductUser.Product.ProductID,
			&transactionProductUser.Product.Title,
			&transactionProductUser.Product.Price,
			&transactionProductUser.Product.Stock,
			&transactionProductUser.Product.CategoryID,
			&transactionProductUser.Product.CreatedAt,
			&transactionProductUser.Product.UpdatedAt,
			&transactionProductUser.User.UserID,
			&transactionProductUser.User.FullName,
			&transactionProductUser.User.Email,
			&transactionProductUser.User.Balance,
			&transactionProductUser.User.CreatedAt,
			&transactionProductUser.User.UpdatedAt,
		)
		if errScan != nil {
			tr.logger.Sugar().Errorf("[ViewUserTransaction] failed to scan data from the database", zap.Error(errScan))
			return nil, errScan
		}
		usersTransactionHistories = append(usersTransactionHistories, &transactionProductUser)
	}
	return usersTransactionHistories, nil
}

func (tr transactionRepository) CountTransaction(ctx context.Context) (int, error) {
	query := COUNT_TRANSACTION

	rows, errQuery := tr.db.QueryContext(ctx, query)
	if errQuery != nil {
		tr.logger.Sugar().Errorf("[CountTransaction] failed to query to the database", zap.Error(errQuery))
		return 0, errQuery
	}
	var countTransaction int
	for rows.Next() {
		errScan := rows.Scan(&countTransaction)
		if errScan != nil {
			tr.logger.Sugar().Errorf("[CountTransaction] failed to scan from the database", zap.Error(errScan))
			return 0, errScan
		}
	}
	return countTransaction, nil
}

func (tr transactionRepository) FindRoleByUserID(ctx context.Context, userID uint64) (*model.User, error) {
	query := FIND_USER_BY_ROLE

	rows, errQuery := tr.db.QueryContext(ctx, query, userID)
	if errQuery != nil {
		tr.logger.Sugar().Errorf("[FindRoleByUserID] failed to query to the database", zap.Error(errQuery))
		return nil, errQuery
	}

	userInfo := &model.User{}
	for rows.Next() {
		errScan := rows.Scan(&userInfo.FullName, &userInfo.Role)
		if errScan != nil {
			tr.logger.Sugar().Errorf("[FindRoleByUserID] failed to scan the data", zap.Error(errScan))
			return nil, errScan
		}
	}
	return userInfo, nil
}
