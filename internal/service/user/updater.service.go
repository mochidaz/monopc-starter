package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"monopc-starter/internal/model"
	"monopc-starter/internal/repository"
	"monopc-starter/utils"
)

type UserUpdater struct {
	userRepo           repository.UserRepositoryUseCase
	roleRepo           repository.RoleRepositoryUseCase
	permissionRepo     repository.PermissionRepositoryUseCase
	userRoleRepo       repository.UserRoleRepositoryUseCase
	rolePermissionRepo repository.RolePermissionRepositoryUseCase
}

type UserUpdaterUseCase interface {
	UpdateUser(ctx context.Context, userID uuid.UUID, updateData *model.User) error
	UpdateRole(ctx context.Context, roleID int, updateData *model.Role) error
	UpdatePermission(ctx context.Context, permissionID int, updateData *model.Permission) error
	UpdatePassword(ctx context.Context, id uuid.UUID, oldPassword, newPassword string) error
}

func NewUserUpdater(
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
	permissionRepo repository.PermissionRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	rolePermissionRepo repository.RolePermissionRepositoryUseCase,
) UserUpdaterUseCase {
	return &UserUpdater{
		userRepo:           userRepo,
		roleRepo:           roleRepo,
		permissionRepo:     permissionRepo,
		userRoleRepo:       userRoleRepo,
		rolePermissionRepo: rolePermissionRepo,
	}
}

func (uu *UserUpdater) UpdateUser(ctx context.Context, userID uuid.UUID, updateData *model.User) error {
	user, err := uu.userRepo.GetById(ctx, userID)
	if err != nil {
		return err
	}

	return uu.userRepo.Update(ctx, user)
}

func (uu *UserUpdater) UpdateRole(ctx context.Context, roleID int, updateData *model.Role) error {
	role, err := uu.roleRepo.GetById(ctx, roleID)
	if err != nil {
		return err
	}
	return uu.roleRepo.Update(ctx, role)
}

func (uu *UserUpdater) UpdatePermission(ctx context.Context, permissionID int, updateData *model.Permission) error {
	permission, err := uu.permissionRepo.GetById(ctx, permissionID)
	if err != nil {
		return err
	}
	return uu.permissionRepo.Update(ctx, permission)
}

func (uu *UserUpdater) UpdatePassword(ctx context.Context, id uuid.UUID, oldPassword, newPassword string) error {
	user, err := uu.userRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if err := utils.ComparePassword(user.Password, oldPassword); err != nil {
		return errors.New("old password is not match")
	}

	hashedPassword, err := utils.HashPassword(newPassword)

	user.Password = hashedPassword

	return uu.userRepo.Update(ctx, user)
}
