package services

import (
	"errors"
	"fmt"
	"icdn/constants"
	"icdn/database"
	"icdn/models"
	"icdn/types"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// Save Permission
func SavePermission(req *types.SavePermissionTypes) (types.PermissionData, error) {
	permission := models.Permissions{}
	response := types.PermissionData{}
	copier.Copy(&permission, req)

	result := database.DB.Create(&permission)

	if result.Error != nil {
		return response, fmt.Errorf("%s: %v", constants.ErrCreatingPermissions, result.Error)
	}
	copier.Copy(&response, &permission)
	return response, nil
}

// Get Permission by ID
func GetPermissionByID(id uuid.UUID) (*types.PermissionData, int, error) {
	permission := models.Permissions{}
	response := types.PermissionData{}

	result := database.DB.First(&permission, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New(constants.ErrPermissionNotFound)
		}
		return nil, 0, errors.New("database error: " + result.Error.Error())
	}
	copier.Copy(&response, &permission)
	return &response, http.StatusOK, nil
}

// Get All Permission
func GetAllPermissions(req *types.PaginationRequest) ([]types.PermissionData, int64, error) {
	permissions := models.Permissions{}
	records := []types.PermissionData{}
	total := int64(0)
	searchFields := []string{"permission_name", "permission_type", "access_level"}

	query := database.DB.Model(permissions)

	if req.Search != "" {
		searchPattern := "%" + strings.ToLower(req.Search) + "%"
		for _, field := range searchFields {
			query = query.Or(fmt.Sprintf("LOWER(%s) ILIKE LOWER(?)", field), searchPattern)
		}
	}

	totalCount := query.Count(&total)
	if totalCount.Error != nil {
		return nil, 0, errors.New(constants.ErrFetchingCount + ": " + totalCount.Error.Error())
	}

	offset := (req.Page - 1) * req.Limit
	result := query.Limit(req.Limit).Offset(offset).Find(&records)
	if result.Error != nil {
		return nil, 0, errors.New(constants.ErrFetchingPermissions + ": " + result.Error.Error())
	}
	return records, total, nil
}

// update Permission
func UpdatePermission(req *types.UpdatePermissionTypes) (*types.PermissionData, error) {
	permission := models.Permissions{}
	response := types.PermissionData{}
	copier.Copy(&permission, req)

	result := database.DB.Model(&permission).Where("id = ?", req.ID).Updates(req)
	if result.Error != nil {
		return nil, errors.New(constants.ErrUpdatingPermissions + result.Error.Error())
	}
	copier.Copy(&response, &permission)
	return &response, nil
}

// Delete Permissionx
func DeletePermission(id uuid.UUID) (string, int, error) {
	permission := models.Permissions{}

	result := database.DB.Where("id = ?", id).Delete(&permission)
	response := fmt.Sprintf("Rows Affected: %d", result.RowsAffected)
	if result.Error != nil {
		return response, 0, errors.New("Database Error :" + result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return response, http.StatusNotFound, errors.New(constants.ErrPermissionNotFound)
	}
	return response, http.StatusOK, nil
}

func PermissionDuplicateCheck(permissionName string, id uuid.UUID) bool {
	permission := models.Permissions{}
	query := database.DB

	if permissionName != "" {
		query = query.Where("permission_name = ?", permissionName)
	}
	if id != uuid.Nil {
		query = query.Where("id != ? AND deleted_at IS NULL", id)
	}

	result := query.First(&permission)
	return result.Error == nil
}
