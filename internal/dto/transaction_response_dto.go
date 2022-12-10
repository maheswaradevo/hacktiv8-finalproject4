package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"
)

type (
	DoTransactionResponse struct {
		Message         string `json:"message"`
		TransactionBill struct {
			TotalPrice   uint64 `json:"total_price"`
			Quantity     uint64 `json:"quantity"`
			ProductTitle string `json:"product_title"`
		} `json:"transaction_bill"`
	}

	ViewMyTransactionResponse struct {
		TransactionID uint64          `json:"id"`
		ProductID     uint64          `json:"product_id"`
		UserID        uint64          `json:"user_id"`
		Quantity      uint64          `json:"quantity"`
		TotalPrice    uint64          `json:"total_price"`
		Product       ProductResponse `json:"product"`
	}

	ProductResponse struct {
		ProductID  uint64    `json:"id"`
		Title      string    `json:"title"`
		Price      uint64    `json:"price"`
		Stock      uint64    `json:"stock"`
		CategoryID uint64    `json:"category_Id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)

func NewDoTransactionResponse(msg string, tr model.TransactionProductJoined) DoTransactionResponse {

	var response DoTransactionResponse

	response.Message = msg
	response.TransactionBill.TotalPrice = tr.TransactionHistory.TotalPrice
	response.TransactionBill.Quantity = tr.TransactionHistory.Quantity
	response.TransactionBill.ProductTitle = tr.Product.Title

	return response
}

func NewViewMyTransactionResponse(tr model.TransactionProductJoined) ViewMyTransactionResponse {
	var response ViewMyTransactionResponse

	response.TransactionID = tr.TransactionHistory.TransactionID
	response.ProductID = tr.Product.ProductID
	response.UserID = tr.TransactionHistory.UserID
	response.Quantity = tr.TransactionHistory.Quantity
	response.TotalPrice = tr.TransactionHistory.TotalPrice

	response.Product.ProductID = tr.Product.ProductID
	response.Product.Title = tr.Product.Title
	response.Product.Price = tr.Product.Price
	response.Product.Stock = tr.Product.Stock
	response.Product.CategoryID = tr.Product.CategoryID
	response.Product.CreatedAt = tr.Product.CreatedAt
	response.Product.UpdatedAt = tr.Product.UpdatedAt

	return response
}

func NewViewMyTransactionsResponse(tr model.TransactionHistories) []ViewMyTransactionResponse {
	var responses []ViewMyTransactionResponse

	for _, trs := range tr {
		transactionHistory := NewViewMyTransactionResponse(*trs)
		responses = append(responses, transactionHistory)
	}
	return responses
}
