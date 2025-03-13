package routes

import (
	"icdn/controllers"
	"icdn/middleware"
	"icdn/types"

	"github.com/gin-gonic/gin"
)

func PermissionRoutes(router *gin.RouterGroup) {
	permissionRoutes := router.Group("/permissions")
	permissionRoutes.Use(middleware.AuthMiddleware())
	{
		permissionRoutes.POST(
			"/save_permission",
			middleware.ValidateRequest(types.SavePermissionTypes{}),
			middleware.PermissionValidator[types.SavePermissionTypes](),
			controllers.CreatePermission,
		)

		permissionRoutes.POST(
			"/get_permission_id",
			middleware.ValidateRequest(types.PermissionID{}),
			controllers.GetPermissionByID,
		)

		permissionRoutes.POST(
			"/get_all_permissions",
			middleware.ValidateRequest(types.PaginationRequest{}),
			middleware.Pagination(),
			controllers.GetAllPermissions,
		)

		permissionRoutes.PUT(
			"/update_permission",
			middleware.ValidateRequest(types.UpdatePermissionTypes{}),
			middleware.PermissionValidator[types.UpdatePermissionTypes](),
			controllers.UpdatePermission,
		)

		permissionRoutes.POST(
			"/delete_permission",
			middleware.ValidateRequest(types.PermissionID{}),
			controllers.DeletePermission,
		)
	}
}
