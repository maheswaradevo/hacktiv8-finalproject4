package product

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
)

type ProductService interface {
	CreateProduct(ctx context.Context, data *dto.CreateProductRequest, userID uint64) (res *dto.CreateProductResponse, err error)
	ViewProduct(ctx context.Context) (dto.ViewProductsResponse, error) 
}
