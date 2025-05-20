package middleware

import (
	"github.com/gin-gonic/gin"
	"prirate-lang-go/core/controller"
	"prirate-lang-go/core/logger"
	"time"
)

type Middleware struct {
	controller.BaseController
}

func NewMiddleware() *Middleware {
	return &Middleware{
		BaseController: controller.NewBaseController(),
	}
}
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		startTime := time.Now()
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		remoteIP := c.ClientIP()

		c.Next()

		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		logger.Info("Request Handled",
			"method", reqMethod,
			"uri", reqURI,
			"status", statusCode,
			"remote_ip", remoteIP,
			"duration", duration,
		)

	}
}
