package router

import (
	"github.com/labstack/echo/v4"
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
func (r *AccountRouter) Setup(e *echo.Echo, middleware *middleware.Middleware) {
	// API v1 group
	v1 := e.Group("/v1")
	test := v1.Group("")
	test.GET("/hello", r.controller.HelloWorld)
}
