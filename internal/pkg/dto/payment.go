package dto

import (
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/payment"
)

type PaymentDto struct {
	PaymentId  string
	UserId     string
	MerchantId string
	Amount     float32
}

func (p *PaymentDto) FromEntity(entity payment.Payment) PaymentDto {
	return PaymentDto{
		PaymentId:  entity.PaymentId,
		UserId:     entity.UserId,
		MerchantId: entity.MerchantId,
		Amount:     entity.Amount,
	}
}

func (p *PaymentDto) ToEntity() payment.Payment {
	return payment.Payment{
		PaymentId:  p.PaymentId,
		UserId:     p.UserId,
		MerchantId: p.MerchantId,
		Amount:     p.Amount,
	}
}
