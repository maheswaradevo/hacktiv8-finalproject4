package dto

import (
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type CreateProductRequest struct {
	Title      string `json:"title" validate:"required"`
	Price      uint64 `json:"price" validate:"required"`
	Stock      uint64 `json:"stock" validate:"required"`
	CategoryID uint64 `json:"category_id"`
}

func (dto *CreateProductRequest) ToProductEntity() (cmt *model.Product) {
	cmt = &model.Product{
		Title:      dto.Title,
		Price:      dto.Price,
		Stock:      dto.Stock,
		CategoryID: dto.CategoryID,
	}
	return
}

type EditProductRequest struct {
	Title      string `json:"title"`
	Price      uint64 `json:"price"`
	Stock      uint64 `json:"stock"`
	CategoryID uint64 `json:"category_id"`
}

func (dto *EditProductRequest) ToProductEntity() *model.ProductCategoryJoined {
	return &model.ProductCategoryJoined{
		Product: model.Product{
			Title:       dto.Title,
			Price:       dto.Price,
			Stock:       dto.Stock,
			CategoryID:  dto.CategoryID,
		},
	}
}
