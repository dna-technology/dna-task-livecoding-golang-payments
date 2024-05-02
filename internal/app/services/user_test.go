package services

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	db, _ := sql.Open("sqlite3", "../../../test-database.sqlite")
	ctx := context.Background()

	us := NewUserService(db)

	createdUser, err := us.CreateUser(ctx, dto.UserDto{FullName: "get test user", Email: "get@test.test"})
	assert.NoError(t, err)

	user, err := us.GetUser(ctx, createdUser.UserId)

	assert.NoError(t, err)
	assert.Equal(t, "get test user", user.FullName)
	assert.Equal(t, "get@test.test", user.Email)
	assert.Equal(t, createdUser.UserId, user.UserId)
}
