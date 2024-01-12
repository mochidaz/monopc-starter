package auth

import (
	"context"
	"monopc-starter/internal/app/config"
	"monopc-starter/internal/repository"
	"monopc-starter/utils"
)

type AuthServiceUseCase interface {
	Login(ctx context.Context, email, password string) (*string, error)

	Verify(ctx context.Context, token string) (bool, error)
}

type AuthService struct {
	cfg      config.Config
	userRepo repository.UserRepositoryUseCase
	authRepo repository.AuthRepositoryUseCase
}

func NewAuthService(config config.Config, userRepo repository.UserRepositoryUseCase, authRepo repository.AuthRepositoryUseCase) AuthServiceUseCase {
	return &AuthService{
		cfg:      config,
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (*string, error) {
	user, err := as.authRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := utils.ComparePassword(user.Password, password); err != nil {
		return nil, err
	}

	iss := "user"

	if user.UserRole.Role.Name == "admin" {
		iss = "cms"
	}

	jwtToken, err := utils.JWTEncode(as.cfg, user.ID, iss)

	if err != nil {
		return nil, err
	}

	return &jwtToken, nil
}

func (as *AuthService) Verify(ctx context.Context, token string) (bool, error) {
	_, err := utils.JWTDecode(as.cfg, token)

	if err != nil {
		return false, err
	}

	return true, nil
}
