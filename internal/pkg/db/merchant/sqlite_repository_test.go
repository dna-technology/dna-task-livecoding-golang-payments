package merchant

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

	err := repo.Create(ctx, Merchant{
		Name:       "Test Merchant",
		MerchantId: uuid.New().String(),
	})

	assert.NoError(t, err)
}

func TestGetByMerchantId(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../../test-database.sqlite")
	ctx := context.Background()

	repo := NewSQLiteRepository(db)

	merchantId := uuid.New().String()

	err := repo.Create(ctx, Merchant{
		Name:       "Test Merchant",
		MerchantId: merchantId,
	})
	assert.NoError(t, err)

	result, err := repo.GetByMerchantId(ctx, merchantId)
	assert.NoError(t, err)
	assert.Equal(t, "Test Merchant", result.Name)
	assert.Equal(t, merchantId, result.MerchantId)
}
