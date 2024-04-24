package payment

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Create(ctx context.Context, payload Payment) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO payment (payment_id, user_id, merchant_id, amount) VALUES (?, ?, ?, ?)",
		payload.PaymentId, payload.UserId, payload.MerchantId, payload.Amount,
	)

	return err
}

func (r *SQLiteRepository) GetByPaymentId(ctx context.Context, paymentId string) (Payment, error) {
	var payment Payment

	row := r.db.QueryRowContext(ctx, "SELECT id, payment_id, user_id, merchant_id, amount FROM payment WHERE payment_id = ?", paymentId)
	err := row.Scan(&payment.Id, &payment.PaymentId, &payment.UserId, &payment.MerchantId, &payment.Amount)

	if err != nil {
		return payment, err
	}

	return payment, nil
}
