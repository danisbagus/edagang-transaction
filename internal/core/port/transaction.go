package port

import (
	"github.com/danisbagus/edagang-pkg/errs"
	"github.com/danisbagus/edagang-transaction/internal/core/domain"
	"github.com/danisbagus/edagang-transaction/internal/dto"
)

type ITransactionRepo interface {
	FindAll() ([]domain.TransactionModel, *errs.AppError)
	FindOneByID(transactionID string) (*domain.TransactionModel, *errs.AppError)
	Create(data *domain.TransactionModel) *errs.AppError
	Delete(transactionID string) *errs.AppError
}

type ITransactionService interface {
	GetAll() (*dto.TransactionListResponse, *errs.AppError)
	GetDetail(transactionID string) (*dto.TransactionResponse, *errs.AppError)
	NewTransaction(data *dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
	RemoveTransaction(transactionID string) *errs.AppError
}
