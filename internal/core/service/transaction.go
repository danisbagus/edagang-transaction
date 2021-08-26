package service

import (
	"github.com/danisbagus/edagang-pkg/errs"
	"github.com/danisbagus/edagang-transaction/internal/core/port"
	"github.com/danisbagus/edagang-transaction/internal/dto"
)

type TransactionService struct {
	repo port.ITransactionRepo
}

func NewTransactionService(repo port.ITransactionRepo) port.ITransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (r TransactionService) GetAll() (*dto.TransactionListResponse, *errs.AppError) {

	dataList, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListTransactionResponse(dataList)

	return response, nil
}

func (r TransactionService) GetDetail(transactionID string) (*dto.TransactionResponse, *errs.AppError) {
	data, err := r.repo.FindOneByID(transactionID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailTransactionResponse(data)

	return response, nil
}
