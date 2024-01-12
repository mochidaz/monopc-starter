package user

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"monopc-starter/internal/app/config"
	"monopc-starter/internal/model"
	"monopc-starter/internal/repository"
	"monopc-starter/utils"
	"time"
)

type UserCreator struct {
	cfg                config.Config
	userRepo           repository.UserRepositoryUseCase
	roleRepo           repository.RoleRepositoryUseCase
	userRoleRepo       repository.UserRoleRepositoryUseCase
	permissionRepo     repository.PermissionRepositoryUseCase
	rolePermissionRepo repository.RolePermissionRepositoryUseCase
}

type UserCreatorUseCase interface {
	CreateUser(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time) (*model.User, error)

	CreateAdmin(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time, roleID int) (*model.User, error)

	CreateRole(ctx context.Context, name string, permissionIDs []int, createdBy string) (*model.Role, error)

	CreatePermission(ctx context.Context, name string, label string) (*model.Permission, error)

	//RegisterUser(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time) (*model.User, error)
}

func NewUserCreator(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	permissionRepo repository.PermissionRepositoryUseCase,
	rolePermissionRepo repository.RolePermissionRepositoryUseCase,
) UserCreatorUseCase {
	return &UserCreator{
		cfg:                cfg,
		userRepo:           userRepo,
		roleRepo:           roleRepo,
		userRoleRepo:       userRoleRepo,
		permissionRepo:     permissionRepo,
		rolePermissionRepo: rolePermissionRepo,
	}
}

func (uc *UserCreator) CreateUser(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time) (*model.User, error) {

	pass, err := utils.HashPassword(password)

	if err != nil {
		return nil, err
	}

	user := model.NewUser(uuid.New(), name, email, pass, phoneNumber, photo, sql.NullTime{Time: dob}, "system")

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserCreator) CreateAdmin(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time, roleID int) (*model.User, error) {
	user := model.NewUser(uuid.New(), name, email, password, phoneNumber, photo, sql.NullTime{Time: dob}, "system")

	user.UserRole = model.NewUserRole(uuid.New(), user.ID, roleID, "system")

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	userRole := model.NewUserRole(uuid.New(), user.ID, roleID, "system")
	if err := uc.userRoleRepo.Create(ctx, userRole); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserCreator) CreateRole(ctx context.Context, name string, permissionIDs []int, createdBy string) (*model.Role, error) {
	role := model.NewRole(name, createdBy)

	roleID, err := uc.roleRepo.Create(ctx, role)

	if err != nil {
		return nil, err
	}

	for _, permissionID := range permissionIDs {
		rolePermission := model.NewRolePermission(uuid.New(), roleID, permissionID, createdBy)
		if err := uc.rolePermissionRepo.Create(ctx, rolePermission); err != nil {
			return nil, err
		}
	}

	return role, nil
}

func (uc *UserCreator) CreatePermission(ctx context.Context, name string, label string) (*model.Permission, error) {
	permission := model.NewPermission(name, label, "system")
	if err := uc.permissionRepo.Create(ctx, permission); err != nil {
		return nil, err
	}
	return permission, nil
}
