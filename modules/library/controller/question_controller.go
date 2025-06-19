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
	groupIdStr := c.Param("groupId")
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

	resultUpdateAudio, errUpload := controller.libraryService.UploadAudioGroup(ctx, file, groupId)
	if errUpload != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating audio: %v", err))
	}
	return controller.SuccessResponse(c, resultUpdateAudio, "Update Audio successfully")
}
func (controller *LibraryController) UploadTranscriptAudioGroup(c echo.Context) error {
	ctx := c.Request().Context()
	// Parse request body
	groupIdStr := c.Param("groupId")
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

	fileResponse, err := controller.libraryService.UploadTranscriptAudioGroup(ctx, file, groupId, lang)
	if err != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating avatar: %v", err))
	}
	return controller.SuccessResponse(c, fileResponse, "Update Transcript Audio successfully")
}
func (controller *LibraryController) UploadImageGroup(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("groupId")
	groupId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid group ID format", errParse)
	}
	file, errFile := c.FormFile("image")
	if errFile != nil {
		return controller.BadRequest(fmt.Sprintf("Error getting image file: %v", errFile))
	}
	contentType := file.Header.Get("Content-Type")
	if !utils.IsImageContentType(contentType) {
		return controller.BadRequest("Invalid file type. Only image files (JPEG, PNG) are allowed.")
	}
	resultUpdateAvatar, err := controller.libraryService.UploadImageGroup(ctx, file, groupId)
	if err != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating image: %v", err))
	}
	return controller.SuccessResponse(c, resultUpdateAvatar, "Update image question successfully")
}
func (controller *LibraryController) CreateQuestionGroup(c echo.Context) error {
	ctx := c.Request().Context()
	requestData := new(dto.CreateQuestionGroupRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidator := validator.ValidateCreateQuestionGroup(requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	groupId, err := controller.libraryService.CreateQuestionGroup(ctx, requestData)
	if err != nil {
		return controller.BadRequest("Error create group", err)
	}
	return controller.SuccessResponse(c, groupId, "Create group successfully")
}
func (controller *LibraryController) UpdateQuestionGroup(c echo.Context) error {
	ctx := c.Request().Context()
	idStr := c.Param("groupId")
	groupId, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return controller.BadRequest("Invalid group ID format", errParse)
	}
	requestData := new(dto.UpdateQuestionGroupRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidator := validator.ValidateUpdateQuestionGroup(requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	err := controller.libraryService.UpdateQuestionGroup(ctx, groupId, requestData)
	if err != nil {
		return controller.BadRequest("Error create group", err)
	}
	return controller.SuccessResponse(c, nil, "Update group successfully")
}

func (controller *LibraryController) GetQuestionGroups(c echo.Context) error {
	ctx := c.Request().Context()
	pageNumber := utils.ToNumberWithDefault(c.QueryParam("pageNumber"), 1)
	pageSize := utils.ToNumberWithDefault(c.QueryParam("pageSize"), 20)
	response, err := controller.libraryService.GetQuestionGroups(ctx, pageNumber, pageSize)
	if err != nil {
		return controller.BadRequest("Error get group", err)
	}
	return controller.SuccessResponse(c, response, "Get Groups successfully")
}
