package user

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

func (r *SQLiteRepository) Create(ctx context.Context, payload User) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO user (user_id, full_name, email) VALUES (?, ?, ?)",
		payload.UserId, payload.FullName, payload.Email,
	)

	return err
}

func (r *SQLiteRepository) GetByUserId(ctx context.Context, userId string) (User, error) {
	var user User

	row := r.db.QueryRowContext(ctx, "SELECT id, user_id, full_name, email FROM user WHERE user_id = ?", userId)
	err := row.Scan(&user.Id, &user.UserId, &user.FullName, &user.Email)

	if err != nil {
		return user, err
	}

	return user, nil
}
