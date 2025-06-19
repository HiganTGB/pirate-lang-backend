package service

import (
	"context"
	"github.com/google/uuid"
	"mime/multipart"
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
	CreateQuestionGroup(ctx context.Context, req *dto.CreateQuestionGroupRequest) (string, *errors.AppError)
	UpdateQuestionGroup(ctx context.Context, groupId uuid.UUID, req *dto.UpdateQuestionGroupRequest) *errors.AppError
	GetQuestionGroups(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedGroupResponse, *errors.AppError)
	UploadAudioGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError)
	UploadTranscriptAudioGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID, language string) (*dto.UpdateContentFileResponse, *errors.AppError)
	UploadImageGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError)
	DeleteAudioGroup(ctx context.Context, groupId uuid.UUID) *errors.AppError
}
