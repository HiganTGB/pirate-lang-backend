package controller

import (
	"pirate-lang-go/core/controller"
	"pirate-lang-go/modules/account/service"
)

type AccountController struct {
	controller.BaseController
	accountService service.IAccountService
}

func NewAccountController(service service.IAccountService) *AccountController {

	return &AccountController{
		BaseController: controller.NewBaseController(),
		accountService: service,
	}
}
