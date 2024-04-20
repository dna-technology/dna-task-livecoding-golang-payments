package account

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

func (r *SQLiteRepository) Create(ctx context.Context, payload Account) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO account (account_id, user_id, balance) VALUES (?, ?, ?)",
		payload.AccountId, payload.UserId, payload.Balance,
	)
	return err
}

func (r *SQLiteRepository) GetByAccountId(ctx context.Context, accountId string) (Account, error) {
	var account Account

	row := r.db.QueryRowContext(ctx, "SELECT id, account_id, user_id, balance FROM account WHERE account_id = ?", accountId)
	err := row.Scan(&account.Id, &account.AccountId, &account.UserId, &account.Balance)

	if err != nil {
		return account, err
	}

	return account, nil
}

func (r *SQLiteRepository) GetByUserId(ctx context.Context, userId string) (Account, error) {
	var account Account

	row := r.db.QueryRowContext(ctx, "SELECT id, account_id, user_id, balance FROM account WHERE user_id = ?", userId)
	err := row.Scan(&account.Id, &account.AccountId, &account.UserId, &account.Balance)

	if err != nil {
		return account, err
	}

	return account, nil
}
