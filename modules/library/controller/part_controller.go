package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/library/dto"
	validator "pirate-lang-go/modules/library/validation"
)

func (controller *LibraryController) CreatePart(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	requestData := new(dto.CreatePartRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidator := validator.ValidateCreatePart(requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	err := controller.libraryService.CreatePart(ctx, requestData)
	if err != nil {
		return controller.BadRequest("Error create part", err)
	}
	return controller.SuccessResponse(c, nil, "Create Part successfully")
}
func (controller *LibraryController) UpdatePart(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	partIdStr := c.Param("partId")
	partId, err := uuid.Parse(partIdStr)
	if err != nil {
		return controller.BadRequest("Invalid part ID format", err)
	}
	requestData := new(dto.UpdatePartRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidator := validator.ValidateUpdatePart(requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	err = controller.libraryService.UpdatePart(ctx, requestData, partId)
	if err != nil {
		return controller.BadRequest("Error get part", err)
	}
	return controller.SuccessResponse(c, nil, "Update Part successfully")
}
func (controller *LibraryController) GetPart(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	partIdStr := c.Param("partId")
	partId, err := uuid.Parse(partIdStr)
	if err != nil {
		return controller.BadRequest("Invalid part ID format", err)
	}
	response, err := controller.libraryService.GetPart(ctx, partId)
	if err != nil {
		return controller.BadRequest("Error get part", err)
	}
	return controller.SuccessResponse(c, response, "Get Part successfully")
}
func (controller *LibraryController) GetParts(c echo.Context) error {
	ctx := c.Request().Context()
	pageNumber := utils.ToNumberWithDefault(c.QueryParam("pageNumber"), 1)
	pageSize := utils.ToNumberWithDefault(c.QueryParam("pageSize"), 20)
	response, err := controller.libraryService.GetParts(ctx, pageNumber, pageSize)
	if err != nil {
		return controller.BadRequest("Error get part", err)
	}
	return controller.SuccessResponse(c, response, "Get Parts successfully")
}
