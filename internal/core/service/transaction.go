package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/danisbagus/edagang-pkg/errs"
	"github.com/danisbagus/edagang-transaction/internal/core/domain"
	"github.com/danisbagus/edagang-transaction/internal/core/port"
	"github.com/danisbagus/edagang-transaction/internal/dto"
)

type TransactionService struct {
	repo port.ITransactionRepo
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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

func (r TransactionService) NewTransaction(data *dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	transactionID := fmt.Sprintf("TR%v", String(6))
	form := domain.TransactionModel{
		TransactionID:   transactionID,
		ProductID:       data.ProductID,
		Quantity:        data.Quantity,
		TransactionDate: time.Now(),
	}

	err := r.repo.Create(&form)
	if err != nil {
		return nil, err
	}

	response := dto.NewNewTransactionResponse(&form)
	return response, nil
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
