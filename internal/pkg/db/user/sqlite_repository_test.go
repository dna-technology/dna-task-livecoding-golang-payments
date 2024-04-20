package user

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

	err := repo.Create(ctx, User{
		UserId:   uuid.New().String(),
		FullName: "Test Testing",
		Email:    "test@testing.test",
	})

	assert.NoError(t, err)
}

func TestGetByUserId(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../../test-database.sqlite")
	ctx := context.Background()

	repo := NewSQLiteRepository(db)

	userId := uuid.New().String()

	err := repo.Create(ctx, User{
		UserId:   userId,
		FullName: "Get Testing",
		Email:    "get@testing.test",
	})
	assert.NoError(t, err)

	result, err := repo.GetByUserId(ctx, userId)
	assert.NoError(t, err)
	assert.Equal(t, userId, result.UserId)
	assert.Equal(t, "Get Testing", result.FullName)
	assert.Equal(t, "get@testing.test", result.Email)
}
