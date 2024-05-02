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
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/stretchr/testify/assert"
)

func TestUserRouter_CreateUser(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")

	fullName := "John Doe"
	email := "john.doe@digitalnewagency.com"
	requestBody, _ := json.Marshal(map[string]any{
		"fullName": fullName,
		"email":    email,
	})

	server := httptest.NewServer(Routes(db))
	defer server.Close()

	resp, err := server.Client().Post(server.URL+"/users", "application/json", bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	response := getUserDtoFrom(resp)
	assertUserDtoResponse(t, response, fullName, email)
	assertUserCreated(t, db, response, fullName, email)
	assertAccountCreated(t, db, response)
}

func assertUserDtoResponse(t *testing.T, response dto.UserDto, fullName string, email string) {
	assert.NotEmpty(t, response.UserId)
	assert.Equal(t, response.FullName, fullName)
	assert.Equal(t, response.Email, email)
}

func assertAccountCreated(t *testing.T, db *sql.DB, response dto.UserDto) {
	accountDto, err := services.NewAccountService(db).GetUserAccount(context.Background(), response.UserId)
	assert.Nil(t, err)
	assert.NotEmpty(t, accountDto.AccountId)
	assert.Zero(t, accountDto.Balance)
}

func assertUserCreated(t *testing.T, db *sql.DB, response dto.UserDto, fullName string, email string) {
	userDto, err := services.NewUserService(db).GetUser(context.Background(), response.UserId)
	assert.Nil(t, err)
	assert.Equal(t, userDto.FullName, fullName)
	assert.Equal(t, userDto.Email, email)
}

func getUserDtoFrom(resp *http.Response) dto.UserDto {
	responseBody, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var response dto.UserDto
	json.Unmarshal(responseBody, &response)
	return response
}
