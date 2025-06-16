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
	GetUserByEmailOrUserNameOrId(ctx context.Context, email, userName string, userId uuid.UUID) (*entity.User, error)
	CreateAccount(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdatePassword(ctx context.Context, user *entity.User) error
	GetUsers(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedUsers, error)
	LockUser(ctx context.Context, userId uuid.UUID, lockReason string) error
	UnlockUser(ctx context.Context, userId uuid.UUID, unlockReason string) error
	// Profile
	CreateProfile(ctx context.Context, profile *entity.UserProfile) error
	UpdateProfile(ctx context.Context, profile *entity.UserProfile) error
	GetProfile(ctx context.Context, userId uuid.UUID) (*entity.UserProfile, *entity.User, error)
	UpdateAvatar(ctx context.Context, updateAvatarUrl string, userID uuid.UUID) error
	GetAvatar(ctx context.Context, userID uuid.UUID) (string, error)
	// Rbac
	CreateRole(ctx context.Context, role *entity.Role) error
	GetRoles(ctx context.Context) ([]*entity.Role, error)
	CreatePermission(ctx context.Context, permission *entity.Permission) error
	GetPermissions(ctx context.Context) ([]*entity.Permission, error)
	AssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error
	AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error
	RoleExists(ctx context.Context, roleID uuid.UUID) (bool, error)
	PermissionExists(ctx context.Context, permissionID uuid.UUID) (bool, error)
	HasPermission(ctx context.Context, userID uuid.UUID, permissionID uuid.UUID) (bool, error)
	DeleteRole(ctx context.Context, roleID uuid.UUID) error
	DeletePermission(ctx context.Context, permissionID uuid.UUID) error
}
