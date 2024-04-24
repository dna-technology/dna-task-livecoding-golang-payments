package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAccount(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	ctx := context.Background()

	as := NewAccountService(db)
	userId := uuid.NewString()

	createdAccount, err := as.CreateUserAccount(ctx, userId)

	assert.NoError(t, err)
	assert.Equal(t, userId, createdAccount.UserId)
	assert.Equal(t, int64(0), createdAccount.Balance)
	assert.NotEmpty(t, createdAccount.AccountId)
}

func TestGetUserAccount(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	ctx := context.Background()

	as := NewAccountService(db)
	userId := uuid.NewString()

	createdAccount, err := as.CreateUserAccount(ctx, userId)
	assert.NoError(t, err)

	getAccount, err := as.GetUserAccount(ctx, userId)

	assert.NoError(t, err)
	assert.Equal(t, createdAccount, getAccount)
}

func TestIncreaseBalance(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	ctx := context.Background()

	as := NewAccountService(db)
	userId := uuid.NewString()

	createdAccount, err := as.CreateUserAccount(ctx, userId)
	assert.NoError(t, err)

	updatedAccount, err := as.IncreaseBalance(ctx, createdAccount.AccountId, 10)

	assert.NoError(t, err)
	assert.Equal(t, int64(10), updatedAccount.Balance)
}

func TestDecreaseBalance(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	ctx := context.Background()

	as := NewAccountService(db)
	userId := uuid.NewString()

	createdAccount, err := as.CreateUserAccount(ctx, userId)
	assert.NoError(t, err)

	updatedAccount, err := as.IncreaseBalance(ctx, createdAccount.AccountId, 10)

	assert.NoError(t, err)
	assert.Equal(t, int64(10), updatedAccount.Balance)

	updatedAccount, err = as.DecreaseBalance(ctx, createdAccount.AccountId, 5)

	assert.NoError(t, err)
	assert.Equal(t, int64(5), updatedAccount.Balance)
}
