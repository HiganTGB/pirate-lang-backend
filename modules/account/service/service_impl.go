package service

import (
	"context"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/modules/account/dto"
	"pirate-lang-go/modules/account/repository"
)

type AccountService struct {
	repo repository.IAccountRepository
}

func NewAccountService(repo repository.IAccountRepository) IAccountService {

	return &AccountService{
		repo: repo,
	}
}

type IAccountService interface {

	// Auth API
	CreateAccount(ctx context.Context, requestData *dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errors.AppError)
	ChangePassword(ctx context.Context, token string, requestData *dto.ChangePasswordRequest) *errors.AppError
	Login(ctx context.Context, requestData *dto.LoginRequest) (*dto.LoginResponse, *errors.AppError)
	Logout(ctx context.Context, token string) *errors.AppError
	// Admin API
	GetUsers(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedUsersResponse, *errors.AppError)
}
