package routers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/app/services"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/payment"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRouter_AddPayment(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	userId := thereIsAUser(db)
	merchantId := thereIsAMerchant(db)
	initialBalance := userHasInitialBalance(db, userId, 100)

	amount := float32(10)
	requestBody, _ := json.Marshal(map[string]any{
		"userId":     userId,
		"merchantId": merchantId,
		"amount":     amount,
	})

	// fixme
	server := httptest.NewServer(Routes(db))
	defer server.Close()

	resp, err := server.Client().Post(server.URL+"/transactions", "application/json", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	response := getPaymentDtoFrom(resp)
	assertPaymentDtoResponse(t, response, userId, merchantId, amount)
	assertPaymentStored(t, db, response, userId, merchantId, amount)
	assertAccountBalanceDecreased(t, db, userId, initialBalance, amount)
}

func thereIsAUser(db *sql.DB) string {
	user, _ := services.NewUserService(db).CreateUser(context.Background(), dto.UserDto{
		FullName: "Jan Kowalski",
		Email:    "jan.kowalski@digitalnewagency.com",
	})
	userId := user.UserId
	return userId
}

func thereIsAMerchant(db *sql.DB) string {
	merchant, _ := services.NewMerchantService(db).CreateMerchant(context.Background(), "DNA")
	return merchant.MerchantId
}

func userHasInitialBalance(db *sql.DB, userId string, amount float32) float32 {
	accountService := services.NewAccountService(db)

	account, _ := accountService.GetUserAccount(context.Background(), userId)

	updatedAccount, _ := accountService.IncreaseBalance(context.Background(), account.AccountId, amount)
	return updatedAccount.Balance
}

func getPaymentDtoFrom(resp *http.Response) dto.PaymentDto {
	responseBody, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response dto.PaymentDto
	json.Unmarshal(responseBody, &response)
	return response
}

func assertPaymentDtoResponse(t *testing.T, response dto.PaymentDto, userId string, merchantId string, amount float32) {
	assert.NotEmpty(t, response.PaymentId)
	assert.Equal(t, response.UserId, userId)
	assert.Equal(t, response.MerchantId, merchantId)
	assert.Equal(t, response.Amount, amount)
}

func assertPaymentStored(t *testing.T, db *sql.DB, response dto.PaymentDto, userId string, merchantId string, amount float32) {
	createdPayment, err := payment.NewSQLiteRepository(db).GetByPaymentId(context.Background(), response.PaymentId)
	assert.Nil(t, err)
	assert.Equal(t, createdPayment.UserId, userId)
	assert.Equal(t, createdPayment.MerchantId, merchantId)
	assert.Equal(t, createdPayment.Amount, amount)
}

func assertAccountBalanceDecreased(t *testing.T, db *sql.DB, userId string, initialBalance float32, amount float32) {
	account, err := services.NewAccountService(db).GetUserAccount(context.Background(), userId)
	assert.Nil(t, err)
	assert.Equal(t, initialBalance-amount, account.Balance)
}
