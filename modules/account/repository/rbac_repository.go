package repository

import (
	"context"
	"database/sql"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/account/entity"

	"github.com/google/uuid"
)

func (r *AccountRepository) CreateRole(ctx context.Context, role *entity.Role) error {
	var (
		name            string
		nullDescription sql.NullString
	)

	name = role.Name
	if role.Description != "" {
		nullDescription = sql.NullString{String: role.Description, Valid: true}
	}

	err := r.Queries.CreateRole(ctx, database.CreateRoleParams{
		Name:        name,
		Description: nullDescription,
	})

	if err != nil {
		logger.Error("AccountRepository:CreateRole:", "error", err)
		return err
	}
	logger.Info("AccountRepository:CreateRole:", err)
	return err
}

func (r *AccountRepository) GetRoles(ctx context.Context) ([]*entity.Role, error) {
	var roles []*entity.Role
	dbRoles, err := r.Queries.GetRoles(ctx)
	for _, dbRole := range dbRoles {
		roles = append(roles, &entity.Role{
			Id:          dbRole.ID,
			Name:        dbRole.Name,
			Description: dbRole.Description.String,
		})
	}
	if err != nil {
		logger.Error("AccountRepository:GetRoles:", err)
		return nil, err
	}
	return roles, err
}

func (r *AccountRepository) CreatePermission(ctx context.Context, permission *entity.Permission) error {
	var (
		name            string
		nullDescription sql.NullString
	)
	name = permission.Name
	if permission.Description != "" {
		nullDescription = sql.NullString{String: permission.Description, Valid: true}
	}
	err := r.Queries.CreatePermission(ctx, database.CreatePermissionParams{
		Name:        name,
		Description: nullDescription,
	})

	if err != nil {
		logger.Error("AccountRepository:CreatePermission:", err)
		return err
	}
	return err
}

func (r *AccountRepository) GetPermissions(ctx context.Context) ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	dbPermission, err := r.Queries.GetPermissions(ctx)
	if err != nil {
		logger.Error("AccountRepository:GetPermissions:", err)
		return nil, err
	}
	for _, dbPermission := range dbPermission {
		permissions = append(permissions, &entity.Permission{
			Id:          dbPermission.ID,
			Name:        dbPermission.Name,
			Description: dbPermission.Description.String,
		})
	}
	return permissions, err
}

func (r *AccountRepository) AssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error {
	err := r.Queries.AssignPermissionToRole(ctx, database.AssignPermissionToRoleParams{RoleID: roleID, PermissionID: permissionID})
	logger.Info("AccountRepository:AssignPermissionToRole:", err)
	return err
}

func (r *AccountRepository) AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	err := r.Queries.AssignRoleToUser(ctx, database.AssignRoleToUserParams{UserID: userID, RoleID: roleID})
	logger.Error("AccountRepository:AssignRoleToUser:", err)
	return err
}

func (r *AccountRepository) RoleExists(ctx context.Context, roleID uuid.UUID) (bool, error) {
	var exists bool

	exists, err := r.Queries.RoleExists(ctx, roleID)
	if err != nil {
		logger.Error("AccountRepository:RoleExists:", err)
		return false, err
	}
	return exists, nil
}

func (r *AccountRepository) PermissionExists(ctx context.Context, permissionID uuid.UUID) (bool, error) {
	var exists bool
	exists, err := r.Queries.PermissionExists(ctx, permissionID)
	if err != nil {
		logger.Error("AccountRepository:PermissionExists:", err)
		return false, err
	}
	return exists, nil
}

func (r *AccountRepository) DeleteRole(ctx context.Context, roleID uuid.UUID) error {
	err := r.Queries.DeleteRole(ctx, roleID)
	if err != nil {
		logger.Error("AccountRepository:DeleteRole:", err)
	}
	return err
}

func (r *AccountRepository) DeletePermission(ctx context.Context, permissionID uuid.UUID) error {
	err := r.Queries.DeletePermission(ctx, permissionID)
	if err != nil {
		logger.Error("AccountRepository:DeletePermission:", err)
	}
	return err
}

func (r *AccountRepository) HasPermission(ctx context.Context, userID uuid.UUID, permissionID uuid.UUID) (bool, error) {
	var exists bool

	exists, err := r.Queries.HasPermission(ctx, database.HasPermissionParams{
		UserID: userID,
		ID:     permissionID,
	})
	if err != nil {
		logger.Error("AccountRepository:HasPermission:", err)
		return false, err
	}
	return exists, nil
}
