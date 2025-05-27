package service

import (
	"context"
	"github.com/google/uuid"
	"pirate-lang-go/core/cache"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/modules/account/dto"
	"pirate-lang-go/modules/account/repository"
)

type AccountService struct {
	repo  repository.IAccountRepository
	cache cache.ICache
}

func NewAccountService(repo repository.IAccountRepository, cache cache.ICache) IAccountService {

	return &AccountService{
		repo:  repo,
		cache: cache,
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

	// Rbac API
	CreateRole(ctx context.Context, role *dto.CreateRoleRequest) *errors.AppError
	GetRoles(ctx context.Context) ([]*dto.RoleResponse, *errors.AppError)
	CreatePermission(ctx context.Context, permission *dto.CreatePermissionRequest) *errors.AppError
	GetPermissions(ctx context.Context) ([]*dto.PermissionResponse, *errors.AppError)
	AssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) *errors.AppError
	AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) *errors.AppError
	HasPermission(ctx context.Context, userID uuid.UUID, permissionID uuid.UUID) (bool, *errors.AppError)
}
