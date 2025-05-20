package router

import (
	"github.com/gin-gonic/gin"
	"prirate-lang-go/core/middleware"
	"prirate-lang-go/modules/account/controller"
)

type AccountRouter struct {
	controller *controller.AccountController
}

func NewAccountRouter(controller *controller.AccountController) *AccountRouter {
	return &AccountRouter{
		controller: controller,
	}
}

func (r *AccountRouter) Setup(engine *gin.Engine, appMiddleware *middleware.Middleware) {
	v1 := engine.Group("/v1")
	test := v1.Group("")
	test.GET("/hello", r.controller.HelloWorld)
}
