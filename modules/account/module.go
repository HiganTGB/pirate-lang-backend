package account

import (
	"github.com/labstack/echo/v4"
	"prirate-lang-go/core/database"
	"prirate-lang-go/core/middleware"
	"prirate-lang-go/modules/account/controller"
	"prirate-lang-go/modules/account/router"
)

func Init(e *echo.Echo, db database.Database) {
	middleware := middleware.NewMiddleware()
	// Update: pass only the controller
	router.NewAccountRouter(
		controller.NewAccountController(),
	).Setup(e, middleware) // Pass middleware to Setup instead
}
