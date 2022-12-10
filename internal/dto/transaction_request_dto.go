package dto

import "github.com/maheswaradevo/hacktiv8-finalproject4/internal/model"

type DoTransactionRequest struct {
	ProductID uint64 `json:"product_id" validate:"required"`
	Quantity  uint64 `json:"quantity" validate:"required"`
}

func (dto *DoTransactionRequest) ToEntity() (tr *model.TransactionHistory) {
	tr = &model.TransactionHistory{
		ProductID: dto.ProductID,
		Quantity:  dto.Quantity,
	}
	return
}
