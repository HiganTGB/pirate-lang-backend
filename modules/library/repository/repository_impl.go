package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/library/entity"
)

type LibraryRepository struct {
	Queries *database.Queries
}

func NewLibraryRepository(sqlDB *sql.DB) ILibraryRepository {
	return &LibraryRepository{
		Queries: database.New(sqlDB),
	}
}

type ILibraryRepository interface {
	CreateExam(ctx context.Context, exam *entity.Exam) error
	UpdateExam(ctx context.Context, exam *entity.Exam, examId uuid.UUID) error
	GetExam(ctx context.Context, examId uuid.UUID) (*entity.Exam, error)
	GetExams(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedExams, error)
	//CreateGroupGroup(ctx context.Context, group *entity.QuestionGroup) (*uuid.UUID, error)
	//GetQuestionGroups(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedQuestionGroup, error)
	//GetAudioGroup(ctx context.Context, groupId uuid.UUID) (string, error)
	//UpdateAudioGroup(ctx context.Context, audioUrl *string, groupId uuid.UUID) error
	//UpdateImageGroup(ctx context.Context, imageUrl *string, groupId uuid.UUID) error
	//UpdateQuestionGroup(ctx context.Context, group *entity.QuestionGroup, groupId uuid.UUID) error
	//GetQuestionsByGroups(ctx context.Context, groupId uuid.UUID) ([]*entity.Question, error)
	//CreateQuestion(ctx context.Context, questionRequest *entity.Question, groupId uuid.UUID) (*entity.Question, error)
	//UpdateQuestion(ctx context.Context, questionRequest *entity.Question, questionId uuid.UUID) error
	CreateExamPart(ctx context.Context, examPart *entity.ExamPart) error
	UpdateExamPart(ctx context.Context, examPart *entity.ExamPart, examPartId uuid.UUID) error
	GetExamPart(ctx context.Context, examPartId uuid.UUID) (*entity.ExamPart, error)
	GetExamParts(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedExamPart, error)
	GetExamPartsByExamId(ctx context.Context, examId uuid.UUID) ([]*entity.ExamPart, error)
	CreateParagraph(ctx context.Context, paragraph *entity.Paragraph) error
	UpdateParagraph(ctx context.Context, paragraph *entity.Paragraph, paragraphId uuid.UUID) error
	GetParagraph(ctx context.Context, paragraphId uuid.UUID) (*entity.Paragraph, error)
	GetParagraphsByPartId(ctx context.Context, partId uuid.UUID) ([]*entity.Paragraph, error)
	UpdateAudioParagraph(ctx context.Context, audioUrl *string, paragraphId uuid.UUID) error
	UpdateImageParagraph(ctx context.Context, imageUrl *string, paragraphId uuid.UUID) error
}
