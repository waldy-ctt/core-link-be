package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/waldy-ctt/core-link-be/internal/domain/entity"
	"github.com/waldy-ctt/core-link-be/internal/domain/repo"
)

type authRepo struct {
	DB *sql.DB
}

func NewAuthRepo(db *sql.DB) repo.AuthRepository {
	return &authRepo{
		DB: db,
	}
}

func (r *authRepo) CreateAuth(ctx context.Context, u *entity.Auth) error {
	query := `
		INSERT INTO auth (user_id, email, password_hash) VALUES ($1, $2, $3)
	`
	_, err := r.DB.ExecContext(ctx, query, u.UserID, u.Email, u.PasswordHash)
	return err
}

func (r *authRepo) GetAuthByEmail(ctx context.Context, email string) (*entity.Auth, error) {
	query := `SELECT user_id, email, password_hash FROM auth WHERE email = $1`

	var auth entity.Auth
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&auth.UserID, &auth.Email, &auth.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("auth not found")
		}
		return nil, err
	}
	return &auth, nil
}

func (r *authRepo) SaveRefreshToken(ctx context.Context, userID, token string) error {
	return nil
}
