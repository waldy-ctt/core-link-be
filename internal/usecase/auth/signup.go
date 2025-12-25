package auth

import (
	"context"
	"time"

	"github.com/waldy-ctt/core-link-be/internal/domain/entity"
	"github.com/waldy-ctt/core-link-be/internal/domain/repo"
	"github.com/waldy-ctt/core-link-be/internal/domain/usecase"
	"github.com/waldy-ctt/core-link-be/internal/platform/hasher"
	"github.com/waldy-ctt/core-link-be/internal/platform/idgen"
)

type signupUseCase struct {
	userRepo       repo.UserRepository
	authRepo       repo.AuthRepository
	contextTimeout time.Duration
}

func NewSignupUseCase(u repo.UserRepository, a repo.AuthRepository, timeout time.Duration) usecase.SignupUseCase {
	return &signupUseCase{
		userRepo:       u,
		authRepo:       a,
		contextTimeout: timeout,
	}
}

func (uc *signupUseCase) Execute(c context.Context, input usecase.SignupInput) error {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	// 1. Generate ID & Hash
	newID, _ := idgen.GenerateV4()
	hashedPwd, _ := hasher.HashPassword(input.Password)

	user := &entity.User{
		ID:          newID,
		Username:    input.Username,
		DisplayName: input.DisplayName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	auth := &entity.Auth{
		UserID:       newID,
		Email:        input.Email,
		PasswordHash: hashedPwd,
	}

	if err := uc.userRepo.CreateUser(ctx, user); err != nil {
		return err
	}
	return uc.authRepo.CreateAuth(ctx, auth)
}
