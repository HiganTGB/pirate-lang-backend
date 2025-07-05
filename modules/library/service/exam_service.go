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

func (s *LibraryService) GetExams(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedExamResponse, *errors.AppError) {

	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resultGetExams, err := s.repo.GetExams(ctx, pageNumber, pageSize)
	if err != nil {
		logger.Error("LibraryService:GetExams:Failed to get exams", "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetExams:Failed to get exams", err)
	}

	examDTOs := mapper.ToPaginatedExamsResponse(resultGetExams)
	return examDTOs, nil
}

func (s *LibraryService) CreateExam(ctx context.Context, dataRequest *dto.CreateExamRequest) *errors.AppError {

	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := s.repo.CreateExam(ctx, mapper.ToCreateExamEntity(dataRequest))
	if err != nil {
		logger.Error("LibraryService:CreateExam:Failed to create exam", "error", err)
		return errors.NewAppError(errors.ErrInternal, "LibraryService:CreateExam:Failed to create exam", err)
	}
	return nil
}

func (s *LibraryService) UpdateExam(ctx context.Context, dataRequest *dto.UpdateExamRequest, examId uuid.UUID) *errors.AppError {

	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := s.repo.UpdateExam(ctx, mapper.ToUpdateExamEntity(dataRequest), examId)
	if err != nil {
		logger.Error("LibraryService:UpdateExam:Failed to update exam", "exam_id", examId, "error", err)
		return errors.NewAppError(errors.ErrInternal, "LibraryService:UpdateExam:Failed to update exam", err)
	}
	return nil
}

func (s *LibraryService) GetExam(ctx context.Context, examId uuid.UUID) (*dto.ExamResponse, *errors.AppError) {
	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	exam, err := s.repo.GetExam(ctx, examId)
	if err != nil {
		logger.Error("LibraryService:GetExam:Failed to retrieve exam", "exam_id", examId, "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetExam:Failed to retrieve exam", err)
	}
	examDTO := mapper.ToExamResponse(exam)
	return examDTO, nil
}
