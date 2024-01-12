package model

import (
	"database/sql"
	"gorm.io/gorm"
	"monopc-starter/utils"
	"time"
)

type Auditable struct {
	CreatedBy sql.NullString `json:"created_by"`
	UpdatedBy sql.NullString `json:"updated_by"`
	DeletedBy sql.NullString `json:"deleted_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func NewAuditable(createdBy string) Auditable {
	return Auditable{
		CreatedBy: utils.StringToNullString(createdBy),
		UpdatedBy: utils.StringToNullString(createdBy),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
