package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/library/dto"
	validator "pirate-lang-go/modules/library/validation"
)

func (controller *LibraryController) CreateExam(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	requestData := new(dto.CreateExamRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err.Error())
	}

	resultValidator := validator.ValidateCreateExam(requestData)
	if resultValidator != nil {
		return controller.BadRequest("Validation failed", resultValidator.Errors)
	}

	appErr := controller.libraryService.CreateExam(ctx, requestData)
	if appErr != nil {
		return controller.BadRequest("Error create exams", appErr.Error())
	}
	return controller.SuccessResponse(c, nil, "Create Exam successfully")
}

func (controller *LibraryController) UpdateExam(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse exam ID from path
	examIdStr := c.Param("examId")
	examId, err := uuid.Parse(examIdStr)
	if err != nil {
		return controller.BadRequest("Invalid exam ID format", err.Error())
	}

	// Parse request body
	requestData := new(dto.UpdateExamRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err.Error())
	}

	resultValidator := validator.ValidateUpdateExam(requestData)
	if resultValidator != nil {
		return controller.BadRequest("Validation failed", resultValidator.Errors)
	}

	appErr := controller.libraryService.UpdateExam(ctx, requestData, examId)
	if appErr != nil {
		return controller.BadRequest("Error update exams", appErr.Error())
	}
	return controller.SuccessResponse(c, nil, "Update Exam successfully")
}

func (controller *LibraryController) GetExam(c echo.Context) error {
	ctx := c.Request().Context()
	examIdStr := c.Param("examId")
	examId, err := uuid.Parse(examIdStr)
	if err != nil {
		return controller.BadRequest("Invalid exam ID format", err.Error())
	}

	response, appErr := controller.libraryService.GetExam(ctx, examId)
	if appErr != nil {
		return controller.BadRequest("Error getting exams", appErr.Error())
	}
	return controller.SuccessResponse(c, response, "Get Exam successfully")
}

func (controller *LibraryController) GetExams(c echo.Context) error {
	ctx := c.Request().Context()
	pageNumber := utils.ToNumberWithDefault(c.QueryParam("pageNumber"), 1)
	pageSize := utils.ToNumberWithDefault(c.QueryParam("pageSize"), 20)

	response, appErr := controller.libraryService.GetExams(ctx, pageNumber, pageSize)
	if appErr != nil {
		return controller.BadRequest("Error getting exams", appErr.Error())
	}
	return controller.SuccessResponse(c, response, "Get Exams successfully")
}
