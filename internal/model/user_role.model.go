package model

import (
	"github.com/google/uuid"
	"time"
)

const (
	userRoleTableName = "user_roles"
)

type UserRole struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	RoleID int       `json:"role_id"`
	Role   *Role     `json:"role"`
	User   *User     `json:"user"`
	Auditable
}

func (*UserRole) TableName() string {
	return userRoleTableName
}

func NewUserRole(
	id uuid.UUID,
	userID uuid.UUID,
	roleID int,
	createdBy string,
) *UserRole {
	return &UserRole{
		ID:        id,
		UserID:    userID,
		RoleID:    roleID,
		Auditable: NewAuditable(createdBy),
	}
}

func (model *UserRole) MapUpdateFrom(from *UserRole) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"role_id":    model.RoleID,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.RoleID != from.RoleID {
		mapped["role_id"] = from.RoleID
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
