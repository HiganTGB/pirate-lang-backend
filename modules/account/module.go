package account

import (
	"github.com/gin-gonic/gin"
	"prirate-lang-go/core/database"
	"prirate-lang-go/core/middleware"
	"prirate-lang-go/modules/account/controller"
	"prirate-lang-go/modules/account/router"
)

func Init(r *gin.Engine, db database.Database) {

	appMiddleware := middleware.NewMiddleware()

	// Update: pass only the controller
	router.NewAccountRouter(
		controller.NewAccountController(),
	).Setup(r, appMiddleware) // Pass middleware to Setup instead
}
