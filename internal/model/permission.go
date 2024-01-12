package model

import (
	"time"
)

const (
	permissionTableName = "permissions"
)

type Permission struct {
	ID    int    `json:"id" gorm:"primary_key;auto_increment"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Auditable
}

func (model *Permission) TableName() string {
	return permissionTableName
}

func NewPermission(
	name string,
	label string,
	createdBy string,
) *Permission {
	return &Permission{
		Name:      name,
		Label:     label,
		Auditable: NewAuditable(createdBy),
	}
}

func (model *Permission) MapUpdateFrom(from *Permission) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":       model.Name,
			"created_by": model.CreatedBy,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	if model.Label != from.Label {
		mapped["label"] = from.Label
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
