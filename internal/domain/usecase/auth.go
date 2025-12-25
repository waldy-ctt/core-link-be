package usecase

import (
	"context"
)

type LoginInput struct {
	Email    string
	Password string
}

type SignupInput struct {
	Username    string
	Email       string
	Password    string
	DisplayName string
}

type LoginOutput struct {
	AccessToken  string
	RefreshToken string
}

type SignupOutput struct {
	AccessToken string
	RefreshToken string
}

type LoginUseCase interface {
	Execute(ctx context.Context, input LoginInput) (LoginOutput, error)
}

type SignupUseCase interface {
	Execute(ctx context.Context, input SignupInput) error
}
