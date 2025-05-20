package controller

import (
	"prirate-lang-go/core/controller"
)

type AccountController struct {
	controller.BaseController
}

func NewAccountController() *AccountController {

	return &AccountController{
		BaseController: controller.NewBaseController(),
	}
}
