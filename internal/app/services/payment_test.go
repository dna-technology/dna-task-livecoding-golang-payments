package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/stretchr/testify/assert"
)

func TestAddPayment_Success(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	ctx := context.Background()

	us := NewUserService(db)
	ms := NewMerchantService(db)
	as := NewAccountService(db)
	ps := NewPaymentService(db)

	createdUser, _ := us.CreateUser(ctx, dto.UserDto{
		FullName: "test user",
		Email:    "test@test.test",
	})
	createdMerchant, _ := ms.CreateMerchant(ctx, "testing payment")
	account, _ := as.GetUserAccount(ctx, createdUser.UserId)
	_, _ = as.IncreaseBalance(ctx, account.AccountId, float32(100))

	addedPayment, err := ps.AddPayment(ctx, dto.PaymentDto{
		MerchantId: createdMerchant.MerchantId,
		Amount:     10,
		UserId:     createdUser.UserId,
	})

	assert.NoError(t, err)
	assert.Equal(t, createdUser.UserId, addedPayment.UserId)
	assert.Equal(t, createdMerchant.MerchantId, addedPayment.MerchantId)
	assert.Equal(t, float32(10), addedPayment.Amount)

	account, _ = as.GetUserAccount(ctx, createdUser.UserId)

	assert.Equal(t, float32(90), account.Balance)
}
