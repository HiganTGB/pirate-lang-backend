package router

import (
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/middleware"
	"pirate-lang-go/modules/account/controller"
)

type AccountRouter struct {
	controller *controller.AccountController
}

func NewAccountRouter(controller *controller.AccountController) *AccountRouter {
	return &AccountRouter{
		controller: controller,
	}
}
func (r *AccountRouter) Setup(e *echo.Echo, middleware *middleware.Middleware) {
	// API v1 group
	v1 := e.Group("/v1")
	// Account routes group
	accounts := v1.Group("/accounts")
	// Auth routes - no middleware needed
	auth := accounts.Group("")
	auth.POST("/register", r.controller.Register)
	auth.POST("/login", r.controller.Login)
	auth.POST("/refresh-token", r.controller.RefreshToken)
	// User routes - requires authentication
	user := accounts.Group("")
	user.Use(middleware.AuthMiddleware())
	user.POST("/logout", r.controller.Logout)
	user.PUT("/change-password", r.controller.ChangePassword)
	user.GET("/profile", r.controller.GetProfile)
	user.POST("/profile", r.controller.CreateProfiles)
	user.PUT("/profile", r.controller.UpdateProfiles)
	user.POST("/profile/avatar", r.controller.UpdateAvatar)
	// Admin routes
	admin := v1.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	// User management routes
	users := admin.Group("/users")
	users.GET("", r.controller.GetUsers)
	users.GET("/:userId/profile", r.controller.GetDetailUser)
	users.POST("/:userId/lock", r.controller.LockUser)
	users.POST("/:userId/unlock", r.controller.UnlockUser)

	test := v1.Group("/test")
	test.GET("/hello", r.controller.HelloWorld)

	// RBAC management routes
	rbac := admin.Group("/rbac")
	rbac.GET("/roles", r.controller.GetRoles)
	rbac.POST("/roles", r.controller.CreateRole)
	rbac.GET("/permissions", r.controller.GetPermissions)
	rbac.POST("/permissions", r.controller.CreatePermission)
	rbac.POST("/roles/:roleId/permissions/:permissionId", r.controller.AssignPermissionToRole)
	rbac.POST("/roles/:roleId/users/:userId", r.controller.AssignRoleToUser)
}
