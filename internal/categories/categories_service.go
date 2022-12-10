package categories

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
)

type CategoriesService interface {
	CreateCategories(ctx context.Context, data *dto.CreateCategoriesRequest, userID uint64) (res *dto.CreateCategoriesResponse, err error)
	ViewCategories(ctx context.Context) ([]dto.ViewCategoriesResponse, error)
	UpdateCategories(ctx context.Context, categoryID uint64, userID uint64, data *dto.EditCategoriesRequest) (*dto.EditCategoriesResponse, error)
	DeleteCategories(ctx context.Context, categoryID uint64, userID uint64) (*dto.DeleteCategoriesResponse, error)
}
