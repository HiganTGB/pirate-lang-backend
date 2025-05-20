package controller

import (
	"github.com/gin-gonic/gin"
)

func (controller *AccountController) HelloWorld(c *gin.Context) {
	controller.SuccessResponse(c, gin.H{
		"message": "Hello, World!",
	}, "Chào mừng!")
}
