package user

import (
	"context"
	"github.com/google/uuid"
	"monopc-starter/internal/repository"
)

type UserDeleter struct {
	userRepo           repository.UserRepositoryUseCase
	userRoleRepo       repository.UserRoleRepositoryUseCase
	roleRepo           repository.RoleRepositoryUseCase
	rolePermissionRepo repository.RolePermissionRepositoryUseCase
	permissionRepo     repository.PermissionRepositoryUseCase
}

type UserDeleterUseCase interface {
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	DeleteRole(ctx context.Context, roleID uuid.UUID) error
	DeletePermission(ctx context.Context, permissionID uuid.UUID) error
}

func NewUserDeleter(
	userRepo repository.UserRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
	rolePermissionRepo repository.RolePermissionRepositoryUseCase,
	permissionRepo repository.PermissionRepositoryUseCase,
) UserDeleterUseCase {
	return &UserDeleter{
		userRepo:           userRepo,
		userRoleRepo:       userRoleRepo,
		roleRepo:           roleRepo,
		rolePermissionRepo: rolePermissionRepo,
		permissionRepo:     permissionRepo,
	}
}

func (ud *UserDeleter) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	if err := ud.userRoleRepo.Delete(ctx, userID); err != nil {
		return err
	}
	return ud.userRepo.Delete(ctx, userID)
}

func (ud *UserDeleter) DeleteRole(ctx context.Context, roleID uuid.UUID) error {
	if err := ud.rolePermissionRepo.Delete(ctx, roleID); err != nil {
		return err
	}
	return ud.roleRepo.Delete(ctx, roleID)
}

func (ud *UserDeleter) DeletePermission(ctx context.Context, permissionID uuid.UUID) error {
	return ud.permissionRepo.Delete(ctx, permissionID)
}
