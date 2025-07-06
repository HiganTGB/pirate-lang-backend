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

func (s *LibraryService) CreateParagraph(ctx context.Context, dataRequest *dto.CreateParagraphRequest) *errors.AppError {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	paragraphEntity := mapper.ToCreateParagraphEntity(dataRequest)
	err := s.repo.CreateParagraph(ctx, paragraphEntity)
	if err != nil {
		logger.Error("LibraryService:CreateParagraph:Failed to create paragraph", "error", err)
		return errors.NewAppError(errors.ErrInternal, "LibraryService:CreateParagraph:Failed to create paragraph", err)
	}
	return nil
}

func (s *LibraryService) UpdateParagraph(ctx context.Context, dataRequest *dto.UpdateParagraphRequest, paragraphId uuid.UUID) *errors.AppError {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	paragraphEntity := mapper.ToUpdateParagraphEntity(dataRequest)
	err := s.repo.UpdateParagraph(ctx, paragraphEntity, paragraphId)
	if err != nil {
		logger.Error("LibraryService:UpdateParagraph:Failed to update paragraph", "paragraph_id", paragraphId, "error", err)
		return errors.NewAppError(errors.ErrInternal, "LibraryService:UpdateParagraph:Failed to update paragraph", err)
	}
	return nil
}

func (s *LibraryService) GetParagraph(ctx context.Context, paragraphId uuid.UUID) (*dto.ParagraphResponse, *errors.AppError) {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	paragraph, err := s.repo.GetParagraph(ctx, paragraphId)
	if err != nil {
		logger.Error("LibraryService:GetParagraph:Failed to retrieve paragraph", "paragraph_id", paragraphId, "error", err)
		// Consider more specific error handling here, e.g., errors.ErrNotFound if the repository returns a specific "not found" error.
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetParagraph:Failed to retrieve paragraph", err)
	}
	paragraphDTO := mapper.ToParagraphResponse(paragraph)
	return paragraphDTO, nil
}

func (s *LibraryService) GetParagraphsByPartId(ctx context.Context, partId uuid.UUID) ([]*dto.ParagraphResponse, *errors.AppError) {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	paragraphs, err := s.repo.GetParagraphsByPartId(ctx, partId)
	if err != nil {
		logger.Error("LibraryService:GetParagraphsByPartId:Failed to retrieve paragraphs by part ID", "part_id", partId, "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetParagraphsByPartId:Failed to retrieve paragraphs by part ID", err)
	}

	var paragraphDTOs []*dto.ParagraphResponse
	for _, paragraph := range paragraphs {
		paragraphDTOs = append(paragraphDTOs, mapper.ToParagraphResponse(paragraph))
	}
	return paragraphDTOs, nil
}
func (s *LibraryService) UploadAudioParagraph(ctx context.Context, file *multipart.FileHeader, paragraphId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError) {
	src, err := file.Open()
	if err != nil {
		logger.Error("LibraryService:UploadAudioParagraph:Failed to open uploaded audio file", "error", err, "paragraphId", paragraphId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
	}
	defer src.Close()
	objectName, objectURL, err := s.storage.UploadAudio(ctx, paragraphId, src, file.Size, file.Filename, GroupFolder)
	if err != nil {
		logger.Error("LibraryService:UploadAudioParagraph:Failed to open uploaded audio file", "error", err, "paragraphId", paragraphId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
	}
	err = s.repo.UpdateAudioParagraph(ctx, &objectURL, paragraphId)
	if err != nil {
		logger.Error("LibraryService:UploadAudioParagraph:Failed to update audio URL in database", "error", err, "paragraphId", paragraphId.String())
		return nil, errors.NewAppError(errors.ErrInternal, "Service:UploadAudioGroup:Failed to persist audio information in database", err)
	}
	response := &dto.UpdateContentFileResponse{
		Filename:  objectName,
		ObjectURL: objectURL,
	}
	return response, nil
}
func (s *LibraryService) UploadTranscriptAudioParagraph(ctx context.Context, file *multipart.FileHeader, paragraphId uuid.UUID, language string) (*dto.UpdateContentFileResponse, *errors.AppError) {
	src, err := file.Open()
	if err != nil {
		logger.Error("LibraryService:UploadAudioParagraph:Failed to open uploaded audio file", "error", err, "paragraphId", paragraphId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
	}
	defer src.Close()
	objectName, objectURL, err := s.storage.UploadTranscriptAudio(ctx, paragraphId, src, file.Size, file.Filename, TranscriptFolder, language)
	if err != nil {
		logger.Error("LibraryService:UploadAudioParagraph:Failed to open uploaded audio file", "error", err, "paragraphId", paragraphId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
	}
	response := &dto.UpdateContentFileResponse{
		Filename:  objectName,
		ObjectURL: objectURL,
	}
	return response, nil
}
func (s *LibraryService) UploadImageParagraph(ctx context.Context, file *multipart.FileHeader, paragraphId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError) {
	src, err := file.Open()
	if err != nil {
		logger.Error("LibraryService:UploadAudioParagraph:Failed to open uploaded audio file", "error", err, "paragraphId", paragraphId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to read audio file", err)
	}
	defer src.Close()
	objectName, objectURL, err := s.storage.UploadImage(ctx, paragraphId, src, file.Size, file.Filename, ImageGroupFolder)
	if err != nil {
		logger.Error("LibraryService:UploadAudioParagraph:Failed to open uploaded audio file", "error", err, "paragraphId", paragraphId.String())
		return nil, errors.NewAppError(errors.ErrInvalidInput, "Service:UploadAudioGroup:Failed to upload audio file", err)
	}
	err = s.repo.UpdateAudioParagraph(ctx, &objectName, paragraphId)
	if err != nil {
		logger.Error("LibraryService:UploadAudioParagraph:Failed to update audio URL in database", "error", err, "paragraphId", paragraphId.String())
		return nil, errors.NewAppError(errors.ErrInternal, "Service:UploadAudioGroup:Failed to persist audio information in database", err)
	}
	response := &dto.UpdateContentFileResponse{
		Filename:  objectName,
		ObjectURL: objectURL,
	}
	return response, nil
}
func (s *LibraryService) DeleteAudioParagraph(ctx context.Context, groupId uuid.UUID) *errors.AppError {
	objectName := ""
	err := s.repo.UpdateAudioParagraph(ctx, &objectName, groupId)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
		return errors.NewAppError(errors.ErrInternal, "LibraryService:UploadAudioGroup:Failed to persist audio information in database", err)
	}
	return nil
}
