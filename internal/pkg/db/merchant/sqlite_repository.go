package merchant

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

func (r *SQLiteRepository) Create(ctx context.Context, payload Merchant) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO merchant (name, merchant_id) VALUES (?, ?)",
		payload.Name, payload.MerchantId,
	)

	return err
}

func (r *SQLiteRepository) GetByMerchantId(ctx context.Context, merchantId string) (Merchant, error) {
	var merchant Merchant

	row := r.db.QueryRowContext(ctx, "SELECT id, name, merchant_id FROM merchant WHERE merchant_id = ?", merchantId)
	err := row.Scan(&merchant.Id, &merchant.Name, &merchant.MerchantId)

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}
