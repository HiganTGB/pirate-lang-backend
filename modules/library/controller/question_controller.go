package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/library/dto"
	validator "pirate-lang-go/modules/library/validation"
)

func (controller *LibraryController) UploadAudioGroup(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	idStr := c.Param("questionId")
	questionId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid questionId ID format", errParse)
	}
	file, errFile := c.FormFile("audio")
	if errFile != nil {
		return controller.BadRequest(fmt.Sprintf("Error getting audio file: %v", errFile))
	}
	contentType := file.Header.Get("Content-Type")
	if !utils.IsAudioMpegContentType(contentType) {
		return controller.BadRequest("Invalid file type. Only MP3 files are allowed.")
	}

	resultUpdateAudio, errUpload := controller.libraryService.UploadAudioQuestion(ctx, file, questionId)
	if errUpload != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating audio: %v", errUpload))
	}
	return controller.SuccessResponse(c, resultUpdateAudio, "Update Audio successfully")
}
func (controller *LibraryController) UploadTranscriptAudioGroup(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	idStr := c.Param("questionId")
	questionId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid questionId ID format", errParse)
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

	fileResponse, err := controller.libraryService.UploadTranscriptAudioParagraph(ctx, file, questionId, lang)
	if err != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating avatar: %v", err))
	}
	return controller.SuccessResponse(c, fileResponse, "Update Transcript Audio successfully")
}
func (controller *LibraryController) UploadImageGroup(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("questionId")
	questionId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid questionId ID format", errParse)
	}
	file, errFile := c.FormFile("image")
	if errFile != nil {
		return controller.BadRequest(fmt.Sprintf("Error getting image file: %v", errFile))
	}
	contentType := file.Header.Get("Content-Type")
	if !utils.IsImageContentType(contentType) {
		return controller.BadRequest("Invalid file type. Only image files (JPEG, PNG) are allowed.")
	}
	resultUpdateAvatar, err := controller.libraryService.UploadImageQuestion(ctx, file, questionId)
	if err != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating image: %v", err))
	}
	return controller.SuccessResponse(c, resultUpdateAvatar, "Update image question successfully")
}

func (controller *LibraryController) GetQuestionsParagraph(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("paragraphId")
	paragraphId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid paragraph ID format", errParse)
	}
	response, err := controller.libraryService.GetQuestionsByParagraph(ctx, paragraphId)
	if err != nil {
		return controller.BadRequest("Error get group", err)
	}
	return controller.SuccessResponse(c, response, "Get Question successfully")
}
func (controller *LibraryController) GetQuestionsPart(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("partId")
	paragraphId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid paragraph ID format", errParse)
	}
	pageNumber := utils.ToNumberWithDefault(c.QueryParam("pageNumber"), 1)
	pageSize := utils.ToNumberWithDefault(c.QueryParam("pageSize"), 20)
	response, err := controller.libraryService.GetQuestionByParts(ctx, pageNumber, pageSize, paragraphId)
	if err != nil {
		return controller.BadRequest("Error get group", err)
	}
	return controller.SuccessResponse(c, response, "Get Question successfully")
}
func (controller *LibraryController) CreateQuestion(c echo.Context) error {
	ctx := c.Request().Context()
	requestData := new(dto.CreateQuestionRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidator := validator.ValidateCreateQuestion(requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	question, err := controller.libraryService.CreateQuestion(ctx, requestData)
	if err != nil {
		return controller.BadRequest("Error create question", err)
	}
	return controller.SuccessResponse(c, question, "Create question successfully")
}
func (controller *LibraryController) UpdateQuestion(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("questionId")
	questionId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid questionId ID format", errParse)
	}
	requestData := new(dto.UpdateQuestionRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidator := validator.ValidateUpdateQuestion(requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	err := controller.libraryService.UpdateQuestion(ctx, requestData, questionId)
	if err != nil {
		return controller.BadRequest("Error create question", err)
	}
	return controller.SuccessResponse(c, nil, "Create question successfully")
}
