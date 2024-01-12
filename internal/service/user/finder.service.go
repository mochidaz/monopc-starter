package user

import (
	"context"
	"github.com/google/uuid"
	"monopc-starter/internal/model"
	"monopc-starter/internal/repository"
)

type UserFinder struct {
	userRepo           repository.UserRepositoryUseCase
	roleRepo           repository.RoleRepositoryUseCase
	permissionRepo     repository.PermissionRepositoryUseCase
	userRoleRepo       repository.UserRoleRepositoryUseCase
	rolePermissionRepo repository.RolePermissionRepositoryUseCase
}

type UserFinderUseCase interface {
	FindUser(ctx context.Context, userID uuid.UUID) (*model.User, error)

	FindRole(ctx context.Context, roleID int) (*model.Role, error)

	FindPermission(ctx context.Context, permissionID int) (*model.Permission, error)

	FindAllUser(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.User, error)

	FindAllRole(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Role, error)

	FindAllPermission(ctx context.Context) ([]*model.Permission, error)
}

func NewUserFinder(
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
	permissionRepo repository.PermissionRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	rolePermissionRepo repository.RolePermissionRepositoryUseCase,
) UserFinderUseCase {
	return &UserFinder{
		userRepo:           userRepo,
		roleRepo:           roleRepo,
		permissionRepo:     permissionRepo,
		userRoleRepo:       userRoleRepo,
		rolePermissionRepo: rolePermissionRepo,
	}
}

func (uf *UserFinder) FindUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	user, err := uf.userRepo.GetById(ctx, userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uf *UserFinder) FindRole(ctx context.Context, roleID int) (*model.Role, error) {
	role, err := uf.roleRepo.GetById(ctx, roleID)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (uf *UserFinder) FindPermission(ctx context.Context, permissionID int) (*model.Permission, error) {
	permission, err := uf.permissionRepo.GetById(ctx, permissionID)

	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (uf *UserFinder) FindAllUser(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.User, error) {
	users, err := uf.userRepo.GetAll(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uf *UserFinder) FindAllRole(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Role, error) {
	roles, err := uf.roleRepo.GetAll(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (uf *UserFinder) FindAllPermission(ctx context.Context) ([]*model.Permission, error) {
	permissions, err := uf.permissionRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
