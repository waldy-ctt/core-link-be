package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/waldy-ctt/core-link-be/internal/domain/repo"
	"github.com/waldy-ctt/core-link-be/internal/domain/usecase"
)

type loginUseCase struct {
	authRepo       repo.AuthRepository
	contextTimeout time.Duration
}

func NewLoginUseCase(a repo.AuthRepository, timeout time.Duration) usecase.LoginUseCase {
	return &loginUseCase{
		authRepo:       a,
		contextTimeout: timeout,
	}
}

func (uc *loginUseCase) Execute(c context.Context, input usecase.LoginInput) (usecase.LoginOutput, error) {
	ctx, cancel := context.WithTimeout(c, uc.contextTimeout)
	defer cancel()

	auth, err := uc.authRepo.GetAuthByEmail(ctx, input.Email)
	if err != nil {
		return usecase.LoginOutput{}, fmt.Errorf("Invalid Credentials")
	}

	if auth.PasswordHash != input.Password {
		return usecase.LoginOutput{}, fmt.Errorf("Invalid Credentials")
	}

	return usecase.LoginOutput{
		AccessToken:  "asd",
		RefreshToken: "asd",
	}, nil
}
