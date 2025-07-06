package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/library/entity"
)

func (r *LibraryRepository) CreateExamPart(ctx context.Context, examPart *entity.ExamPart) error {
	var (
		ExamID              uuid.NullUUID
		PartTitle           string
		PartOrder           sql.NullInt32
		Description         sql.NullString
		IsPracticeComponent sql.NullBool
		PlanType            string
		ToeicPartNumber     sql.NullInt32
	)
	ExamID = uuid.NullUUID{UUID: examPart.ExamID, Valid: examPart.ExamID != uuid.Nil}
	PartTitle = examPart.PartTitle
	PlanType = examPart.PlanType

	if examPart.PartOrder != 0 {
		PartOrder = sql.NullInt32{Int32: examPart.PartOrder, Valid: true}
	}
	if examPart.Description != "" {
		Description = sql.NullString{String: examPart.Description, Valid: true}
	}
	IsPracticeComponent = sql.NullBool{Bool: examPart.IsPracticeComponent, Valid: true}

	if examPart.ToeicPartNumber != 0 {
		ToeicPartNumber = sql.NullInt32{Int32: examPart.ToeicPartNumber, Valid: true}
	}

	_, err := r.Queries.CreateExamPart(ctx, database.CreateExamPartParams{
		ExamID:              ExamID,
		PartTitle:           PartTitle,
		PartOrder:           PartOrder,
		Description:         Description,
		IsPracticeComponent: IsPracticeComponent,
		PlanType:            PlanType,
		ToeicPartNumber:     ToeicPartNumber,
	})
	if err != nil {
		logger.Error("LibraryRepository.CreateExamPart: failed to create exam part",
			"error", err)
		return err
	}
	return nil
}

func (r *LibraryRepository) UpdateExamPart(ctx context.Context, examPart *entity.ExamPart, examPartId uuid.UUID) error {
	var (
		ExamID              uuid.NullUUID
		PartTitle           string
		PartOrder           sql.NullInt32
		Description         sql.NullString
		IsPracticeComponent sql.NullBool
		PlanType            string
		ToeicPartNumber     sql.NullInt32
	)

	ExamID = uuid.NullUUID{UUID: examPart.ExamID, Valid: examPart.ExamID != uuid.Nil}
	PartTitle = examPart.PartTitle
	PlanType = examPart.PlanType

	if examPart.PartOrder != 0 {
		PartOrder = sql.NullInt32{Int32: examPart.PartOrder, Valid: true}
	}
	if examPart.Description != "" {
		Description = sql.NullString{String: examPart.Description, Valid: true}
	}
	IsPracticeComponent = sql.NullBool{Bool: examPart.IsPracticeComponent, Valid: true}
	if examPart.ToeicPartNumber != 0 {
		ToeicPartNumber = sql.NullInt32{Int32: examPart.ToeicPartNumber, Valid: true}
	}

	err := r.Queries.UpdateExamPart(ctx, database.UpdateExamPartParams{
		ExamID:              ExamID,
		PartTitle:           PartTitle,
		PartOrder:           PartOrder,
		Description:         Description,
		IsPracticeComponent: IsPracticeComponent,
		PlanType:            PlanType,
		ToeicPartNumber:     ToeicPartNumber,
		PartID:              examPartId,
	})
	if err != nil {
		logger.Error("LibraryRepository.UpdateExamPart: failed to update exam part",
			"exam_part_id", examPartId,
			"error", err)
		return err
	}
	return nil
}

func (r *LibraryRepository) GetExamPart(ctx context.Context, examPartId uuid.UUID) (*entity.ExamPart, error) {
	dbExamPart, err := r.Queries.GetExamPartByID(ctx, examPartId)
	if err != nil {
		logger.Error("LibraryRepository.GetExamPart: failed to retrieve exam part",
			"exam_part_id", examPartId,
			"error", err)
		return nil, err
	}

	return &entity.ExamPart{
		PartID:              dbExamPart.PartID,
		PartTitle:           dbExamPart.PartTitle,
		PartOrder:           dbExamPart.PartOrder.Int32,
		Description:         dbExamPart.Description.String,
		IsPracticeComponent: dbExamPart.IsPracticeComponent.Bool,
		PlanType:            dbExamPart.PlanType,
		CreatedAt:           dbExamPart.CreatedAt.Time,
		UpdatedAt:           dbExamPart.UpdatedAt.Time,
		ToeicPartNumber:     dbExamPart.ToeicPartNumber.Int32,
	}, nil
}

func (r *LibraryRepository) GetPracticeExamParts(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedExamPart, error) {
	totalItems, err := r.Queries.GetPracticeExamPartCount(ctx)
	if err != nil {
		logger.Error("LibraryRepository.GetExamParts: failed to get total count of exam parts",
			"page_number", pageNumber,
			"page_size", pageSize,
			"error", err)
		return nil, err
	}

	offset := (pageNumber - 1) * pageSize
	listParams := database.GetPaginatedPracticeExamPartsParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	}

	dbExamParts, err := r.Queries.GetPaginatedPracticeExamParts(ctx, listParams)
	if err != nil {
		logger.Error("LibraryRepository.GetExamParts: failed to retrieve paginated exam parts",
			"page_number", pageNumber,
			"page_size", pageSize,
			"offset", offset,
			"error", err)
		return nil, err
	}
	var examParts []*entity.ExamPart

	for _, dbExamPart := range dbExamParts {
		examPart := &entity.ExamPart{
			PartID:              dbExamPart.PartID,
			PartTitle:           dbExamPart.PartTitle,
			PartOrder:           dbExamPart.PartOrder.Int32,
			Description:         dbExamPart.Description.String,
			IsPracticeComponent: true,
			PlanType:            dbExamPart.PlanType,
			CreatedAt:           dbExamPart.CreatedAt.Time,
			UpdatedAt:           dbExamPart.UpdatedAt.Time,
			ToeicPartNumber:     dbExamPart.ToeicPartNumber.Int32,
		}
		examParts = append(examParts, examPart)
	}
	totalPages := (totalItems + int64(pageSize) - 1) / int64(pageSize)

	return &entity.PaginatedExamPart{
		Items:       examParts,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
		PageSize:    pageSize,
	}, nil
}
func (r *LibraryRepository) GetExamPartsByExamId(ctx context.Context, examPartId uuid.UUID) ([]*entity.ExamPart, error) {
	dbExamParts, err := r.Queries.GetExamPartsByExamId(ctx, uuid.NullUUID{UUID: examPartId, Valid: true})
	if err != nil {
		logger.Error("LibraryRepository.GetExamParts: failed to retrieve paginated exam parts",
			"error", err)
		return nil, err
	}
	var examParts []*entity.ExamPart
	for _, dbExamPart := range dbExamParts {
		examPart := &entity.ExamPart{
			PartID:              dbExamPart.PartID,
			ExamID:              examPartId,
			PartTitle:           dbExamPart.PartTitle,
			PartOrder:           dbExamPart.PartOrder.Int32,
			Description:         dbExamPart.Description.String,
			IsPracticeComponent: false,
			CreatedAt:           dbExamPart.CreatedAt.Time,
			UpdatedAt:           dbExamPart.UpdatedAt.Time,
			ToeicPartNumber:     dbExamPart.ToeicPartNumber.Int32,
		}
		examParts = append(examParts, examPart)
	}
	return examParts, nil
}
