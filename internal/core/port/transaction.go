package port

import (
	"github.com/danisbagus/edagang-pkg/errs"
	"github.com/danisbagus/edagang-transaction/internal/core/domain"
	"github.com/danisbagus/edagang-transaction/internal/dto"
)

type ITransactionRepo interface {
	FindAll() ([]domain.TransactionModel, *errs.AppError)
	// FindOneByID(transactionID string) (*domain.TransactionModel, *errs.AppError)
}

type ITransactionService interface {
	GetAll() (*dto.TransactionListResponse, *errs.AppError)
	// GetDetail(productID int64) (*dto.TransactionResponse, *errs.AppError)
}
