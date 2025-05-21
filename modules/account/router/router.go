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
	user.POST("/logout", r.controller.Logout)
	user.PUT("/change-password", r.controller.ChangePassword)
	// Admin routes
	admin := v1.Group("/admin")

	// User management routes
	users := admin.Group("/users")
	users.GET("", r.controller.GetUsers)

	test := v1.Group("/test")
	test.GET("/hello", r.controller.HelloWorld)
}
