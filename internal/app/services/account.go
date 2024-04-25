package services

import (
	"context"
	"database/sql"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/account"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/google/uuid"
)

type AccountService struct {
	accountRepository account.Repository
}

func NewAccountService(db *sql.DB) *AccountService {
	ar := account.NewSQLiteRepository(db)

	return &AccountService{accountRepository: ar}
}

func (as *AccountService) CreateUserAccount(ctx context.Context, userId string) (dto.AccountDto, error) {
	accountDto := dto.AccountDto{
		UserId:    userId,
		AccountId: uuid.NewString(),
		Balance:   float32(0),
	}

	err := as.accountRepository.Create(ctx, accountDto.ToEntity())
	if err != nil {
		return accountDto, err
	}

	return accountDto, nil
}

func (as *AccountService) GetUserAccount(ctx context.Context, userId string) (dto.AccountDto, error) {
	var accountDto dto.AccountDto

	account, err := as.accountRepository.GetByUserId(ctx, userId)
	if err != nil {
		return accountDto, err
	}

	return accountDto.FromEntity(account), nil
}

func (as *AccountService) IncreaseBalance(ctx context.Context, accountId string, amount float32) (dto.AccountDto, error) {
	var accountDto dto.AccountDto

	a, err := as.accountRepository.GetByAccountId(ctx, accountId)
	if err != nil {
		return accountDto, err
	}

	payload := dto.AccountDto{
		AccountId: accountId,
		Balance:   a.Balance + amount,
	}

	err = as.accountRepository.Update(ctx, payload.ToEntity())
	if err != nil {
		return payload, err
	}

	return payload, nil
}

func (as *AccountService) DecreaseBalance(ctx context.Context, accountId string, amount float32) (dto.AccountDto, error) {
	var accountDto dto.AccountDto

	a, err := as.accountRepository.GetByAccountId(ctx, accountId)
	if err != nil {
		return accountDto, err
	}

	payload := dto.AccountDto{
		AccountId: accountId,
		Balance:   a.Balance - amount,
	}

	err = as.accountRepository.Update(ctx, payload.ToEntity())
	if err != nil {
		return payload, err
	}

	return payload, nil
}
