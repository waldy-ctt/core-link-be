package pg

import (
	"context"
	"database/sql"

	"github.com/waldy-ctt/core-link-be/internal/domain/entity"
	"github.com/waldy-ctt/core-link-be/internal/domain/repo"
)

type userRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepository {
	return &userRepo{
		DB: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *entity.User) error {
	query := `
		INSERT INTO users (id, username, display_name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.DB.ExecContext(ctx, query, u.ID, u.Username, u.DisplayName, u.CreatedAt, u.UpdatedAt)
	return err
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return nil, nil
}
