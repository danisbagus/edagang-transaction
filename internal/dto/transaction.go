package dto

import (
	"time"

	"github.com/danisbagus/edagang-pkg/errs"
	"github.com/danisbagus/edagang-transaction/internal/core/domain"
)

type TransactionResponse struct {
	TransactionID   string    `json:"transaction_id"`
	UserID          int64     `json:"user_id"`
	ProductID       int64     `json:"product_id"`
	Quantity        int64     `json:"quantity"`
	TransactionDate time.Time `json:"transaction_date"`
}

type TransactionListResponse struct {
	Transaction []TransactionResponse `json:"data"`
}

type NewTransactionRequest struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type NewTransactionResponse struct {
	TransactionID   string    `json:"transaction_id"`
	TransactionDate time.Time `json:"transaction_date"`
}

func NewGetListTransactionResponse(data []domain.TransactionModel) *TransactionListResponse {
	dataList := make([]TransactionResponse, len(data))

	for k, v := range data {
		dataList[k] = TransactionResponse{
			TransactionID:   v.TransactionID,
			ProductID:       v.ProductID,
			UserID:          v.UserID,
			Quantity:        v.Quantity,
			TransactionDate: v.TransactionDate,
		}
	}
	return &TransactionListResponse{Transaction: dataList}
}

func NewGetDetailTransactionResponse(data *domain.TransactionModel) *TransactionResponse {
	result := &TransactionResponse{
		TransactionID:   data.TransactionID,
		ProductID:       data.ProductID,
		Quantity:        data.Quantity,
		TransactionDate: data.TransactionDate,
	}
	return result
}

func NewNewTransactionResponse(data *domain.TransactionModel) *NewTransactionResponse {
	result := &NewTransactionResponse{
		TransactionID:   data.TransactionID,
		TransactionDate: data.TransactionDate,
	}

	return result
}

func (r NewTransactionRequest) Validate() *errs.AppError {
	if r.Quantity <= 0 {
		return errs.NewValidationError("Product quantity must more than 0")
	}
	return nil
}
