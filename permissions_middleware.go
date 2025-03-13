package middleware

import (
	"errors"
	"icdn/constants"
	"icdn/handlers"
	"icdn/services"
	"icdn/types"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	validPermissionTypes = map[types.PermissionType]bool{
		types.Brand: true, types.User: true, types.Approval: true,
	}
	validAccessLevels = map[types.AccessLevelType]bool{
		types.Public: true, types.Private: true,
	}
)

// save, update and delete permission genaric validator
func PermissionValidator[modelType interface{}]() gin.HandlerFunc {
	return func(context *gin.Context) {
		req := GetRequestData[modelType](context, "validatedData")
		reflectRequest := reflect.Indirect(reflect.ValueOf(req))

		permissionName := reflectRequest.FieldByName("PermissionName")
		permissionType := types.PermissionType(reflectRequest.FieldByName("PermissionType").String())
		accessLevel := types.AccessLevelType(reflectRequest.FieldByName("AccessLevel").String())

		permissionID := uuid.Nil
		if idField := reflectRequest.FieldByName("ID"); idField.IsValid() && idField.Type() == reflect.TypeOf(uuid.UUID{}) {
			permissionID = idField.Interface().(uuid.UUID)
			permission, statusCode, err := services.GetPermissionByID(permissionID)
			if err != nil || permission == nil {
				errorMessage := map[bool]string{true: constants.ErrSomeThingWrong, false: constants.ErrPermissionNotFound}[statusCode == 0]
				handlers.ErrorResponse(context, statusCode, errors.New(errorMessage), err)
				context.Abort()
				return
			}
		}

		if services.PermissionDuplicateCheck(permissionName.String(), permissionID) {
			handlers.ErrorResponse(context, http.StatusConflict, errors.New(constants.ErrPermissionExists))
			context.Abort()
			return
		}

		if !ValidatePermissionTypes(permissionType, accessLevel, context) {
			context.Abort()
			return
		}

		context.Next()
	}
}

func ValidatePermissionTypes(PermitionType types.PermissionType, AccessLevel types.AccessLevelType, context *gin.Context) bool {
	if !validPermissionTypes[PermitionType] {
		handlers.ErrorResponse(context, http.StatusBadRequest, errors.New(constants.ErrPermissionType))
		return false
	}

	if !validAccessLevels[AccessLevel] {
		handlers.ErrorResponse(context, http.StatusBadRequest, errors.New(constants.ErrPermissionAccessLevel))
		return false
	}

	return true
}
