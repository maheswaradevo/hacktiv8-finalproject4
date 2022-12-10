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
	CategoryID        uint64                        `json:"id"`
	Type              string                        `json:"type"`
	SoldProductAmount uint64                        `json:"sold_product_amount"`
	CreatedAt         time.Time                     `json:"created_at"`
	UpdatedAt         time.Time                     `json:"updated_at"`
	Product           ViewCategoriesProductResponse `json:"product"`
}

type ViewCategoriesProductResponse struct {
	ProductID uint64    `db:"id"`
	Title     string    `db:"title"`
	Price     uint64    `db:"price"`
	Stock     uint64    `db:"stock"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// type ViewAllCategoriesResponse []*ViewCategoriesResponse

func NewViewCategoryResponse(ctg model.CategoriesProductJoined) ViewCategoriesResponse {
	var response ViewCategoriesResponse

	response.CategoryID = ctg.Categories.CategoryID
	response.Type = ctg.Categories.Type
	response.SoldProductAmount = ctg.Categories.SoldProductAmount
	response.CreatedAt = ctg.Categories.CreatedAt
	response.UpdatedAt = ctg.Categories.UpdatedAt

	response.Product.ProductID = ctg.Product.ProductID
	response.Product.Title = ctg.Product.Title
	response.Product.Price = ctg.Product.Price
	response.Product.Stock = ctg.Product.Stock
	response.Product.CreatedAt = ctg.Categories.CreatedAt
	response.Product.UpdatedAt = ctg.Categories.UpdatedAt

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

// func NewViewCategoriesResponse(ctg model.CategoriesProductJoined) *ViewCategoriesResponse {
// 	return &ViewCategoriesResponse{
// 		CategoryID:        ctg.Categories.CategoryID,
// 		Type:              ctg.Categories.Type,
// 		SoldProductAmount: ctg.Categories.SoldProductAmount,
// 		CreatedAt:         ctg.Categories.CreatedAt,
// 		UpdatedAt:         ctg.Categories.UpdatedAt,
// 		Product: ViewCategoriesProductResponse{
// 			ProductID: ctg.Product.ProductID,
// 			Title:     ctg.Product.Title,
// 			Price:     ctg.Product.Price,
// 			Stock:     ctg.Product.Stock,
// 			CreatedAt: ctg.Product.CreatedAt,
// 			UpdatedAt: ctg.Product.UpdatedAt,
// 		},
// 	}
// }

// func NewViewAllCategoriesResponse(ctg model.CategoriesJoined) *ViewAllCategoriesResponse {
// 	var viewAllCategoriesResponse ViewAllCategoriesResponse

// 	for idx := range ctg {
// 		peopleCategories := NewViewCategoriesResponse(*ctg[idx])
// 		viewAllCategoriesResponse = append(viewAllCategoriesResponse, peopleCategories)
// 	}
// 	return &viewAllCategoriesResponse
// }

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
