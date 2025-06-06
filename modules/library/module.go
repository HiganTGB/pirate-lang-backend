package library

import (
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/cache"
	"pirate-lang-go/core/database"
	"pirate-lang-go/core/middleware"
	"pirate-lang-go/core/storage"
	accountrepo "pirate-lang-go/modules/account/repository"
	accountservice "pirate-lang-go/modules/account/service"
	"pirate-lang-go/modules/library/controller"
	"pirate-lang-go/modules/library/repository"
	"pirate-lang-go/modules/library/router"
	"pirate-lang-go/modules/library/service"
)

func Init(e *echo.Echo, db database.Database, cache *cache.Cache, storage *storage.Storage) {
	accountService := accountservice.NewAccountService(accountrepo.NewAccountRepository(db.DB()), cache, storage)
	middleware := middleware.NewMiddleware(accountService)
	repository := repository.NewLibraryRepository(db.DB())

	libraryService := service.NewLibraryService(repository, cache, storage)
	// Update: pass only the controller
	router.NewLibraryRouter(
		controller.NewLibraryController(libraryService),
	).Setup(e, middleware)
}
