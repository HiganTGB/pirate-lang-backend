package service

import (
	"context"
	"github.com/google/uuid"
	"mime/multipart"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/library/dto"
	"pirate-lang-go/modules/library/mapper"
	"time"
)

const (
	GroupFolder      = "GroupFolder"
	QuestionFolder   = "QuestionFolder"
	TranscriptFolder = "TranscriptFolder"
	ImageGroupFolder = "TranscriptFolder"
)

func (s *LibraryService) CreateQuestionGroup(ctx context.Context, req *dto.CreateQuestionGroupRequest) (string, *errors.AppError) {
	groupEntity := mapper.ToQuestionGroupEntity(req)
	groupId, err := s.repo.CreateGroupGroup(ctx, groupEntity)
	if err != nil {
		logger.Error("LibraryService:CreateQuestionGroup:Failed to create question group", "error", err)
		return "", errors.NewAppError(errors.ErrInternal, "Service:CreateQuestionGroup:Failed to create question group", err)
	}
	return groupId.String(), nil
}
func (s *LibraryService) UpdateQuestionGroup(ctx context.Context, groupId uuid.UUID, req *dto.UpdateQuestionGroupRequest) *errors.AppError {

	groupEntity := mapper.ToQuestionGroupEntityForUpdate(req)
	err := s.repo.UpdateQuestionGroup(ctx, groupEntity, groupId)
	if err != nil {
		logger.Error("LibraryService:UpdateQuestionGroup:Failed to update question group", "error", err)
		return errors.NewAppError(errors.ErrInternal, "Service:UpdateQuestionGroup:Failed to update question group", err)
	}
	return nil
}
func (s *LibraryService) UploadAudioGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError) {
	src, err := file.Open()
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
	}
	defer src.Close()
	objectName, objectURL, err := s.storage.UploadAudio(ctx, groupId, src, file.Size, file.Filename, GroupFolder)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
	}
	err = s.repo.UpdateAudioGroup(ctx, &objectURL, groupId)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInternal, "Service:UploadAudioGroup:Failed to persist audio information in database", err)
	}
	response := &dto.UpdateContentFileResponse{
		Filename:  objectName,
		ObjectURL: objectURL,
	}
	return response, nil
}
func (s *LibraryService) UploadTranscriptAudioGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID, language string) (*dto.UpdateContentFileResponse, *errors.AppError) {
	src, err := file.Open()
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
	}
	defer src.Close()
	objectName, objectURL, err := s.storage.UploadTranscriptAudio(ctx, groupId, src, file.Size, file.Filename, TranscriptFolder, language)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
	}
	err = s.repo.UpdateAudioGroup(ctx, &objectName, groupId)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInternal, "Service:UploadAudioGroup:Failed to persist audio information in database", err)
	}
	response := &dto.UpdateContentFileResponse{
		Filename:  objectName,
		ObjectURL: objectURL,
	}
	return response, nil
}
func (s *LibraryService) UploadImageGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError) {
	src, err := file.Open()
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
	}
	defer src.Close()
	objectName, objectURL, err := s.storage.UploadImage(ctx, groupId, src, file.Size, file.Filename, ImageGroupFolder)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to open uploaded audio file", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
	}
	err = s.repo.UpdateAudioGroup(ctx, &objectName, groupId)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
		return nil, errors.NewAppError(errors.ErrInternal, "Service:UploadAudioGroup:Failed to persist audio information in database", err)
	}
	response := &dto.UpdateContentFileResponse{
		Filename:  objectName,
		ObjectURL: objectURL,
	}
	return response, nil
}
func (s *LibraryService) DeleteAudioGroup(ctx context.Context, groupId uuid.UUID) *errors.AppError {
	objectName := ""
	err := s.repo.UpdateAudioGroup(ctx, &objectName, groupId)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
		return errors.NewAppError(errors.ErrInternal, "LibraryService:UploadAudioGroup:Failed to persist audio information in database", err)
	}
	return nil
}
func (s *LibraryService) GetQuestionGroups(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedGroupResponse, *errors.AppError) {

	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	getQuestionGroups, err := s.repo.GetQuestionGroups(ctx, pageNumber, pageSize)
	if err != nil {
		logger.Error("LibraryService:GetParts:Failed to get parts", "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetQuestionGroups:Failed to Get Question group", err)
	}
	// Convert to DTO
	groupDTOs := mapper.ToPaginatedGroupsResponse(getQuestionGroups)
	return groupDTOs, nil
}
