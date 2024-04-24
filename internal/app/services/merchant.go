package services

import (
	"context"
	"database/sql"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/merchant"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/google/uuid"
)

type MerchantService struct {
	merchantRepository merchant.Repository
}

func NewMerchantService(db *sql.DB) *MerchantService {
	mr := merchant.NewSQLiteRepository(db)

	return &MerchantService{merchantRepository: mr}
}

func (ms *MerchantService) CreateMerchant(ctx context.Context, name string) (dto.MerchantDto, error) {
	merchantDto := dto.MerchantDto{
		Name:       name,
		MerchantId: uuid.NewString(),
	}

	err := ms.merchantRepository.Create(ctx, merchantDto.ToEntity())
	if err != nil {
		return merchantDto, err
	}

	return merchantDto, nil
}

func (ms *MerchantService) GetMerchant(ctx context.Context, merchantId string) (dto.MerchantDto, error) {
	var merchantDto dto.MerchantDto

	merchant, err := ms.merchantRepository.GetByMerchantId(ctx, merchantId)
	if err != nil {
		return merchantDto, err
	}

	return merchantDto.FromEntity(merchant), nil
}
