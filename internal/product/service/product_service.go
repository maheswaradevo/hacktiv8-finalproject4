package service

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/product"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
	"go.uber.org/zap"
)

type ProductServiceImpl struct {
	repo   product.ProductRepository
	logger *zap.Logger
}

func ProvideProductService(repo product.ProductRepository, logger *zap.Logger) *ProductServiceImpl {
	return &ProductServiceImpl{
		repo: repo,
		logger: logger,
	}
}

func (p *ProductServiceImpl) CreateProduct(ctx context.Context, data *dto.CreateProductRequest, userID uint64) (res *dto.CreateProductResponse, err error) {
	productData := data.ToProductEntity()
	validate := validator.New()
	validateError := validate.Struct(data)
	if validateError != nil {
		validateError = errors.ErrInvalidRequestBody
		p.logger.Sugar().Errorf("[CreateProduct] there's data that not through the validate process")
		return nil, validateError
	}

	check, err := p.repo.CheckCategory(ctx, data.CategoryID)
	if err != nil {
		p.logger.Sugar().Errorf("[CreateProduct] failed to check category with id: %v, err: %v", data.CategoryID, zap.Error(err))
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		p.logger.Sugar().Errorf("[CreateProduct] there's no category data with id: %v", data.CategoryID)
		return nil, err
	}
	productID, errCreate := p.repo.CreateProduct(ctx, *productData)
	if errCreate != nil {
		p.logger.Sugar().Errorf("[CreateProduct] failed to store user data to database: %v", zap.Error(errCreate))
		return
	}
	return dto.NewProductCreateResponse(*productData, productID), nil
}

func (p *ProductServiceImpl) ViewProduct(ctx context.Context) (dto.ViewProductsResponse, error) {
	count, errCount := p.repo.CountProduct(ctx)

	if errCount != nil {
		p.logger.Sugar().Errorf("[ViewProduct] failed to count the task, err: %v", zap.Error(errCount))
		return nil, errCount
	}
	if count == 0 {
		errCount = errors.ErrDataNotFound
		p.logger.Sugar().Errorf("[ViewProduct] no data exists in the database: %v", zap.Error(errCount))
		return nil, errCount
	}
	res, errView := p.repo.ViewProduct(ctx)
	if errView != nil {
		p.logger.Sugar().Errorf("[ViewProduct] failed to view the task, err: %v", zap.Error(errView))
		return nil, errView
	}
	return dto.NewViewProductsResponse(res), nil
}

func (p *ProductServiceImpl) UpdateProduct(ctx context.Context, productID uint64, userID uint64, data *dto.EditProductRequest) (*dto.EditProductResponse, error) {
	editedProduct := data.ToProductEntity()

	check, errCheck := p.repo.CheckProduct(ctx, productID)
	if errCheck != nil {
		p.logger.Sugar().Errorf("[UpdateProduct] failed to check product with, err: %v", zap.Error(errCheck))
		return nil, errCheck
	}
	if !check {
		errCheck = errors.ErrDataNotFound
		p.logger.Sugar().Errorf("[UpdateProduct] Product not found, err: %v", zap.Error(errCheck))
		return nil, errCheck
	}

	checkCategory, errCheckCategory := p.repo.CheckCategory(ctx, data.CategoryID)
	if errCheckCategory != nil {
		p.logger.Sugar().Errorf("[CreateProduct] failed to check category with id: %v, err: %v", data.CategoryID, zap.Error(errCheckCategory))
		return nil, errCheckCategory
	}
	if !checkCategory {
		errCheckCategory = errors.ErrDataNotFound
		p.logger.Sugar().Errorf("[CreateProduct] there's no category data with id: %v, err: %v", data.CategoryID, zap.Error(errCheckCategory))
		return nil, errCheckCategory
	}

	errCheckCategory = p.repo.UpdateProduct(ctx, *editedProduct, productID)
	if errCheckCategory != nil {
		p.logger.Sugar().Errorf("[UpdateProduct] failed to update product, err: %v", zap.Error(errCheckCategory))
		return nil, errCheckCategory
	}
	task, errGet := p.repo.GetProductByID(ctx, productID)
	if errGet != nil {
		p.logger.Sugar().Errorf("[UpdateProduct] failed to get product, err: %v", zap.Error(errGet))
		return nil, errGet
	}
	return task, nil
}

func (p *ProductServiceImpl) DeleteProduct(ctx context.Context, productID uint64, userID uint64) (*dto.DeleteProductResponse, error) {
	check, errCheck := p.repo.CheckProduct(ctx, productID)
	if errCheck != nil {
		p.logger.Sugar().Errorf("[DeleteProduct] failed to check product with,, err: %v", zap.Error(errCheck))
		return nil, errCheck
	}
	if !check {
		errCheck = errors.ErrDataNotFound
		p.logger.Sugar().Errorf("[DeleteProduct] no product in database")
		return nil, errCheck
	}

	errCheck = p.repo.DeleteProduct(ctx, productID)
	if errCheck != nil {
		p.logger.Sugar().Errorf("[DeleteProduct] failed to delete product, id: %v", productID)
		return nil, errCheck
	}
	message := "Your product has been successfully deleted"
	return dto.NewDeleteProductResponse(message), nil
}
