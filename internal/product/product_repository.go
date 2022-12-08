package product

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, data model.Product) (productID uint64, err error)
	CheckCategory(ctx context.Context, categoryID uint64) (bool, error)
	ViewProduct(ctx context.Context) (model.Products, error)
	CountProduct(ctx context.Context) (int, error)
	CheckProduct(ctx context.Context, productID uint64) (bool, error)
	UpdateProduct(ctx context.Context, reqData model.ProductCategoryJoined, productID uint64) error
	GetProductByID(ctx context.Context, productID uint64) (*dto.EditProductResponse, error)
	DeleteProduct(ctx context.Context, productID uint64) error 
}
