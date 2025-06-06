package service

import (
	"pirate-lang-go/core/cache"
	"pirate-lang-go/core/storage"
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
}
