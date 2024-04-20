package account

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../../test-database.sqlite")
	ctx := context.Background()

	repo := NewSQLiteRepository(db)

	err := repo.Create(ctx, Account{
		AccountId: uuid.New().String(),
		UserId:    uuid.New().String(),
		Balance:   0,
	})

	assert.NoError(t, err)
}

func TestGetByAccountId(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../../test-database.sqlite")
	ctx := context.Background()

	repo := NewSQLiteRepository(db)

	userId := uuid.New().String()
	accountId := uuid.New().String()

	err := repo.Create(ctx, Account{
		AccountId: accountId,
		UserId:    userId,
		Balance:   0,
	})
	assert.NoError(t, err)

	result, err := repo.GetByAccountId(ctx, accountId)
	assert.NoError(t, err)
	assert.Equal(t, userId, result.UserId)
	assert.Equal(t, accountId, result.AccountId)
	assert.Equal(t, int64(0), result.Balance)
}

func TestGetByUserId(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../../test-database.sqlite")
	ctx := context.Background()

	repo := NewSQLiteRepository(db)

	userId := uuid.New().String()
	accountId := uuid.New().String()

	err := repo.Create(ctx, Account{
		AccountId: accountId,
		UserId:    userId,
		Balance:   0,
	})
	assert.NoError(t, err)

	result, err := repo.GetByUserId(ctx, userId)
	assert.NoError(t, err)
	assert.Equal(t, userId, result.UserId)
	assert.Equal(t, accountId, result.AccountId)
	assert.Equal(t, int64(0), result.Balance)
}
