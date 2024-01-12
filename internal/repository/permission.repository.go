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

type PermissionRepositoryUseCase interface {
	Create(ctx context.Context, permission *model.Permission) error
	Update(ctx context.Context, permission *model.Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*model.Permission, error)
	GetById(ctx context.Context, id int) (*model.Permission, error)
}

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepositoryUseCase {
	return &PermissionRepository{db}
}

func (r *PermissionRepository) Create(ctx context.Context, permission *model.Permission) error {
	if err := r.db.WithContext(ctx).Model(&model.Permission{}).Create(permission).Error; err != nil {
		return errors.Wrap(err, "error creating permission")
	}
	return nil
}

func (r *PermissionRepository) Update(ctx context.Context, permission *model.Permission) error {
	oldTime := permission.UpdatedAt
	newTime := time.Now()
	var txnError error

	txnError = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sourceModel := new(model.Permission)
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(sourceModel, permission.ID).Error; err != nil {
			log.Println("[PermissionRepository-Update]", err)
			return err
		}

		updates := sourceModel.MapUpdateFrom(permission)
		if len(*updates) > 0 {
			(*updates)["updated_at"] = newTime
			if err := tx.Model(&model.Permission{}).Where("id = ?", permission.ID).UpdateColumns(updates).Error; err != nil {
				log.Println("[PermissionRepository-Update]", err)
				return err
			}
		}

		return nil
	})

	if txnError != nil {
		permission.UpdatedAt = oldTime
		return txnError
	}

	permission.UpdatedAt = newTime
	return nil
}

func (r *PermissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&model.Permission{}).Delete(&model.Permission{}, id).Error; err != nil {
		return errors.Wrap(err, "error deleting permission")
	}
	return nil
}

func (r *PermissionRepository) GetAll(ctx context.Context) ([]*model.Permission, error) {
	var permissions []*model.Permission
	if err := r.db.WithContext(ctx).Model(&model.Permission{}).Find(&permissions).Error; err != nil {
		return nil, errors.Wrap(err, "error getting all permissions")
	}
	return permissions, nil
}

func (r *PermissionRepository) GetById(ctx context.Context, id int) (*model.Permission, error) {
	var permission model.Permission
	if err := r.db.WithContext(ctx).Model(&model.Permission{}).First(&permission, id).Error; err != nil {
		return nil, errors.Wrap(err, "error getting permission by id")
	}
	return &permission, nil
}
