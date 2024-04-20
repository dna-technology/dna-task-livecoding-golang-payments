package payment

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

	err := repo.Create(ctx, Payment{
		UserId:     uuid.New().String(),
		PaymentId:  uuid.New().String(),
		MerchantId: uuid.New().String(),
		Amount:     0,
	})

	assert.NoError(t, err)
}

func TestGetByPaymentId(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../../test-database.sqlite")
	ctx := context.Background()

	repo := NewSQLiteRepository(db)

	userId := uuid.New().String()
	paymentId := uuid.New().String()
	merchantId := uuid.New().String()

	err := repo.Create(ctx, Payment{
		UserId:     userId,
		PaymentId:  paymentId,
		MerchantId: merchantId,
		Amount:     0,
	})
	assert.NoError(t, err)

	result, err := repo.GetByPaymentId(ctx, paymentId)
	assert.NoError(t, err)
	assert.Equal(t, userId, result.UserId)
	assert.Equal(t, paymentId, result.PaymentId)
	assert.Equal(t, merchantId, result.MerchantId)
	assert.Equal(t, int64(0), result.Amount)
}
