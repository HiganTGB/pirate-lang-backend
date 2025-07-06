package service

import (
	"context"
	"github.com/google/uuid"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/library/dto"
	"pirate-lang-go/modules/library/mapper"
	"time"
)

func (s *LibraryService) CreateExamPart(ctx context.Context, dataRequest *dto.CreateExamPartRequest) *errors.AppError {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	examPartEntity := mapper.ToCreateExamPartEntity(dataRequest)
	err := s.repo.CreateExamPart(ctx, examPartEntity)
	if err != nil {
		logger.Error("LibraryService:CreateExamPart:Failed to create exam part", "error", err)
		return errors.NewAppError(errors.ErrInternal, "LibraryService:CreateExamPart:Failed to create exam part", err)
	}
	return nil
}
func (s *LibraryService) UpdateExamPart(ctx context.Context, dataRequest *dto.UpdateExamPartRequest, examPartId uuid.UUID) *errors.AppError {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	examPartEntity := mapper.ToUpdateExamPartEntity(dataRequest)
	err := s.repo.UpdateExamPart(ctx, examPartEntity, examPartId)
	if err != nil {
		logger.Error("LibraryService:UpdateExamPart:Failed to update exam part", "exam_part_id", examPartId, "error", err)
		return errors.NewAppError(errors.ErrInternal, "LibraryService:UpdateExamPart:Failed to update exam part", err)
	}
	return nil
}
func (s *LibraryService) GetExamPart(ctx context.Context, examPartId uuid.UUID) (*dto.ExamPartResponse, *errors.AppError) {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	examPart, err := s.repo.GetExamPart(ctx, examPartId)
	if err != nil {
		logger.Error("LibraryService:GetExamPart:Failed to retrieve exam part", "exam_part_id", examPartId, "error", err)

		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetExamPart:Failed to retrieve exam part", err)
	}
	examPartDTO := mapper.ToExamPartResponse(examPart)
	return examPartDTO, nil
}
func (s *LibraryService) GetPracticeExamParts(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedExamPartResponse, *errors.AppError) {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resultGetExamParts, err := s.repo.GetPracticeExamParts(ctx, pageNumber, pageSize)
	if err != nil {
		logger.Error("LibraryService:GetExamParts:Failed to get exam parts", "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetExamParts:Failed to get exam parts", err)
	}

	examPartDTOs := mapper.ToPaginatedExamPartsResponse(resultGetExamParts)
	return examPartDTOs, nil
}
func (s *LibraryService) GetExamPartsByExamId(ctx context.Context, examId uuid.UUID) ([]*dto.ExamPartResponse, *errors.AppError) {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	examParts, err := s.repo.GetExamPartsByExamId(ctx, examId)
	if err != nil {
		logger.Error("LibraryService:GetExamPartsByExamId:Failed to retrieve exam parts by exam ID", "exam_id", examId, "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetExamPartsByExamId:Failed to retrieve exam parts by exam ID", err)
	}

	var examPartDTOs []*dto.ExamPartResponse
	for _, examPart := range examParts {
		examPartDTOs = append(examPartDTOs, mapper.ToExamPartResponse(examPart))
	}
	return examPartDTOs, nil
}
