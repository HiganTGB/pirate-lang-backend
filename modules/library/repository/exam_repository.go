package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/library/entity"
)

func (r *LibraryRepository) CreateExam(ctx context.Context, exam *entity.Exam) error {
	var (
		ExamTitle         string
		Description       sql.NullString
		DurationMinutes   sql.NullInt32
		ExamType          string
		MaxListeningScore sql.NullInt32
		MaxReadingScore   sql.NullInt32
		MaxSpeakingScore  sql.NullInt32
		MaxWritingScore   sql.NullInt32
		TotalScore        sql.NullInt32
	)
	ExamTitle = exam.ExamTitle
	ExamType = exam.ExamType

	if exam.Description != "" {
		Description = sql.NullString{String: exam.Description, Valid: true}
	}

	if exam.DurationMinutes != 0 {
		DurationMinutes = sql.NullInt32{Int32: exam.DurationMinutes, Valid: true}
	}
	if exam.MaxListeningScore != 0 {
		MaxListeningScore = sql.NullInt32{Int32: exam.MaxListeningScore, Valid: true}
	}
	if exam.MaxReadingScore != 0 {
		MaxReadingScore = sql.NullInt32{Int32: exam.MaxReadingScore, Valid: true}
	}
	if exam.MaxSpeakingScore != 0 {
		MaxSpeakingScore = sql.NullInt32{Int32: exam.MaxSpeakingScore, Valid: true}
	}
	if exam.MaxWritingScore != 0 {
		MaxWritingScore = sql.NullInt32{Int32: exam.MaxWritingScore, Valid: true}
	}
	if exam.TotalScore != 0 {
		TotalScore = sql.NullInt32{Int32: exam.TotalScore, Valid: true}
	}

	_, err := r.Queries.CreateExam(ctx, database.CreateExamParams{
		ExamTitle:         ExamTitle,
		Description:       Description,
		DurationMinutes:   DurationMinutes,
		ExamType:          ExamType,
		MaxListeningScore: MaxListeningScore,
		MaxReadingScore:   MaxReadingScore,
		MaxSpeakingScore:  MaxSpeakingScore,
		MaxWritingScore:   MaxWritingScore,
		TotalScore:        TotalScore,
	})
	if err != nil {
		logger.Error("LibraryRepository.CreateExam: failed to create exam",
			"error", err)
		return err
	}
	return err
}

func (r *LibraryRepository) UpdateExam(ctx context.Context, exam *entity.Exam, examId uuid.UUID) error {
	var (
		ExamTitle         string
		Description       sql.NullString
		DurationMinutes   sql.NullInt32
		ExamType          string
		MaxListeningScore sql.NullInt32
		MaxReadingScore   sql.NullInt32
		MaxSpeakingScore  sql.NullInt32
		MaxWritingScore   sql.NullInt32
		TotalScore        sql.NullInt32
	)
	ExamTitle = exam.ExamTitle
	ExamType = exam.ExamType

	if exam.Description != "" {
		Description = sql.NullString{String: exam.Description, Valid: true}
	}

	if exam.DurationMinutes != 0 {
		DurationMinutes = sql.NullInt32{Int32: exam.DurationMinutes, Valid: true}
	}
	if exam.MaxListeningScore != 0 {
		MaxListeningScore = sql.NullInt32{Int32: exam.MaxListeningScore, Valid: true}
	}
	if exam.MaxReadingScore != 0 {
		MaxReadingScore = sql.NullInt32{Int32: exam.MaxReadingScore, Valid: true}
	}
	if exam.MaxSpeakingScore != 0 {
		MaxSpeakingScore = sql.NullInt32{Int32: exam.MaxSpeakingScore, Valid: true}
	}
	if exam.MaxWritingScore != 0 {
		MaxWritingScore = sql.NullInt32{Int32: exam.MaxWritingScore, Valid: true}
	}
	if exam.TotalScore != 0 {
		TotalScore = sql.NullInt32{Int32: exam.TotalScore, Valid: true}
	}

	err := r.Queries.UpdateExam(ctx, database.UpdateExamParams{
		ExamTitle:         ExamTitle,
		Description:       Description,
		DurationMinutes:   DurationMinutes,
		ExamType:          ExamType,
		MaxListeningScore: MaxListeningScore,
		MaxReadingScore:   MaxReadingScore,
		MaxSpeakingScore:  MaxSpeakingScore,
		MaxWritingScore:   MaxWritingScore,
		TotalScore:        TotalScore,
		ExamID:            examId,
	})
	if err != nil {
		logger.Error("LibraryRepository.UpdateExam: failed to update exam",
			"exam_id", examId,
			"error", err)
		return err
	}
	return err
}

func (r *LibraryRepository) GetExam(ctx context.Context, examId uuid.UUID) (*entity.Exam, error) {
	dbExam, err := r.Queries.GetExam(ctx, examId)
	if err != nil {
		logger.Error("LibraryRepository.GetExam: failed to retrieve exam",
			"exam_id", examId,
			"error", err)
		return nil, err
	}

	getInt32 := func(ni sql.NullInt32) int32 {
		if ni.Valid {
			return ni.Int32
		}
		return 0
	}

	return &entity.Exam{
		ExamID:            dbExam.ExamID,
		ExamTitle:         dbExam.ExamTitle,
		Description:       dbExam.Description.String,
		DurationMinutes:   getInt32(dbExam.DurationMinutes),
		ExamType:          dbExam.ExamType,
		MaxListeningScore: getInt32(dbExam.MaxListeningScore),
		MaxReadingScore:   getInt32(dbExam.MaxReadingScore),
		MaxSpeakingScore:  getInt32(dbExam.MaxSpeakingScore),
		MaxWritingScore:   getInt32(dbExam.MaxWritingScore),
		TotalScore:        getInt32(dbExam.TotalScore),
		CreatedAt:         dbExam.CreatedAt.Time,
		UpdatedAt:         dbExam.UpdatedAt.Time,
	}, err
}

func (r *LibraryRepository) GetExams(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedExams, error) {
	totalItems, err := r.Queries.GetExamsCount(ctx)
	if err != nil {
		logger.Error("LibraryRepository.GetExams: failed to get total count of exams",
			"page_number", pageNumber,
			"page_size", pageSize,
			"error", err)
		return nil, err
	}

	offset := (pageNumber - 1) * pageSize
	listParams := database.GetPaginatedExamsParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	}

	dbExams, err := r.Queries.GetPaginatedExams(ctx, listParams)
	if err != nil {
		logger.Error("LibraryRepository.GetExams: failed to retrieve paginated exams",
			"page_number", pageNumber,
			"page_size", pageSize,
			"offset", offset,
			"error", err)
		return nil, err
	}
	var exams []*entity.Exam
	getInt32 := func(ni sql.NullInt32) int32 {
		if ni.Valid {
			return ni.Int32
		}
		return 0
	}
	for _, dbExam := range dbExams {
		exam := &entity.Exam{
			ExamID:            dbExam.ExamID,
			ExamTitle:         dbExam.ExamTitle,
			Description:       dbExam.Description.String,
			DurationMinutes:   getInt32(dbExam.DurationMinutes),
			ExamType:          dbExam.ExamType,
			MaxListeningScore: getInt32(dbExam.MaxListeningScore),
			MaxReadingScore:   getInt32(dbExam.MaxReadingScore),
			MaxSpeakingScore:  getInt32(dbExam.MaxSpeakingScore),
			MaxWritingScore:   getInt32(dbExam.MaxWritingScore),
			TotalScore:        getInt32(dbExam.TotalScore),
			CreatedAt:         dbExam.CreatedAt.Time,
			UpdatedAt:         dbExam.UpdatedAt.Time,
		}
		exams = append(exams, exam)
	}
	totalPages := (totalItems + int64(pageSize) - 1) / int64(pageSize)

	return &entity.PaginatedExams{
		Items:       exams,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
		PageSize:    pageSize,
	}, nil
}
