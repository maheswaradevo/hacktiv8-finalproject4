package repository

import (
	"context"
	"database/sql"
	"log"

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
	CREATE_PRODUCT = "INSERT INTO `products`(category_id, title, price, stock) VALUES (?, ?, ?, ?);"
	CHECK_CATEGORY = "SELECT id FROM categories;"
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
