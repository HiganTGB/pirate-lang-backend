package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/account/entity"
)

func (r *AccountRepository) GetUserByEmailOrUserNameOrId(ctx context.Context, email, userName string, userId uuid.UUID) (*entity.User, error) {
	var (
		nullEmail    sql.NullString
		nullUserName sql.NullString
		nullUserID   uuid.NullUUID
	)

	if email != "" {
		nullEmail = sql.NullString{String: email, Valid: true}
	}

	if userName != "" {
		nullUserName = sql.NullString{String: userName, Valid: true}
	}

	if userId != uuid.Nil {
		nullUserID = uuid.NullUUID{UUID: userId, Valid: true}
	}
	userParams := database.GetUserByEmailOrUserNameOrIdParams{
		Email:    nullEmail,
		UserName: nullUserName,
		ID:       nullUserID,
	}

	dbUser, err := r.Queries.GetUserByEmailOrUserNameOrId(ctx, userParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	var user *entity.User
	user = &entity.User{
		ID:        dbUser.ID,
		UserName:  dbUser.UserName,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
	}
	return user, nil
}

func (r *AccountRepository) GetUsers(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedUsers, error) {

	// Get total count
	totalItems, err := r.Queries.GetUsersCount(ctx)
	if err != nil {
		logger.Error("AccountRepository:GetUsers:Error when count users", "error", err)
		return nil, err
	}

	offset := (pageNumber - 1) * pageSize
	// Get paginated Users
	listParams := database.GetPaginatedUsersParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	}

	dbUsers, err := r.Queries.GetPaginatedUsers(ctx, listParams)
	if err != nil {
		logger.Error("AccountRepository:GetUsers:Error when get users", "error", err)
		return nil, err
	}
	var users []*entity.User
	for _, dbUser := range dbUsers {
		user := &entity.User{
			ID:        dbUser.ID,
			UserName:  dbUser.UserName,
			Email:     dbUser.Email,
			CreatedAt: dbUser.CreatedAt.Time,
			UpdatedAt: dbUser.UpdatedAt.Time,
		}
		users = append(users, user)
	}
	totalPages := (totalItems + int64(pageSize) - 1) / int64(pageSize)

	return &entity.PaginatedUsers{
		Items:       users,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
		PageSize:    pageSize,
	}, nil
}
func (r *AccountRepository) CreateAccount(ctx context.Context, user *entity.User) (*entity.User, error) {

	params := database.CreateAccountParams{
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
	dbUser, err := r.Queries.CreateAccount(ctx, params)
	if err != nil {
		logger.Error("AccountRepository:CreateAccount:Error when creating user", "error", err)
		return nil, err
	}

	createdUser := &entity.User{
		ID:       dbUser.ID,
		UserName: dbUser.UserName,
		Email:    dbUser.Email,
	}

	return createdUser, nil
}
func (r *AccountRepository) UpdatePassword(ctx context.Context, user *entity.User) error {

	params := database.UpdatePasswordParams{
		Password: user.Password,
	}
	result, err := r.Queries.UpdatePassword(ctx, params)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no user found to update")
	}

	return nil
}
