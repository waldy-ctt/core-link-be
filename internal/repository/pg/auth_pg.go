package pg

import (
	"context"
	"database/sql"

	"github.com/waldy-ctt/core-link-be/internal/domain/entity"
)

func NewAuthRepo(db *sql.DB) *PostgreRepo {
	return &PostgreRepo{
		DB: db,
	}
}

func (r *PostgreRepo) CreateAuth(ctx context.Context, u *entity.Auth) error {
	query := `
		INSERT INTO auth (userid, email, password) VALUES ($1, $2, $3)
	`

	_, err := r.DB.ExecContext(ctx, query, u.UserId, u.Email, u.PasswordHash)
	return err
}
