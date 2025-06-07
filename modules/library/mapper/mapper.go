package mapper

import (
	"pirate-lang-go/modules/library/dto"
	"pirate-lang-go/modules/library/entity"
)

func ToPaginatedPartsResponse(parts *entity.PaginatedParts) *dto.PaginatedPartResponse {
	if parts == nil {
		return nil
	}

	partDtos := make([]*dto.PartResponse, 0, len(parts.Items))
	for _, part := range parts.Items {
		partDtos = append(partDtos, &dto.PartResponse{
			ID:          part.ID,
			Name:        part.Name,
			Skill:       part.Skill,
			Description: part.Description,
			Sequence:    part.Sequence,
			CreatedAt:   part.CreatedAt,
		})
	}

	return &dto.PaginatedPartResponse{
		Items:       partDtos,
		TotalItems:  parts.TotalItems,
		TotalPages:  parts.TotalPages,
		CurrentPage: parts.CurrentPage,
		PageSize:    parts.PageSize,
	}
}
func ToPartResponse(part *entity.Part) *dto.PartResponse {
	if part == nil {
		return nil
	}
	return &dto.PartResponse{
		ID:          part.ID,
		Name:        part.Name,
		Skill:       part.Skill,
		Description: part.Description,
		Sequence:    part.Sequence,
		CreatedAt:   part.CreatedAt,
	}
}
func ToCreatePartEntity(part *dto.CreatePartRequest) *entity.Part {
	if part == nil {
		return nil
	}
	return &entity.Part{
		Name:        part.Name,
		Skill:       part.Skill,
		Description: part.Description,
		Sequence:    part.Sequence,
	}
}
func ToUpdatePartEntity(part *dto.UpdatePartRequest) *entity.Part {
	if part == nil {
		return nil
	}
	return &entity.Part{
		Name:        part.Name,
		Skill:       part.Skill,
		Description: part.Description,
		Sequence:    part.Sequence,
	}
}
