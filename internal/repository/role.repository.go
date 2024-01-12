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

type RoleRepositoryUseCase interface {
	Create(ctx context.Context, role *model.Role) (int, error)
	Update(ctx context.Context, role *model.Role) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Role, error)
	GetById(ctx context.Context, id int) (*model.Role, error)
}

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepositoryUseCase {
	return &RoleRepository{db}
}

func (r *RoleRepository) Create(ctx context.Context, role *model.Role) (int, error) {
	if err := r.db.WithContext(ctx).Model(&model.Role{}).Create(role).Error; err != nil {
		return 0, errors.Wrap(err, "error creating role")
	}
	return role.ID, nil
}

func (r *RoleRepository) Update(ctx context.Context, role *model.Role) error {
	oldTime := role.UpdatedAt
	newTime := time.Now()
	var txnError error

	txnError = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		sourceModel := new(model.Role)
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(sourceModel, role.ID).Error; err != nil {
			log.Println("[RoleRepository-Update]", err)
			return err
		}

		updates := sourceModel.MapUpdateFrom(role)
		if len(*updates) > 0 {
			(*updates)["updated_at"] = newTime
			if err := tx.Model(&model.Role{}).Where("id = ?", role.ID).UpdateColumns(updates).Error; err != nil {
				log.Println("[RoleRepository-Update]", err)
				return err
			}
		}

		return nil
	})

	if txnError != nil {
		role.UpdatedAt = oldTime
		return txnError
	}

	role.UpdatedAt = newTime
	return nil
}

func (r *RoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Model(&model.Role{}).Delete(&model.Role{}, id).Error; err != nil {
		return errors.Wrap(err, "error deleting role")
	}
	return nil
}

func (repo *RoleRepository) GetAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*model.Role, error) {
	var roles []*model.Role
	q := repo.db.Preload("RolePermissions").WithContext(ctx).Model(&model.Role{})
	if query != "" {
		q = q.Where("name LIKE ?", "%"+query+"%")
	}
	if sort != "" {
		q = q.Order(sort + " " + order)
	}
	if limit > 0 {
		q = q.Limit(limit).Offset(offset)
	}
	if err := q.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) GetById(ctx context.Context, id int) (*model.Role, error) {
	var role model.Role
	if err := r.db.Preload("RolePermissions").WithContext(ctx).Model(&model.Role{}).First(&role, id).Error; err != nil {
		return nil, errors.Wrap(err, "error getting role by id")
	}
	return &role, nil
}
