package pg

import (
	"context"
	"database/sql"

	"github.com/waldy-ctt/core-link-be/internal/domain/entity"
)

type PostgreRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *PostgreRepo {
	return &PostgreRepo{
		DB: db,
	}
}

func (r *PostgreRepo) CreateUser(ctx context.Context, u *entity.User) error {
	query := `
		INSERT INTO users (id, username, display_name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.DB.ExecContext(ctx, query, u.ID, u.Username, u.DisplayName, u.CreatedAt, u.UpdatedAt)
	return err
}

