package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/library/entity"
)

func (r *LibraryRepository) CreatePart(ctx context.Context, part *entity.Part) error {
	var (
		Skill       string
		Name        string
		Description sql.NullString
		Sequence    int32
	)
	Skill = part.Skill
	Name = part.Name
	Sequence = int32(part.Sequence)
	if part.Description != "" {
		Description = sql.NullString{String: part.Description, Valid: true}
	}

	err := r.Queries.CreatePart(ctx, database.CreatePartParams{
		Skill:       Skill,
		Name:        Name,
		Description: Description,
		Sequence:    Sequence,
	})
	if err != nil {
		logger.Error("PartRepository:CreatePart:", "error", err)
		return err
	}
	return err
}
func (r *LibraryRepository) UpdatePart(ctx context.Context, part *entity.Part, partId uuid.UUID) error {
	var (
		Skill       string
		Name        string
		Description sql.NullString
		Sequence    int32
	)
	Skill = part.Skill
	Name = part.Name
	Sequence = int32(part.Sequence)
	if part.Description != "" {
		Description = sql.NullString{String: part.Description, Valid: true}
	}

	err := r.Queries.UpdatePart(ctx, database.UpdatePartParams{
		Skill:       Skill,
		Name:        Name,
		Description: Description,
		Sequence:    Sequence,
		PartID:      partId,
	})
	if err != nil {
		logger.Error("PartRepository:CreatePart:", "error", err)
		return err
	}
	return err
}
func (r *LibraryRepository) GetPart(ctx context.Context, partId uuid.UUID) (*entity.Part, error) {
	dbPart, err := r.Queries.GetPart(ctx, partId)
	if err != nil {
		logger.Error("PartRepository:GetParts:", "error", err)
		return nil, err
	}
	return &entity.Part{
		ID:          dbPart.PartID,
		Name:        dbPart.Name,
		Description: dbPart.Description.String,
		Sequence:    int(dbPart.Sequence),
		Skill:       dbPart.Skill,
		CreatedAt:   dbPart.CreatedAt.Time,
		UpdatedAt:   dbPart.UpdatedAt.Time,
	}, err
}
func (r *LibraryRepository) GetParts(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedParts, error) {
	// Get total count
	totalItems, err := r.Queries.GetPartsCount(ctx)
	if err != nil {
		logger.Error("PartRepository:GetParts:Error when count parts", "error", err)
		return nil, err
	}

	offset := (pageNumber - 1) * pageSize
	// Get paginated Parts
	listParams := database.GetPaginatedPartsParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	}

	dbParts, err := r.Queries.GetPaginatedParts(ctx, listParams)
	if err != nil {
		logger.Error("PartRepository:GetParts:Error when get parts", "error", err)
		return nil, err
	}
	var parts []*entity.Part
	for _, dbPart := range dbParts {
		part := &entity.Part{
			ID:          dbPart.PartID,
			Skill:       dbPart.Skill,
			Name:        dbPart.Name,
			Description: dbPart.Description.String,
			Sequence:    int(dbPart.Sequence),
			CreatedAt:   dbPart.CreatedAt.Time,
			UpdatedAt:   dbPart.UpdatedAt.Time,
		}
		parts = append(parts, part)
	}
	totalPages := (totalItems + int64(pageSize) - 1) / int64(pageSize)

	return &entity.PaginatedParts{
		Items:       parts,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
		PageSize:    pageSize,
	}, nil
}
