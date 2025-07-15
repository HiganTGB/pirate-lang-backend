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

func (s *LibraryService) UploadAudioQuestion(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError) {
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
	err = s.repo.UpdateQuestionAudioUrl(ctx, &objectURL, groupId)
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
func (s *LibraryService) UploadTranscriptQuestion(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID, language string) (*dto.UpdateContentFileResponse, *errors.AppError) {
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
	response := &dto.UpdateContentFileResponse{
		Filename:  objectName,
		ObjectURL: objectURL,
	}
	return response, nil
}
func (s *LibraryService) UploadImageQuestion(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError) {
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
	err = s.repo.UpdateQuestionImageUrl(ctx, &objectName, groupId)
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
	err := s.repo.UpdateQuestionAudioUrl(ctx, &objectName, groupId)
	if err != nil {
		logger.Error("LibraryService:UploadAudioGroup:Failed to update audio URL in database", "error", err, "groupId", groupId.String())
		return errors.NewAppError(errors.ErrInternal, "LibraryService:UploadAudioGroup:Failed to persist audio information in database", err)
	}
	return nil
}
func (s *LibraryService) GetQuestionByParts(ctx context.Context, pageNumber, pageSize int, partId uuid.UUID) (*dto.PaginatedQuestionResponse, *errors.AppError) {

	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	getQuestionGroups, err := s.repo.GetSeparateQuestionsByPart(ctx, partId, pageNumber, pageSize)
	if err != nil {
		logger.Error("LibraryService:GetParts:Failed to get parts", "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetQuestionGroups:Failed to Get Question group", err)
	}
	groupDTOs := mapper.ToPaginatedQuestionResponse(getQuestionGroups)
	return groupDTOs, nil
}
func (s *LibraryService) GetQuestionsByParagraph(ctx context.Context, paragraphId uuid.UUID) ([]*dto.QuestionResponse, error) {
	questionDBs, err := s.repo.GetQuestionsByParagraph(ctx, paragraphId)
	if err != nil {
		logger.Error("LibraryService:GetQuestionGroup:Failed to get questions from group", err)
		return nil, err
	}
	var questions []*dto.QuestionResponse
	for _, questionDB := range questionDBs {
		question := mapper.ToQuestionResponse(questionDB)
		questions = append(questions, question)
	}
	return questions, nil
}
func (s *LibraryService) CreateQuestion(ctx context.Context, request *dto.CreateQuestionRequest) (*dto.QuestionResponse, error) {
	questionEntity := mapper.ToCreateQuestionEntity(request)
	question, err := s.repo.CreateQuestion(ctx, questionEntity)
	if err != nil {
		logger.Error("LibraryService:CreateQuestion: failed to create question", err)
		return nil, err
	}
	response := mapper.ToQuestionResponse(question)
	return response, nil
}
func (s *LibraryService) UpdateQuestion(ctx context.Context, request *dto.UpdateQuestionRequest, questionId uuid.UUID) error {
	questionEntity := mapper.ToUpdateQuestionEntity(request)
	err := s.repo.UpdateQuestion(ctx, questionEntity, questionId)
	if err != nil {
		logger.Error("LibraryService:CreateQuestion: failed to create question", err)
		return err
	}
	return nil
}
func (s *LibraryService) GetQuestion(ctx context.Context, questionId uuid.UUID) (*dto.QuestionResponse, error) {
	question, err := s.repo.GetQuestion(ctx, questionId)
	if err != nil {
		logger.Error("LibraryService:CreateQuestion: failed to create question", err)
		return nil, err
	}
	response := mapper.ToQuestionResponse(question)
	return response, nil
}
