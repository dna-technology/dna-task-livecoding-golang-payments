package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/db/payment"
	"github.com/dna-technology/dna-task-livecoding-golang/internal/pkg/dto"
	"github.com/google/uuid"
)

type PaymentService struct {
	paymentRepository payment.Repository
	accountService    *AccountService
	userService       *UserService
	merchantService   *MerchantService
}

func NewPaymentService(db *sql.DB) *PaymentService {
	pr := payment.NewSQLiteRepository(db)
	as := NewAccountService(db)
	us := NewUserService(db)
	ms := NewMerchantService(db)

	return &PaymentService{
		paymentRepository: pr,
		accountService:    as,
		userService:       us,
		merchantService:   ms,
	}
}

func (ps *PaymentService) AddPayment(ctx context.Context, payload dto.PaymentDto) (dto.PaymentDto, error) {
	paymentDto := dto.PaymentDto{
		PaymentId:  uuid.NewString(),
		UserId:     payload.UserId,
		Amount:     payload.Amount,
		MerchantId: payload.MerchantId,
	}

	_, err := ps.userService.GetUser(ctx, paymentDto.UserId)
	if err != nil {
		return paymentDto, err
	}

	_, err = ps.merchantService.GetMerchant(ctx, paymentDto.MerchantId)
	if err != nil {
		return paymentDto, err
	}

	a, err := ps.accountService.GetUserAccount(ctx, paymentDto.UserId)
	if err != nil {
		return paymentDto, err
	}

	if a.Balance < paymentDto.Amount {
		return paymentDto, errors.New("insufficient funds")
	}

	a, err = ps.accountService.DecreaseBalance(ctx, a.AccountId, paymentDto.Amount)
	if err != nil {
		return paymentDto, err
	}

	err = ps.paymentRepository.Create(ctx, paymentDto.ToEntity())
	if err != nil {
		return paymentDto, err
	}

	return paymentDto, err
}
