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
	GetPart(ctx context.Context, partId uuid.UUID) (*entity.Part, error)
	GetParts(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedParts, error)
	CreatePart(ctx context.Context, part *entity.Part) error
	UpdatePart(ctx context.Context, part *entity.Part, partId uuid.UUID) error
	CreateGroupGroup(ctx context.Context, group *entity.QuestionGroup) (*uuid.UUID, error)
	GetQuestionGroups(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedQuestionGroup, error)
	GetAudioGroup(ctx context.Context, groupId uuid.UUID) (string, error)
	UpdateAudioGroup(ctx context.Context, audioUrl *string, groupId uuid.UUID) error
	UpdateImageGroup(ctx context.Context, imageUrl *string, groupId uuid.UUID) error
	UpdateQuestionGroup(ctx context.Context, group *entity.QuestionGroup, groupId uuid.UUID) error
}
