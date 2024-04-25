package dto

import "github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/account"

type AccountDto struct {
	AccountId string
	UserId    string
	Balance   float32
}

func (a *AccountDto) FromEntity(entity account.Account) AccountDto {
	return AccountDto{
		UserId:    entity.UserId,
		AccountId: entity.AccountId,
		Balance:   entity.Balance,
	}
}

func (a *AccountDto) ToEntity() account.Account {
	return account.Account{
		UserId:    a.UserId,
		AccountId: a.AccountId,
		Balance:   a.Balance,
	}
}
