package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/library/dto"
	validator "pirate-lang-go/modules/library/validation"
)

func (controller *LibraryController) CreateExamPart(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	requestData := new(dto.CreateExamPartRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err.Error())
	}

	resultValidator := validator.ValidateCreateExamPart(requestData)
	if resultValidator != nil {
		return controller.BadRequest("Validation failed", resultValidator.Errors)
	}

	appErr := controller.libraryService.CreateExamPart(ctx, requestData)
	if appErr != nil {
		return controller.BadRequest("Error create exams", appErr.Error())
	}
	return controller.SuccessResponse(c, nil, "Create Exam successfully")
}

func (controller *LibraryController) UpdateExamPart(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse exam ID from path
	examIdStr := c.Param("partId")
	examId, err := uuid.Parse(examIdStr)
	if err != nil {
		return controller.BadRequest("Invalid exam ID format", err.Error())
	}

	// Parse request body
	requestData := new(dto.UpdateExamPartRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err.Error())
	}

	resultValidator := validator.ValidateUpdateExamPart(requestData)
	if resultValidator != nil {
		return controller.BadRequest("Validation failed", resultValidator.Errors)
	}

	appErr := controller.libraryService.UpdateExamPart(ctx, requestData, examId)
	if appErr != nil {
		return controller.BadRequest("Error update exams", appErr.Error())
	}
	return controller.SuccessResponse(c, nil, "Update Exam successfully")
}

func (controller *LibraryController) GetExamPart(c echo.Context) error {
	ctx := c.Request().Context()
	examIdStr := c.Param("partId")
	examId, err := uuid.Parse(examIdStr)
	if err != nil {
		return controller.BadRequest("Invalid exam ID format", err.Error())
	}

	response, appErr := controller.libraryService.GetExamPart(ctx, examId)
	if appErr != nil {
		return controller.BadRequest("Error getting exams", appErr.Error())
	}
	return controller.SuccessResponse(c, response, "Get Exam successfully")
}

func (controller *LibraryController) GetPracticeParts(c echo.Context) error {
	ctx := c.Request().Context()
	pageNumber := utils.ToNumberWithDefault(c.QueryParam("pageNumber"), 1)
	pageSize := utils.ToNumberWithDefault(c.QueryParam("pageSize"), 20)

	response, appErr := controller.libraryService.GetPracticeExamParts(ctx, pageNumber, pageSize)
	if appErr != nil {
		return controller.BadRequest("Error getting exams", appErr.Error())
	}
	return controller.SuccessResponse(c, response, "Get Exams successfully")
}
func (controller *LibraryController) GetExamPartsByExam(c echo.Context) error {
	ctx := c.Request().Context()
	examIdStr := c.Param("examId")
	examId, err := uuid.Parse(examIdStr)
	if err != nil {
		return controller.BadRequest("Invalid exam ID format", err.Error())
	}
	response, appErr := controller.libraryService.GetExamPartsByExamId(ctx, examId)
	if appErr != nil {
		return controller.BadRequest("Error getting exams", appErr.Error())
	}
	return controller.SuccessResponse(c, response, "Get Exams successfully")
}
