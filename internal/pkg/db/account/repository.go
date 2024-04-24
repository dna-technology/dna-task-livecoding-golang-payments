package account

import "context"

type Repository interface {
	GetByAccountId(ctx context.Context, accountId string) (Account, error)
	GetByUserId(ctx context.Context, userId string) (Account, error)
	Create(ctx context.Context, payload Account) error
	Update(ctx context.Context, payload Account) error
}
