package model

import "time"

type Role struct {
	ID              int `json:"id" gorm:"primary_key;auto_increment"`
	Name            string
	RolePermissions []*RolePermission `gorm:"foreignKey:RoleID"`
	Auditable
}

func NewRole(
	name string,
	createdBy string,
) *Role {
	return &Role{
		Name:      name,
		Auditable: NewAuditable(createdBy),
	}
}

func (model *Role) MapUpdateFrom(from *Role) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":            model.Name,
			"role_permission": model.RolePermissions,
			"updated_at":      model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	for i := range model.RolePermissions {
		if model.RolePermissions[i].PermissionID != from.RolePermissions[i].PermissionID {
			mapped["role_permission"] = from.RolePermissions
		}
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
