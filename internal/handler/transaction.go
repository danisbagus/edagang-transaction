package handler

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/edagang-transaction/internal/core/port"
)

type TransactionHandler struct {
	Service port.ITransactionService
}

func (rc TransactionHandler) GetTransactionList(w http.ResponseWriter, r *http.Request) {
	dataList, err := rc.Service.GetAll()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, dataList)
}

// func (rc TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
// 	var request dto.NewTransactionRequest

// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		writeResponse(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	data, err := rc.Service.NewTransaction(&request)
// 	if err != nil {
// 		writeResponse(w, err.Code, err.AsMessage())
// 		return
// 	}
// 	writeResponse(w, http.StatusCreated, data)
// }

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
