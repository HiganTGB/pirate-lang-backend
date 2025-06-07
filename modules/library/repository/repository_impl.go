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
}
