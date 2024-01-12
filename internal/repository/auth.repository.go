package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"monopc-starter/internal/model"
)

type AuthRepositoryUseCase interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepositoryUseCase {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := ar.db.WithContext(ctx).Preload("UserRole.Role").Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, errors.Wrap(err, "error getting user by email")
	}

	return &user, nil
}
