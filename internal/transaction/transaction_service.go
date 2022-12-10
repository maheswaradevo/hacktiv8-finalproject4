package transaction

import (
	"context"

	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
)

type TransactionService interface {
	DoTransaction(ctx context.Context, data *dto.DoTransactionRequest, userID uint64) (dto.DoTransactionResponse, error)
	ViewMyTransaction(ctx context.Context, userID uint64) ([]dto.ViewMyTransactionResponse, error)
}
