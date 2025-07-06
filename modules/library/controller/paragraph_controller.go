package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/library/dto"
	validator "pirate-lang-go/modules/library/validation"
)

func (controller *LibraryController) CreateParagraph(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	requestData := new(dto.CreateParagraphRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err.Error())
	}

	resultValidator := validator.ValidateCreateParagraph(requestData)
	if resultValidator != nil {
		return controller.BadRequest("Validation failed", resultValidator.Errors)
	}

	appErr := controller.libraryService.CreateParagraph(ctx, requestData)
	if appErr != nil {
		return controller.BadRequest("Error create exams", appErr.Error())
	}
	return controller.SuccessResponse(c, nil, "Create Exam successfully")
}

func (controller *LibraryController) UpdateParagraph(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse exam ID from path
	examIdStr := c.Param("paragraphId")
	examId, err := uuid.Parse(examIdStr)
	if err != nil {
		return controller.BadRequest("Invalid exam ID format", err.Error())
	}

	// Parse request body
	requestData := new(dto.UpdateParagraphRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err.Error())
	}

	resultValidator := validator.ValidateUpdateParagraph(requestData)
	if resultValidator != nil {
		return controller.BadRequest("Validation failed", resultValidator.Errors)
	}

	appErr := controller.libraryService.UpdateParagraph(ctx, requestData, examId)
	if appErr != nil {
		return controller.BadRequest("Error update paragraphs", appErr.Error())
	}
	return controller.SuccessResponse(c, nil, "Update Exam successfully")
}

func (controller *LibraryController) GetParagraph(c echo.Context) error {
	ctx := c.Request().Context()
	IdStr := c.Param("paragraphId")
	id, err := uuid.Parse(IdStr)
	if err != nil {
		return controller.BadRequest("Invalid exam ID format", err.Error())
	}

	response, appErr := controller.libraryService.GetExamPart(ctx, id)
	if appErr != nil {
		return controller.BadRequest("Error getting exams", appErr.Error())
	}
	return controller.SuccessResponse(c, response, "Get Exam successfully")
}

func (controller *LibraryController) GetParagraphsByPart(c echo.Context) error {
	ctx := c.Request().Context()
	IdStr := c.Param("partId")
	id, err := uuid.Parse(IdStr)
	if err != nil {
		return controller.BadRequest("Invalid exam ID format", err.Error())
	}

	response, appErr := controller.libraryService.GetParagraphsByPartId(ctx, id)
	if appErr != nil {
		return controller.BadRequest("Error getting exams", appErr.Error())
	}
	return controller.SuccessResponse(c, response, "Get Exams successfully")
}
func (controller *LibraryController) UploadAudioParagraph(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	groupIdStr := c.Param("paragraphId")
	groupId, err := uuid.Parse(groupIdStr)
	if err != nil {
		return controller.BadRequest("Invalid group ID format", err)
	}
	file, errFile := c.FormFile("audio")
	if errFile != nil {
		return controller.BadRequest(fmt.Sprintf("Error getting audio file: %v", errFile))
	}
	contentType := file.Header.Get("Content-Type")
	if !utils.IsAudioMpegContentType(contentType) {
		return controller.BadRequest("Invalid file type. Only MP3 files are allowed.")
	}

	resultUpdateAudio, errUpload := controller.libraryService.UploadAudioParagraph(ctx, file, groupId)
	if errUpload != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating audio: %v", err))
	}
	return controller.SuccessResponse(c, resultUpdateAudio, "Update Audio successfully")
}
func (controller *LibraryController) UploadTranscriptAudioParagraph(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	groupIdStr := c.Param("paragraphId")
	groupId, errParse := uuid.Parse(groupIdStr)
	if errParse != nil {
		return controller.BadRequest("Invalid group ID format", errParse)
	}

	lang := c.QueryParam("lang")
	if validator.ValidateLang(lang) {
		return controller.BadRequest("Invalid lang type")
	}

	file, errFile := c.FormFile("transcript")
	if errFile != nil {
		return controller.BadRequest(fmt.Sprintf("Error getting audio file: %v", errFile))
	}
	contentType := file.Header.Get("Content-Type")
	if !utils.IsTextPlainContentType(contentType) {
		return controller.BadRequest("Invalid file type. Only text files are allowed.")
	}

	fileResponse, err := controller.libraryService.UploadTranscriptAudioParagraph(ctx, file, groupId, lang)
	if err != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating avatar: %v", err))
	}
	return controller.SuccessResponse(c, fileResponse, "Update Transcript Audio successfully")
}
func (controller *LibraryController) UploadImageParagraph(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("paragraphId")
	groupId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid paragraph ID format", errParse)
	}
	file, errFile := c.FormFile("image")
	if errFile != nil {
		return controller.BadRequest(fmt.Sprintf("Error getting image file: %v", errFile))
	}
	contentType := file.Header.Get("Content-Type")
	if !utils.IsImageContentType(contentType) {
		return controller.BadRequest("Invalid file type. Only image files (JPEG, PNG) are allowed.")
	}
	resultUpdateAvatar, err := controller.libraryService.UploadImageParagraph(ctx, file, groupId)
	if err != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating image: %v", err))
	}
	return controller.SuccessResponse(c, resultUpdateAvatar, "Update image question successfully")
}
