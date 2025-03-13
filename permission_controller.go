package controllers

import (
	"errors"
	"icdn/constants"
	"icdn/handlers"
	"icdn/middleware"
	"icdn/services"
	"icdn/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePermission(context *gin.Context) {
	req := middleware.GetRequestData[types.SavePermissionTypes](context, "validatedData")

	createdPermission, err := services.SavePermission(req)
	if err != nil {
		handlers.ErrorResponse(context, http.StatusBadRequest, errors.New(constants.ErrSomeThingWrong), err)
		return
	}

	handlers.SuccessResponse(context, http.StatusOK, createdPermission, constants.SuccessPermissionCreated)
}

func GetPermissionByID(context *gin.Context) {
	req := middleware.GetRequestData[types.PermissionID](context, "validatedData")

	permission, statusCode, err := services.GetPermissionByID(req.ID)
	if err != nil {
		errorMessage := map[bool]string{true: constants.ErrSomeThingWrong, false: constants.ErrPermissionNotFound}[statusCode == 0]
		handlers.ErrorResponse(context, statusCode, errors.New(errorMessage), err)
		context.Abort()
		return
	}

	handlers.SuccessResponse(context, statusCode, permission, constants.SuccessPermissionFetch)
}

func GetAllPermissions(context *gin.Context) {
	req := middleware.GetRequestData[types.PaginationRequest](context, "validatedData")

	page := req.Page
	limit := req.Limit

	permissions, total, err := services.GetAllPermissions(req)
	if err != nil {
		handlers.ErrorResponse(context, http.StatusBadRequest, errors.New(constants.ErrSomeThingWrong), err)
		return
	}

	response := gin.H{
		"data":       permissions,
		"total":      total,
		"totalPages": (int(total) + limit - 1) / limit,
		"page":       page,
		"limit":      limit,
	}

	handlers.SuccessResponse(context, http.StatusOK, response, constants.SuccessPermissionFetch)
}

func UpdatePermission(context *gin.Context) {
	req := middleware.GetRequestData[types.UpdatePermissionTypes](context, "validatedData")

	updatePermission, err := services.UpdatePermission(req)
	if err != nil {
		handlers.ErrorResponse(context, http.StatusBadRequest, errors.New(constants.ErrSomeThingWrong), err)
		return
	}

	handlers.SuccessResponse(context, http.StatusOK, updatePermission, constants.SuccessPermissionUpdate)
}

func DeletePermission(context *gin.Context) {
	req := middleware.GetRequestData[types.PermissionID](context, "validatedData")

	delPermission, statusCode, err := services.DeletePermission(req.ID)
	if err != nil {
		handlers.ErrorResponse(context, statusCode, errors.New(constants.ErrPermissionNotFound), err)
		return
	}

	handlers.SuccessResponse(context, statusCode, delPermission, constants.SuccessPermissionDelete)
}
