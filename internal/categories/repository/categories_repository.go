package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type CategoriesImplRepo struct {
	db *sql.DB
}

func ProvideCategoriesRepository(db *sql.DB) *CategoriesImplRepo {
	return &CategoriesImplRepo{
		db: db,
	}
}

var (
	CREATE_CATEGORIES    = "INSERT INTO categories(type) VALUES (?);"
	VIEW_CATEGORIES      = "SELECT c.id, c.type, c.sold_product_amount, c.updated_at, c.created_at, p.id, p.title, p.price, p.stock, p.created_at, p.updated_at FROM `product` p JOIN  `categories` c ON c.id = p.category_id ORDER BY c.created_at DESC;"
	COUNT_CATEGORIES     = "SELECT COUNT(*) FROM categories;"
	UPDATE_CATEGORIES    = "UPDATE categories SET type = ? WHERE id = ?;"
	CHECK_CATEGORY       = "SELECT id FROM categories WHERE id = ?;"
	DELETE_CATEGORIES    = "DELETE FROM categories WHERE id = ?;"
	GET_CATEGORIES_BY_ID = "SELECT c.id, c.type, c.updated_at FROM categories c WHERE c.id = ?;"
)

type CreateCategoriesResponse struct {
	CategoryID        uint64    `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint64    `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

func (ctg CategoriesImplRepo) CreateCategories(ctx context.Context, data model.Category) (categoryID uint64, err error) {
	query := CREATE_CATEGORIES
	stmt, err := ctg.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CreateTask] failed to prepare statement: %v", err)
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, data.Type)
	if err != nil {
		log.Printf("[CreateComment] failed to insert user to the database: %v", err)
		return
	}

	id, _ := res.LastInsertId()
	categoryID = uint64(id)

	return categoryID, nil
}

func (ctg CategoriesImplRepo) ViewCategories(ctx context.Context) (model.CategoriesJoined, error) {
	query := VIEW_CATEGORIES
	stmt, err := ctg.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[ViewCategory] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[ViewCategory] failed to query to the database, err: %v", err)
		return nil, err
	}
	var peopleCategories model.CategoriesJoined
	for rows.Next() {
		personCategories := model.CategoriesProductJoined{}
		err := rows.Scan(
			&personCategories.Categories.CategoryID,
			&personCategories.Categories.Type,
			&personCategories.Categories.SoldProductAmount,
			&personCategories.Categories.CreatedAt,
			&personCategories.Categories.UpdatedAt,
			&personCategories.Product.ProductID,
			&personCategories.Product.Title,
			&personCategories.Product.Price,
			&personCategories.Product.Stock,
			&personCategories.Product.CreatedAt,
			&personCategories.Product.UpdatedAt,
		)
		if err != nil {
			log.Printf("[ViewCategory] failed to scan the data from the database, err: %v", err)
			return nil, err
		}
		peopleCategories = append(peopleCategories, &personCategories)
	}
	return peopleCategories, nil
}

func (ctg CategoriesImplRepo) CountCategories(ctx context.Context) (int, error) {
	query := COUNT_CATEGORIES
	rows := ctg.db.QueryRowContext(ctx, query)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		log.Printf("[CountCategory] failed to scan the data from the database, err: %v", err)
		return 0, err
	}
	return count, nil
}

func (ctg CategoriesImplRepo) UpdateCategories(ctx context.Context, reqData model.Category, categoryID uint64) error {
	query := UPDATE_CATEGORIES

	stmt, err := ctg.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[UpdateCategory] failed to prepare the statement, err: %v", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, reqData.Type, categoryID)
	if err != nil {
		log.Printf("[UpdateCategory] failed to store data to the database, err: %v", err)
		return err
	}
	return nil
}

func (ctg CategoriesImplRepo) GetCategoriesByID(ctx context.Context, categoryID uint64) (*dto.EditCategoriesResponse, error) {
	query := GET_CATEGORIES_BY_ID
	stmt, err := ctg.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[GetCategoriesByID] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows := stmt.QueryRowContext(ctx, categoryID)
	if err != nil {
		log.Printf("[GetCategoriesByID] failed to query to the database, err: %v", err)
		return nil, err
	}
	personCategories := model.CategoriesProductJoined{}
	err = rows.Scan(
		&personCategories.Categories.CategoryID,
		&personCategories.Categories.Type,
		&personCategories.Categories.UpdatedAt,
	)
	if err != nil {
		log.Printf("[GetCategoriesByID] failed to scan the data from the database, err: %v", err)
		return nil, err
	}
	return dto.NewEditCategoriesResponse(personCategories.Categories), err
}

func (ctg CategoriesImplRepo) CheckCategories(ctx context.Context, categoryID uint64) (bool, error) {
	query := CHECK_CATEGORY
	stmt, err := ctg.db.PrepareContext(ctx, query)
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

func (ctg CategoriesImplRepo) DeleteCategories(ctx context.Context, categoryID uint64) error {
	query := DELETE_CATEGORIES

	stmt, err := ctg.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[DeleteCategory] failed to prepare the statement, err: %v", err)
		return err
	}

	_, err = stmt.QueryContext(ctx, categoryID)
	if err != nil {
		log.Printf("[DeleteCategory] failed to delete the category, err: %v", err)
		return err
	}
	return nil
}
