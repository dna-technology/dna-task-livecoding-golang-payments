package user

import "context"

type Repository interface {
	GetByUserId(ctx context.Context, userId string) (User, error)
	Create(ctx context.Context, payload User) error
}
