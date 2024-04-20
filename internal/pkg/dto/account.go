package dto

import "github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/account"

type AccountDto struct {
	AccountId string
	UserId    string
	Balance   int64
}

func (a *AccountDto) FromEntity(entity account.Account) AccountDto {
	return AccountDto{
		UserId:    entity.UserId,
		AccountId: entity.AccountId,
		Balance:   entity.Balance,
	}
}
