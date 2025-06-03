package service

import (
	"context"
	"github.com/google/uuid"
	"mime/multipart"
	"pirate-lang-go/core/cache"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/storage"
	"pirate-lang-go/modules/account/dto"
	"pirate-lang-go/modules/account/entity"
	"pirate-lang-go/modules/account/repository"
)

type AccountService struct {
	repo    repository.IAccountRepository
	cache   cache.ICache
	storage storage.IStorage
}

func NewAccountService(repo repository.IAccountRepository, cache cache.ICache, storage storage.IStorage) IAccountService {

	return &AccountService{
		repo:    repo,
		cache:   cache,
		storage: storage,
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
	GetManagerProfile(ctx context.Context, userId uuid.UUID) (*entity.UserProfile, *errors.AppError)
	LockUser(ctx context.Context, requestData *dto.LockUserRequest, userId uuid.UUID) *errors.AppError
	UnlockUser(ctx context.Context, requestData *dto.UnlockUserRequest, userId uuid.UUID) *errors.AppError
	// UserProfile
	GetProfile(ctx context.Context, token string) (*dto.ProfileResponse, *errors.AppError)
	CreateProfile(ctx context.Context, token string, requestData *dto.CreateUserProfile) *errors.AppError
	UpdateProfile(ctx context.Context, token string, requestData *dto.UpdateUserProfile) *errors.AppError
	UpdateAvatar(ctx context.Context, file *multipart.FileHeader, token string) (*dto.UpdateUserAvatarResponse, *errors.AppError)
	// Rbac API
	CreateRole(ctx context.Context, role *dto.CreateRoleRequest) *errors.AppError
	GetRoles(ctx context.Context) ([]*dto.RoleResponse, *errors.AppError)
	CreatePermission(ctx context.Context, permission *dto.CreatePermissionRequest) *errors.AppError
	GetPermissions(ctx context.Context) ([]*dto.PermissionResponse, *errors.AppError)
	AssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) *errors.AppError
	AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) *errors.AppError
	HasPermission(ctx context.Context, userID uuid.UUID, permissionID uuid.UUID) (bool, *errors.AppError)
}
