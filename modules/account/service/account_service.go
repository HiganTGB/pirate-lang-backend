package service

import (
	"context"
	"github.com/google/uuid"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/account/dto"
	"pirate-lang-go/modules/account/mapper"
	"time"
)

func (s *AccountService) GetUsers(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedUsersResponse, *errors.AppError) {

	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resultGetUsers, err := s.repo.GetUsers(ctx, pageNumber, pageSize)
	if err != nil {
		logger.Error("AccountService:GetUsers:Failed to get users", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	// Convert to DTO
	usersDTO := mapper.ToPaginatedUsersResponse(resultGetUsers)
	return usersDTO, nil
}
func (s *AccountService) LockUser(ctx context.Context, requestData *dto.LockUserRequest, userId uuid.UUID) *errors.AppError {
	err := s.repo.LockUser(ctx, userId, requestData.LockReason)
	if err != nil {
		logger.Error("AccountService:LockUser:Failed to lock user", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:user is already locked", err)
	}
	return nil
}
func (s *AccountService) UnlockUser(ctx context.Context, requestData *dto.UnlockUserRequest, userId uuid.UUID) *errors.AppError {
	err := s.repo.UnlockUser(ctx, userId, requestData.UnlockReason)
	if err != nil {
		logger.Error("AccountService:LockUser:Failed to unlock user", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:user is not locked", err)
	}
	return nil
}
