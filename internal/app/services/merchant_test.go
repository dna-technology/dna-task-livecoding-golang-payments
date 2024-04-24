package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMerchant(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	ctx := context.Background()

	ms := NewMerchantService(db)

	createdMerchant, err := ms.CreateMerchant(ctx, "testing")

	assert.NoError(t, err)
	assert.Equal(t, "testing", createdMerchant.Name)
	assert.NotEmpty(t, createdMerchant.MerchantId)
}

func TestGetMerchant(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	ctx := context.Background()

	ms := NewMerchantService(db)

	createdMerchant, err := ms.CreateMerchant(ctx, "testing get")
	assert.NoError(t, err)

	merchant, err := ms.GetMerchant(ctx, createdMerchant.MerchantId)

	assert.NoError(t, err)
	assert.Equal(t, "testing get", merchant.Name)
	assert.Equal(t, createdMerchant.MerchantId, merchant.MerchantId)
}
