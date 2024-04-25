package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/payment"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddUserSavesUser(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../test-database.sqlite")
	app := NewApplication(db)

	fullName := "John Doe"
	email := "john.doe@digitalnewagency.com"

	requestBody, _ := json.Marshal(map[string]any{
		"fullName": fullName,
		"email":    email,
	})
	server := httptest.NewServer(app.routes())
	defer server.Close()

	resp, err := server.Client().Post(server.URL+"/users", "application/json", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	responseBody, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response dto.UserDto
	json.Unmarshal(responseBody, &response)
	assert.NotEmpty(t, response.UserId)
	assert.Equal(t, response.FullName, fullName)
	assert.Equal(t, response.Email, email)

	userDto, err := app.userService.GetUser(context.Background(), response.UserId)
	assert.Nil(t, err)
	assert.Equal(t, userDto.FullName, fullName)
	assert.Equal(t, userDto.Email, email)

	accountDto, err := app.accountService.GetUserAccount(context.Background(), response.UserId)
	assert.Nil(t, err)
	assert.NotEmpty(t, accountDto.AccountId)
	assert.Zero(t, accountDto.Balance)
}

func TestAddPaymentSavesTransaction(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../test-database.sqlite")
	app := NewApplication(db)

	userId := thereIsAUser(app)
	merchantId := thereIsAMerchant(app)
	initialBalance := userHasPositiveAccountBalance(app, userId)
	amount := int64(10)

	requestBody, _ := json.Marshal(map[string]any{
		"userId":     userId,
		"merchantId": merchantId,
		"amount":     amount,
	})

	server := httptest.NewServer(app.routes())
	defer server.Close()

	resp, err := server.Client().Post(server.URL+"/transactions", "application/json", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	responseBody, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response dto.PaymentDto
	json.Unmarshal(responseBody, &response)

	assert.NotEmpty(t, response.PaymentId)
	assert.Equal(t, response.UserId, userId)
	assert.Equal(t, response.MerchantId, merchantId)
	assert.Equal(t, response.Amount, amount)

	createdPayment, err := payment.NewSQLiteRepository(db).GetByPaymentId(context.Background(), response.PaymentId)
	assert.Nil(t, err)
	assert.Equal(t, createdPayment.UserId, userId)
	assert.Equal(t, createdPayment.MerchantId, merchantId)
	assert.Equal(t, createdPayment.Amount, amount)

	account, _ := app.accountService.GetUserAccount(context.Background(), userId)
	assert.Equal(t, initialBalance-amount, account.Balance)
}

func thereIsAUser(app *Application) string {
	user, _ := app.userService.CreateUser(context.Background(), dto.UserDto{
		FullName: "Jan Kowalski",
		Email:    "jan.kowalski@digitalnewagency.com",
	})
	return user.UserId
}

func thereIsAMerchant(app *Application) string {
	merchant, _ := app.merchantService.CreateMerchant(context.Background(), "DNA")
	return merchant.MerchantId
}

func userHasPositiveAccountBalance(app *Application, userId string) int64 {
	account, _ := app.accountService.GetUserAccount(context.Background(), userId)

	updatedAccount, _ := app.accountService.IncreaseBalance(context.Background(), account.AccountId, 100)
	return updatedAccount.Balance
}
