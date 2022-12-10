package categories

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type CategoriesRepository interface {
	CreateCategories(ctx context.Context, data model.Categories) (categoryID uint64, err error)
	CheckCategories(ctx context.Context, categoryID uint64) (bool, error)
	ViewCategories(ctx context.Context) (model.CategoriesJoined, error)
	CountCategories(ctx context.Context) (int, error)
	UpdateCategories(ctx context.Context, reqData model.Categories, categoryID uint64) error
	DeleteCategories(ctx context.Context, categoryID uint64) error
	GetCategoriesByID(ctx context.Context, categoryID uint64) (*dto.EditCategoriesResponse, error)
}
