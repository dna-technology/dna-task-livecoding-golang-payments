package services

import (
	"context"
	"database/sql"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/account"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
)

type AccountService struct {
	accountRepository account.Repository
}

func NewAccountService(db *sql.DB) *AccountService {
	ar := account.NewSQLiteRepository(db)

	return &AccountService{accountRepository: ar}
}

func (as *AccountService) CreateUserAccount(ctx context.Context, userId string) {}

func (as *AccountService) GetUserAccount(ctx context.Context, userId string) (dto.AccountDto, error) {
	var accountDto dto.AccountDto

	account, err := as.accountRepository.GetByUserId(ctx, userId)
	if err != nil {
		return accountDto, err
	}

	return accountDto.FromEntity(account), nil
}
