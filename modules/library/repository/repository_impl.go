package repository

import (
	"database/sql"
	"pirate-lang-go/internal/database"
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
}
