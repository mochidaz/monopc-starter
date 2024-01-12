package model

import (
	"github.com/google/uuid"
)

const (
	rolePermissionTableName = "role_permissions"
)

type RolePermission struct {
	ID           uuid.UUID `json:"id"`
	RoleID       int       `json:"role_id"`
	PermissionID int       `json:"permission_id"`
	Permission   *Permission
	Auditable
}

func (model *RolePermission) TableName() string {
	return rolePermissionTableName
}

func NewRolePermission(
	id uuid.UUID,
	roleID int,
	permissionID int,
	createdBy string,
) *RolePermission {
	return &RolePermission{
		ID:           id,
		RoleID:       roleID,
		PermissionID: permissionID,
		Auditable:    NewAuditable(createdBy),
	}
}
