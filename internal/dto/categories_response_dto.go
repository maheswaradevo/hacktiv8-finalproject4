package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type CreateCategoriesResponse struct {
	CategoryID        uint64    `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint64    `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

func NewCategoriesCreateResponse(ctg model.Category, userID uint64, categoryID uint64) *CreateCategoriesResponse {
	return &CreateCategoriesResponse{
		CategoryID:        categoryID,
		Type:              ctg.Type,
		SoldProductAmount: ctg.SoldProductAmount,
		CreatedAt:         time.Now(),
	}
}

type ViewCategoriesResponse struct {
	CategoryID        uint64                          `json:"id"`
	Type              string                          `json:"type"`
	SoldProductAmount uint64                          `json:"sold_product_amount"`
	CreatedAt         time.Time                       `json:"created_at"`
	UpdatedAt         time.Time                       `json:"updated_at"`
	Product           []ViewCategoriesProductResponse `json:"products"`
}

type ViewCategoriesProductResponse struct {
	ProductID uint64    `json:"id"`
	Title     string    `json:"title"`
	Price     uint64    `json:"price"`
	Stock     uint64    `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProductResponse(tsk model.Product) *ViewCategoriesProductResponse {
	return &ViewCategoriesProductResponse{
		ProductID: tsk.ProductID,
		Title:     tsk.Title,
		Price:     tsk.Price,
		Stock:     tsk.Stock,
		CreatedAt: time.Now(),
	}
}

func NewViewCategoryResponse(ctg model.CategoriesProductJoined) ViewCategoriesResponse {
	var response ViewCategoriesResponse

	response.CategoryID = ctg.Categories.CategoryID
	response.Type = ctg.Categories.Type
	response.SoldProductAmount = ctg.Categories.SoldProductAmount
	response.CreatedAt = ctg.Categories.CreatedAt
	response.UpdatedAt = ctg.Categories.UpdatedAt

	for _, p := range ctg.Product {
		product := NewProductResponse(p)
		response.Product = append(response.Product, *product)
	}
	return response
}

func NewViewCategoriesResponse(ctg model.CategoriesJoined) []ViewCategoriesResponse {
	var responses []ViewCategoriesResponse

	for _, ctgs := range ctg {
		categoriesHistory := NewViewCategoryResponse(*ctgs)
		responses = append(responses, categoriesHistory)
	}
	return responses
}

type EditCategoriesResponse struct {
	CategoryID        uint64    `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint64    `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func NewEditCategoriesResponse(ctg model.Category) *EditCategoriesResponse {
	return &EditCategoriesResponse{
		CategoryID:        ctg.CategoryID,
		Type:              ctg.Type,
		SoldProductAmount: ctg.SoldProductAmount,
		UpdatedAt:         ctg.UpdatedAt,
	}
}

type DeleteCategoriesResponse struct {
	Message string `json:"message"`
}

func NewDeleteCategoriesResponse(message string) *DeleteCategoriesResponse {
	return &DeleteCategoriesResponse{
		Message: message,
	}
}
