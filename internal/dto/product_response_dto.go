package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type CreateProductResponse struct {
	ProductID  uint64    `json:"id"`
	Title      string    `json:"title"`
	Price      uint64    `json:"price"`
	Stock      uint64    `json:"stock"`
	CategoryID uint64    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewProductCreateResponse(tsk model.Product, productID uint64) *CreateProductResponse {
	return &CreateProductResponse{
		ProductID:  productID,
		Title:      tsk.Title,
		Price:      tsk.Price,
		Stock:      tsk.Stock,
		CategoryID: tsk.CategoryID,
		CreatedAt:  time.Now(),
	}
}

type ViewProductResponse struct {
	ProductID  uint64    `json:"id"`
	Title      string    `json:"title"`
	Price      uint64    `json:"price"`
	Stock      uint64    `json:"stock"`
	CategoryID uint64    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type ViewProductsResponse []*ViewProductResponse

func NewViewProductResponse(p model.Product) *ViewProductResponse {
	return &ViewProductResponse{
		ProductID:  p.ProductID,
		Title:      p.Title,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryID,
		CreatedAt:  p.CreatedAt,
	}
}

func NewViewProductsResponse(p model.Products) ViewProductsResponse {
	var viewProductsResponse ViewProductsResponse

	for idx := range p {
		peopleTask := NewViewProductResponse(*p[idx])
		viewProductsResponse = append(viewProductsResponse, peopleTask)
	}
	return viewProductsResponse
}
