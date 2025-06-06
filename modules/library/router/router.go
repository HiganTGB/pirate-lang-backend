package router

import (
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/middleware"
	"pirate-lang-go/modules/library/controller"
)

type LibraryRouter struct {
	controller *controller.LibraryController
}

func NewLibraryRouter(controller *controller.LibraryController) *LibraryRouter {
	return &LibraryRouter{
		controller: controller,
	}
}
func (r *LibraryRouter) Setup(e *echo.Echo, middleware *middleware.Middleware) {
	// API v1 group
	v1 := e.Group("/v1")

	test := v1.Group("/test2")
	test.GET("/hello", r.controller.HelloWorld)

}
