package service

import (
	"context"
	"github.com/google/uuid"
	"pirate-lang-go/core/cache"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/storage"
	"pirate-lang-go/modules/library/dto"
	"pirate-lang-go/modules/library/repository"
)

type LibraryService struct {
	repo    repository.ILibraryRepository
	cache   cache.ICache
	storage storage.IStorage
}

func NewLibraryService(repo repository.ILibraryRepository, cache cache.ICache, storage storage.IStorage) ILibraryService {

	return &LibraryService{
		repo:    repo,
		cache:   cache,
		storage: storage,
	}
}

type ILibraryService interface {
	GetPart(ctx context.Context, partId uuid.UUID) (*dto.PartResponse, error)
	GetParts(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedPartResponse, *errors.AppError)
	CreatePart(ctx context.Context, dto *dto.CreatePartRequest) error
	UpdatePart(ctx context.Context, dto *dto.UpdatePartRequest, partId uuid.UUID) error
}
