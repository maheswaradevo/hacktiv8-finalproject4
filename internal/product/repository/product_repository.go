package repository

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
)

type ProductImplRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

func ProvideProductRepository(db *sql.DB, logger *zap.Logger) *ProductImplRepo {
	return &ProductImplRepo{
		db:     db,
		logger: logger,
	}
}

var (
	CREATE_PRODUCT     = "INSERT INTO `products`(category_id, title, price, stock) VALUES (?, ?, ?, ?);"
	CHECK_CATEGORY     = "SELECT id FROM categories WHERE id = ?;"
	VIEW_PRODUCT       = "SELECT p.id, p.title, p.price, p.stock, p.category_id, p.created_at FROM products p ORDER BY p.created_at DESC;"
	COUNT_PRODUCT      = "SELECT COUNT(*) FROM products;"
	CHECK_PRODUCT      = "SELECT id FROM products WHERE id = ?;"
	UPDATE_PRODUCT     = "UPDATE products SET title = ?, price = ?, stock = ?, category_id = ? WHERE id = ?;"
	GET_PRODUCT_BY_ID  = "SELECT p.id, p.title, p.price, p.stock, p.category_id, p.updated_at FROM `products` p WHERE p.id = ?;"
	DELETE_PRODUCT     = "DELETE FROM products WHERE id = ?;"
	FIND_PRODUCT_BY_ID = "SELECT id, title, price, stock, category_id FROM products WHERE id=?;"
)

func (p ProductImplRepo) CreateProduct(ctx context.Context, data model.Product) (productID uint64, err error) {
	query := CREATE_PRODUCT
	stmt, errPrepare := p.db.PrepareContext(ctx, query)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[CreateProduct] failed to prepare statement: %v", zap.Error(errPrepare))
		return
	}
	defer stmt.Close()

	res, errExec := stmt.ExecContext(ctx, data.CategoryID, data.Title, data.Price, data.Stock)
	if errExec != nil {
		p.logger.Sugar().Errorf("[CreateProduct] failed to insert user to the database: %v", zap.Error(errExec))
		return
	}

	id, _ := res.LastInsertId()
	productID = uint64(id)

	return productID, nil
}

func (p ProductImplRepo) CheckCategory(ctx context.Context, categoryID uint64) (bool, error) {
	query := CHECK_CATEGORY
	stmt, errPrepare := p.db.PrepareContext(ctx, query)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[CheckCategory] failed to prepare the statement, err: %v", zap.Error(errPrepare))
		return false, errPrepare
	}
	rows, errQuery := stmt.QueryContext(ctx, categoryID)
	if errQuery != nil {
		p.logger.Sugar().Errorf("[CheckCategory] failed to query to the database, err: %v", zap.Error(errQuery))
		return false, errQuery
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
func (p ProductImplRepo) ViewProduct(ctx context.Context) (model.Products, error) {
	query := VIEW_PRODUCT
	stmt, errPrepare := p.db.PrepareContext(ctx, query)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[ViewTask] failed to prepare the statement, err: %v", zap.Error(errPrepare))
		return nil, errPrepare
	}
	rows, errQuery := stmt.QueryContext(ctx)
	if errQuery != nil {
		p.logger.Sugar().Errorf("[ViewTask] failed to query to the database, err: %v", zap.Error(errQuery))
		return nil, errQuery
	}
	var products model.Products
	for rows.Next() {
		product := model.Product{}
		errScan := rows.Scan(
			&product.ProductID,
			&product.Title,
			&product.Price,
			&product.Stock,
			&product.CategoryID,
			&product.CreatedAt,
		)
		if errScan != nil {
			p.logger.Sugar().Errorf("[ViewTask] failed to scan the data from the database, err: %v", zap.Error(errScan))
			return nil, errScan
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p ProductImplRepo) CountProduct(ctx context.Context) (int, error) {
	query := COUNT_PRODUCT
	rows := p.db.QueryRowContext(ctx, query)
	var count int
	errScan := rows.Scan(&count)
	if errScan != nil {
		p.logger.Sugar().Errorf("[CountProduct] failed to scan the data from the database, err: %v", zap.Error(errScan))
		return 0, errScan
	}
	return count, nil
}

func (p ProductImplRepo) CheckProduct(ctx context.Context, productID uint64) (bool, error) {
	query := CHECK_PRODUCT
	stmt, errPrepare := p.db.PrepareContext(ctx, query)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[CheckProduct] failed to prepare the statement, err: %v", zap.Error(errPrepare))
		return false, errPrepare
	}
	rows, errQuery := stmt.QueryContext(ctx, productID)
	if errQuery != nil {
		p.logger.Sugar().Errorf("[CheckProduct] failed to query to the database, err: %v", zap.Error(errQuery))
		return false, errQuery
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func (p ProductImplRepo) UpdateProduct(ctx context.Context, reqData model.ProductCategoryJoined, productID uint64) error {
	query := UPDATE_PRODUCT

	stmt, errPrepare := p.db.PrepareContext(ctx, query)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[UpdateProduct] failed to prepare the statement, err: %v", zap.Error(errPrepare))
		return errPrepare
	}
	_, errExec := stmt.ExecContext(ctx, reqData.Product.Title, reqData.Product.Price, reqData.Product.Stock, reqData.Product.CategoryID, productID)
	if errExec != nil {
		p.logger.Sugar().Errorf("[UpdateProduct] failed to store data to the database, err: %v", zap.Error(errExec))
		return errExec
	}
	return nil
}

func (p ProductImplRepo) GetProductByID(ctx context.Context, productID uint64) (*dto.EditProductResponse, error) {
	query := GET_PRODUCT_BY_ID
	stmt, errPrepare := p.db.PrepareContext(ctx, query)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[GetProductByID] failed to prepare the statement, err: %v", zap.Error(errPrepare))
		return nil, errPrepare
	}
	rows := stmt.QueryRowContext(ctx, productID)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[GetProductByID] failed to query to the database, err: %v", zap.Error(errPrepare))
		return nil, errPrepare
	}
	product := model.ProductCategoryJoined{}
	errPrepare = rows.Scan(
		&product.Product.ProductID,
		&product.Product.Title,
		&product.Product.Price,
		&product.Product.Stock,
		&product.Product.CategoryID,
		&product.Product.UpdatedAt,
	)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[GetTaskByID] failed to scan the data from the database, err: %v", zap.Error(errPrepare))
		return nil, errPrepare
	}
	return dto.NewEditProductResponse(product.Product), errPrepare
}

func (p ProductImplRepo) DeleteProduct(ctx context.Context, productID uint64) error {
	query := DELETE_PRODUCT

	stmt, errPrepare := p.db.PrepareContext(ctx, query)
	if errPrepare != nil {
		p.logger.Sugar().Errorf("[DeleteTask] failed to prepare the statement, err: %v", zap.Error(errPrepare))
		return errPrepare
	}

	_, errQuery := stmt.QueryContext(ctx, productID)
	if errQuery != nil {
		p.logger.Sugar().Errorf("[DeleteTask] failed to delete the product, err: %v", zap.Error(errQuery))
		return errQuery
	}
	return nil
}

func (p ProductImplRepo) FindProductByID(ctx context.Context, productID uint64) (*model.Product, error) {
	query := FIND_PRODUCT_BY_ID

	rows := p.db.QueryRowContext(ctx, query, productID)

	product := &model.Product{}

	errScanData := rows.Scan(&product.ProductID, &product.Title, &product.Price, &product.Stock, &product.CategoryID)
	if errScanData != nil && errScanData != sql.ErrNoRows {
		p.logger.Sugar().Errorf("[FindProductByID] failed to scan data", zap.Error(errScanData))
		return nil, errScanData
	} else if errScanData == sql.ErrNoRows {
		p.logger.Sugar().Errorf("[FindProductByID] there's no data with id %v", productID)
		return nil, errors.ErrInvalidResources
	}
	return product, nil
}
