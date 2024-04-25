package routers

import (
	"database/sql"
	"encoding/json"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/app/services"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"net/http"
)

type TransactionRouter struct {
	paymentService *services.PaymentService
}

func NewTransactionRouter(db *sql.DB) *TransactionRouter {
	return &TransactionRouter{
		paymentService: services.NewPaymentService(db),
	}
}

func (t *TransactionRouter) AddPayment(w http.ResponseWriter, r *http.Request) {
	var paymentDto dto.PaymentDto
	err := json.NewDecoder(r.Body).Decode(&paymentDto)
	if err != nil {
		responseJson(w, http.StatusBadRequest, err.Error())
		return
	}

	paymentDto, err = t.paymentService.AddPayment(r.Context(), paymentDto)
	if err != nil {
		responseJson(w, http.StatusBadRequest, err.Error())
		return
	}

	responseJson(w, http.StatusCreated, paymentDto)
}
