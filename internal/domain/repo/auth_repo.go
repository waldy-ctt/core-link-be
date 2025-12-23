package repo

import (
	"context"

	"github.com/waldy-ctt/core-link-be/internal/domain/entity"
)

type AuthRepository interface {
    // Save the credentials (email + password_hash)
    CreateAuth(ctx context.Context, auth *entity.Auth) error 
    
    // Get credentials so the UseCase can check the password
    GetAuthByEmail(ctx context.Context, email string) (*entity.Auth, error)
    
    // If you strictly want to store refresh tokens in DB
    SaveRefreshToken(ctx context.Context, userID, token string) error
}
