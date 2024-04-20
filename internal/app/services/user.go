package services

import (
	"context"
	"database/sql"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/user"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
)

type UserService struct {
	userRepository user.Repository
}

func NewUserService(db *sql.DB) *UserService {
	ur := user.NewSQLiteRepository(db)

	return &UserService{userRepository: ur}
}

func (us *UserService) CreateUser() {}

func (us *UserService) GetUser(ctx context.Context, userId string) (dto.UserDto, error) {
	var userDto dto.UserDto

	user, err := us.userRepository.GetByUserId(ctx, userId)
	if err != nil {
		return userDto, err
	}

	return userDto.FromEntity(user), nil
}
