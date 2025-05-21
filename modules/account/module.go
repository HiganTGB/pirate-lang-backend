package account

import (
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/database"
	"pirate-lang-go/core/middleware"
	"pirate-lang-go/modules/account/controller"
	"pirate-lang-go/modules/account/repository"
	"pirate-lang-go/modules/account/router"
	"pirate-lang-go/modules/account/service"
)

func Init(e *echo.Echo, db database.Database) {
	repository := repository.NewAccountRepository(db.DB())
	accountService := service.NewAccountService(repository)
	middleware := middleware.NewMiddleware(accountService)
	// Update: pass only the controller
	router.NewAccountRouter(
		controller.NewAccountController(accountService),
	).Setup(e, middleware) // Pass middleware to Setup instead
}
