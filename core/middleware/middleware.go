package middleware

import (
	"github.com/labstack/echo/v4"
	"prirate-lang-go/core/controller"
	"prirate-lang-go/core/logger"
)

type Middleware struct {
	controller.BaseController
}

func NewMiddleware() *Middleware {
	return &Middleware{
		BaseController: controller.NewBaseController(),
	}
}
func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			if err := next(c); err != nil {
				c.Error(err)
			}

			// Log request details
			logger.Info("Request",
				"method", req.Method,
				"uri", req.RequestURI,
				"status", res.Status,
				"remote_ip", c.RealIP(),
			)

			return nil
		}
	}
}
