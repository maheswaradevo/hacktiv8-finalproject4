package service

import (
	"context"
	"log"

	"github.com/go-playground/validator"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/product"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
)

type ProductServiceImpl struct {
	repo product.ProductRepository
}

func ProvideProductService(repo product.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		repo: repo,
	}
}

func (p *ProductServiceImpl) CreateProduct(ctx context.Context, data *dto.CreateProductRequest, userID uint64) (res *dto.CreateProductResponse, err error) {
	productData := data.ToProductEntity()
	validate := validator.New()
	validateError := validate.Struct(data)
	if validateError != nil {
		validateError = errors.ErrInvalidRequestBody
		log.Printf("[CreateProduct] there's data that not through the validate process")
		return nil, validateError
	}
	productID, err := p.repo.CreateProduct(ctx, *productData)
	if err != nil {
		log.Printf("[CreateProduct] failed to store user data to database: %v", err)
		return
	}
	return dto.NewProductCreateResponse(*productData, productID), nil
}

func (p *ProductServiceImpl) ViewProduct(ctx context.Context) (dto.ViewProductsResponse, error) {
	count, err := p.repo.CountProduct(ctx)

	if err != nil {
		log.Printf("[ViewProduct] failed to count the task, err: %v", err)
		return nil, err
	}
	if count == 0 {
		err = errors.ErrDataNotFound
		log.Printf("[ViewProduct] no data exists in the database: %v", err)
		return nil, err
	}
	res, err := p.repo.ViewProduct(ctx)
	if err != nil {
		log.Printf("[ViewProduct] failed to view the task, err: %v", err)
		return nil, err
	}
	return dto.NewViewProductsResponse(res), nil
}
