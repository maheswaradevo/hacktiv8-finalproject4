package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type ProductImplRepo struct {
	db *sql.DB
}

func ProvideProductRepository(db *sql.DB) *ProductImplRepo {
	return &ProductImplRepo{
		db: db,
	}
}

var (
	CREATE_PRODUCT    = "INSERT INTO `products`(category_id, title, price, stock) VALUES (?, ?, ?, ?);"
	CHECK_CATEGORY    = "SELECT id FROM categories WHERE id = ?;"
	VIEW_PRODUCT      = "SELECT p.id, p.title, p.price, p.stock, p.category_id, p.created_at FROM products p ORDER BY p.created_at DESC;"
	COUNT_PRODUCT     = "SELECT COUNT(*) FROM products;"
	CHECK_PRODUCT     = "SELECT id FROM products WHERE id = ?;"
	UPDATE_PRODUCT    = "UPDATE products SET title = ?, price = ?, stock = ?, category_id = ? WHERE id = ?;"
	GET_PRODUCT_BY_ID = "SELECT p.id, p.title, p.price, p.stock, p.category_id, p.updated_at FROM `products` p WHERE p.id = ?;"
	DELETE_PRODUCT    = "DELETE FROM products WHERE id = ?;"
)

func (p ProductImplRepo) CreateProduct(ctx context.Context, data model.Product) (productID uint64, err error) {
	query := CREATE_PRODUCT
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CreateProduct] failed to prepare statement: %v", err)
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, data.CategoryID, data.Title, data.Price, data.Stock)
	if err != nil {
		log.Printf("[CreateProduct] failed to insert user to the database: %v", err)
		return
	}

	id, _ := res.LastInsertId()
	productID = uint64(id)

	return productID, nil
}

func (p ProductImplRepo) CheckCategory(ctx context.Context, categoryID uint64) (bool, error) {
	query := CHECK_CATEGORY
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CheckCategory] failed to prepare the statement, err: %v", err)
		return false, err
	}
	rows, err := stmt.QueryContext(ctx, categoryID)
	if err != nil {
		log.Printf("[CheckCategory] failed to query to the database, err: %v", err)
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
func (p ProductImplRepo) ViewProduct(ctx context.Context) (model.Products, error) {
	query := VIEW_PRODUCT
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[ViewTask] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[ViewTask] failed to query to the database, err: %v", err)
		return nil, err
	}
	var products model.Products
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(
			&product.ProductID,
			&product.Title,
			&product.Price,
			&product.Stock,
			&product.CategoryID,
			&product.CreatedAt,
		)
		if err != nil {
			log.Printf("[ViewTask] failed to scan the data from the database, err: %v", err)
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p ProductImplRepo) CountProduct(ctx context.Context) (int, error) {
	query := COUNT_PRODUCT
	rows := p.db.QueryRowContext(ctx, query)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		log.Printf("[CountProduct] failed to scan the data from the database, err: %v", err)
		return 0, err
	}
	return count, nil
}

func (p ProductImplRepo) CheckProduct(ctx context.Context, productID uint64) (bool, error) {
	query := CHECK_PRODUCT
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CheckProduct] failed to prepare the statement, err: %v", err)
		return false, err
	}
	rows, err := stmt.QueryContext(ctx, productID)
	if err != nil {
		log.Printf("[CheckProduct] failed to query to the database, err: %v", err)
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func (p ProductImplRepo) UpdateProduct(ctx context.Context, reqData model.ProductCategoryJoined, productID uint64) error {
	query := UPDATE_PRODUCT

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[UpdateProduct] failed to prepare the statement, err: %v", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, reqData.Product.Title, reqData.Product.Price, reqData.Product.Stock, reqData.Product.CategoryID, productID)
	if err != nil {
		log.Printf("[UpdateProduct] failed to store data to the database, err: %v", err)
		return err
	}
	return nil
}

func (p ProductImplRepo) GetProductByID(ctx context.Context, productID uint64) (*dto.EditProductResponse, error) {
	query := GET_PRODUCT_BY_ID
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[GetProductByID] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows := stmt.QueryRowContext(ctx, productID)
	if err != nil {
		log.Printf("[GetProductByID] failed to query to the database, err: %v", err)
		return nil, err
	}
	product := model.ProductCategoryJoined{}
	err = rows.Scan(
		&product.Product.ProductID,
		&product.Product.Title,
		&product.Product.Price,
		&product.Product.Stock,
		&product.Product.CategoryID,
		&product.Product.UpdatedAt,
	)
	if err != nil {
		log.Printf("[GetTaskByID] failed to scan the data from the database, err: %v", err)
		return nil, err
	}
	return dto.NewEditProductResponse(product.Product), err
}

func (p ProductImplRepo) DeleteProduct(ctx context.Context, productID uint64) error {
	query := DELETE_PRODUCT

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[DeleteTask] failed to prepare the statement, err: %v", err)
		return err
	}

	_, err = stmt.QueryContext(ctx, productID)
	if err != nil {
		log.Printf("[DeleteTask] failed to delete the product, err: %v", err)
		return err
	}
	return nil
}
