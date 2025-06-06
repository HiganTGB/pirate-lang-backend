package controller

import (
	"pirate-lang-go/core/controller"
	"pirate-lang-go/modules/library/service"
)

type LibraryController struct {
	controller.BaseController
	libraryService service.ILibraryService
}

func NewLibraryController(service service.ILibraryService) *LibraryController {

	return &LibraryController{
		BaseController: controller.NewBaseController(),
		libraryService: service,
	}
}
