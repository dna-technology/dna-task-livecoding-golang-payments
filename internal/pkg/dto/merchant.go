package dto

import (
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/merchant"
)

type MerchantDto struct {
	Name       string
	MerchantId string
}

func (m *MerchantDto) FromEntity(entity merchant.Merchant) MerchantDto {
	return MerchantDto{
		Name:       entity.Name,
		MerchantId: entity.MerchantId,
	}
}

func (m *MerchantDto) ToEntity() merchant.Merchant {
	return merchant.Merchant{
		Name:       m.Name,
		MerchantId: m.MerchantId,
	}
}
