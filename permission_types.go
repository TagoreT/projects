package types

import (
	"time"

	"github.com/google/uuid"
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

// Request
type SavePermissionTypes struct {
	PermissionName string `json:"permission_name" validate:"required"`
	PermissionType string `json:"permission_type" validate:"required"`
	AccessLevel    string `json:"access_level" validate:"required"`
	Description    string `json:"description"`
	CreatedBy      string
}

type PermissionID struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type UpdatePermissionTypes struct {
	ID             uuid.UUID `json:"id" validate:"required"`
	PermissionName string    `json:"permission_name"`
	PermissionType string    `json:"permission_type"`
	Description    string    `json:"description"`
	AccessLevel    string    `json:"access_level"`
	UpdatedBy      string
}

type PaginationRequest struct {
	Search string `json:"search"`
	Page   int    `json:"page" validate:"required"`
	Limit  int    `json:"limit" validate:"required"`
}

// Response
type PermissionData struct {
	ID             uuid.UUID `json:"_id"`
	PermissionName string    `json:"permission_name" validate:"required"`
	PermissionType string    `json:"permission_type" validate:"required"`
	AccessLevel    string    `json:"access_level" validate:"required"`
	Description    string    `json:"description"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
