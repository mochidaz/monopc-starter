package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"monopc-starter/internal/model"
)

type RolePermissionRepositoryUseCase interface {
	Create(ctx context.Context, rolePermission *model.RolePermission) error
	Update(ctx context.Context, rolePermission *model.RolePermission) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context) ([]*model.RolePermission, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.RolePermission, error)
}

type RolePermissionRepository struct {
	db *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB) RolePermissionRepositoryUseCase {
	return &RolePermissionRepository{db}
}

func (r *RolePermissionRepository) Create(ctx context.Context, rolePermission *model.RolePermission) error {
	if err := r.db.WithContext(ctx).Model(&model.RolePermission{}).Create(rolePermission).Error; err != nil {
		return errors.Wrap(err, "error creating role permission")
	}
	return nil
}

func (r *RolePermissionRepository) Update(ctx context.Context, rolePermission *model.RolePermission) error {
	if err := r.db.WithContext(ctx).Model(&model.RolePermission{}).Save(rolePermission).Error; err != nil {
		return errors.Wrap(err, "error updating role permission")
	}
	return nil
}

func (r *RolePermissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&model.RolePermission{}).Delete(&model.RolePermission{}, id).Error; err != nil {
		return errors.Wrap(err, "error deleting role permission")
	}
	return nil
}

func (r *RolePermissionRepository) GetAll(ctx context.Context) ([]*model.RolePermission, error) {
	var rolePermissions []*model.RolePermission
	if err := r.db.WithContext(ctx).Model(&model.RolePermission{}).Find(&rolePermissions).Error; err != nil {
		return nil, errors.Wrap(err, "error getting all role permissions")
	}
	return rolePermissions, nil
}

func (r *RolePermissionRepository) GetById(ctx context.Context, id uuid.UUID) (*model.RolePermission, error) {
	var rolePermission model.RolePermission
	if err := r.db.WithContext(ctx).Model(&model.RolePermission{}).First(&rolePermission, id).Error; err != nil {
		return nil, errors.Wrap(err, "error getting role permission by id")
	}
	return &rolePermission, nil
}
