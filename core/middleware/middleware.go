package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pirate-lang-go/core/controller"
	"pirate-lang-go/core/logger"
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
func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(http.StatusOK)
			}

			return next(c)
		}
	}
}
