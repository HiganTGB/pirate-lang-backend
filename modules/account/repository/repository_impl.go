package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/account/entity"
)

type AccountRepository struct {
	Queries *database.Queries
}

func NewAccountRepository(sqlDB *sql.DB) IAccountRepository {
	return &AccountRepository{
		Queries: database.New(sqlDB),
	}
}

type IAccountRepository interface {
	CreateAccount(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUsers(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedUsers, error)
	GetUserByEmailOrUserNameOrId(ctx context.Context, email, userName string, userId uuid.UUID) (*entity.User, error)
	UpdatePassword(ctx context.Context, user *entity.User) error
}
