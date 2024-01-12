package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"monopc-starter/internal/model"
	"time"
)

type UserRoleRepositoryUseCase interface {
	Create(ctx context.Context, userRole *model.UserRole) error
	Update(ctx context.Context, userRole *model.UserRole) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*model.UserRole, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.UserRole, error)
}

type UserRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepositoryUseCase {
	return &UserRoleRepository{db}
}

func (r *UserRoleRepository) Create(ctx context.Context, userRole *model.UserRole) error {
	if err := r.db.WithContext(ctx).Model(&model.UserRole{}).Create(userRole).Error; err != nil {
		return errors.Wrap(err, "error creating user role")
	}
	return nil
}

func (r *UserRoleRepository) Update(ctx context.Context, userRole *model.UserRole) error {
	oldTime := userRole.UpdatedAt
	newTime := time.Now()
	var txnError error

	txnError = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sourceModel := new(model.UserRole)
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(sourceModel, userRole.ID).Error; err != nil {
			log.Println("[UserRoleRepository-Update]", err)
			return err
		}

		updates := sourceModel.MapUpdateFrom(userRole)
		if len(*updates) > 0 {
			(*updates)["updated_at"] = newTime
			if err := tx.Model(&model.UserRole{}).Where("id = ?", userRole.ID).UpdateColumns(updates).Error; err != nil {
				log.Println("[UserRoleRepository-Update]", err)
				return err
			}
		}

		return nil
	})

	if txnError != nil {
		userRole.UpdatedAt = oldTime
		return txnError
	}

	userRole.UpdatedAt = newTime
	return nil
}

func (r *UserRoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&model.UserRole{}).Delete(&model.UserRole{}, id).Error; err != nil {
		return errors.Wrap(err, "error deleting user role")
	}
	return nil
}

func (r *UserRoleRepository) GetAll(ctx context.Context) ([]*model.UserRole, error) {
	var userRoles []*model.UserRole
	if err := r.db.WithContext(ctx).Model(&model.UserRole{}).Find(&userRoles).Error; err != nil {
		return nil, errors.Wrap(err, "error getting all user roles")
	}
	return userRoles, nil
}

func (r *UserRoleRepository) GetById(ctx context.Context, id uuid.UUID) (*model.UserRole, error) {
	var userRole model.UserRole
	if err := r.db.WithContext(ctx).Model(&model.UserRole{}).First(&userRole, id).Error; err != nil {
		return nil, errors.Wrap(err, "error getting user role by id")
	}
	return &userRole, nil
}
