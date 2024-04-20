package payment

import "context"

type Repository interface {
	GetByPaymentId(ctx context.Context, paymentId string) (Payment, error)
	Create(ctx context.Context, payload Payment) error
}
