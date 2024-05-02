package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/user"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/google/uuid"
)

type UserService struct {
	userRepository user.Repository
	accountService *AccountService
}

func NewUserService(db *sql.DB) *UserService {
	ur := user.NewSQLiteRepository(db)
	as := NewAccountService(db)

	return &UserService{
		userRepository: ur,
		accountService: as,
	}
}

func (us *UserService) CreateUser(ctx context.Context, payload dto.UserDto) (dto.UserDto, error) {
	userDto := dto.UserDto{
		UserId:   uuid.NewString(),
		FullName: payload.FullName,
		Email:    payload.Email,
	}

	log.Printf("About to create user (userId: %q, fullname: %q, email: %q)", userDto.UserId, userDto.FullName, userDto.Email)

	err := us.userRepository.Create(ctx, userDto.ToEntity())
	if err != nil {
		return userDto, err
	}

	_, err = us.accountService.CreateUserAccount(ctx, userDto.UserId)
	if err != nil {
		return userDto, err
	}

	return userDto, err
}

func (us *UserService) GetUser(ctx context.Context, userId string) (dto.UserDto, error) {
	var userDto dto.UserDto

	user, err := us.userRepository.GetByUserId(ctx, userId)
	if err != nil {
		return userDto, err
	}

	return userDto.FromEntity(user), nil
}
