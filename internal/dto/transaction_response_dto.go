package dto

import "github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"

type DoTransactionResponse struct {
	Message         string `json:"message"`
	TransactionBill struct {
		TotalPrice   uint64 `json:"total_price"`
		Quantity     uint64 `json:"quantity"`
		ProductTitle string `json:"product_title"`
	} `json:"transaction_bill"`
}

func NewDoTransactionResponse(msg string, tr model.TransactionProductJoined) DoTransactionResponse {

	var response DoTransactionResponse

	response.Message = msg
	response.TransactionBill.TotalPrice = tr.TransactionHistory.TotalPrice
	response.TransactionBill.Quantity = tr.TransactionHistory.Quantity
	response.TransactionBill.ProductTitle = tr.Product.Title

	return response
}
