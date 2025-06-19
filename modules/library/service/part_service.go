package service

import (
	"context"
	"github.com/google/uuid"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/library/dto"
	"pirate-lang-go/modules/library/mapper"
	"time"
)

func (s *LibraryService) GetParts(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedPartResponse, *errors.AppError) {

	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resultGetParts, err := s.repo.GetParts(ctx, pageNumber, pageSize)
	if err != nil {
		logger.Error("LibraryService:GetParts:Failed to get parts", "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "LibraryService:GetParts:Failed to get parts", err)
	}
	// Convert to DTO
	partDTOs := mapper.ToPaginatedPartsResponse(resultGetParts)
	return partDTOs, nil
}
func (s *LibraryService) CreatePart(ctx context.Context, dto *dto.CreatePartRequest) error {
	err := s.repo.CreatePart(ctx, mapper.ToCreatePartEntity(dto))
	if err != nil {
		logger.Error("LibraryService:CreatePart:Failed to create parts", "error", err)
	}
	return err
}
func (s *LibraryService) UpdatePart(ctx context.Context, dto *dto.UpdatePartRequest, partId uuid.UUID) error {
	err := s.repo.UpdatePart(ctx, mapper.ToUpdatePartEntity(dto), partId)
	if err != nil {
		logger.Error("LibraryService:CreatePart:Failed to create parts", "error", err)
	}
	return err
}
func (s *LibraryService) GetPart(ctx context.Context, partId uuid.UUID) (*dto.PartResponse, error) {
	part, err := s.repo.GetPart(ctx, partId)
	if err != nil {
		logger.Error("LibraryService:CreatePart:Failed to create parts", "error", err)
	}
	partDTOs := mapper.ToPartResponse(part)
	return partDTOs, err
}
