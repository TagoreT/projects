package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionType string

const (
	Brand    PermissionType = "Brand"
	User     PermissionType = "User"
	Approval PermissionType = "Approval"
)

type AccessLevelType string

const (
	Public  AccessLevelType = "Public"
	Private AccessLevelType = "Private"
)

type Permissions struct {
	ID             uuid.UUID       `gorm:"type:uuid; default:uuid_generate_v4();primaryKey" json:"_id"`
	PermissionName string          `gorm:"not null" json:"permission_name"`
	PermissionType PermissionType  `gorm:"type:varchar(20);not null" json:"permission_type"`
	PermissionText string          `json:"permissions_text"`
	Description    string          `gorm:"type:text" json:"description"`
	AccessLevel    AccessLevelType `gorm:"type:varchar(20);not null" json:"access_level"`
	CreatedBy      string          `gorm:"type:text;not null" json:"created_by"`
	CreatedAt      time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedBy      *uuid.UUID      `json:"updated_by"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `json:"deleted_at"`
}
