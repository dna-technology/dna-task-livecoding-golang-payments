package merchant

import "context"

type Repository interface {
	GetByMerchantId(ctx context.Context, merchantId string) (Merchant, error)
	Create(ctx context.Context, payload Merchant) error
}
