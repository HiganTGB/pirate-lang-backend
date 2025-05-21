package service

import (
	"context"
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
