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
func ToQuestionGroupEntity(dto *dto.CreateQuestionGroupRequest) *entity.QuestionGroup {
	if dto == nil {
		return nil
	}
	return &entity.QuestionGroup{
		Name:        dto.Name,
		Description: dto.Description,
		PartID:      dto.PartID,
		PlanType:    dto.PlanType,
		GroupType:   dto.GroupType,
	}
}
func ToQuestionGroupEntityForUpdate(dto *dto.UpdateQuestionGroupRequest) *entity.QuestionGroup {
	if dto == nil {
		return nil
	}
	return &entity.QuestionGroup{
		Name:               dto.Name,
		Description:        dto.Description,
		PartID:             dto.PartID,
		PlanType:           dto.PlanType,
		GroupType:          dto.GroupType,
		ContextTextContent: dto.ContextTextContent,
	}
}
func ToQuestionGroupResponse(ent *entity.QuestionGroup) *dto.QuestionGroupResponse {
	if ent == nil {
		return nil
	}
	return &dto.QuestionGroupResponse{
		ID:                 ent.QuestionGroupID,
		Name:               ent.Name,
		Description:        ent.Description,
		PartID:             ent.PartID,
		PlanType:           ent.PlanType,
		GroupType:          ent.GroupType,
		ContextTextContent: ent.ContextTextContent,
		ContextAudioUrl:    ent.ContextAudioURL,
		ContextImageUrl:    ent.ContextImageURL,
	}
}
func ToPaginatedGroupsResponse(groups *entity.PaginatedQuestionGroup) *dto.PaginatedGroupResponse {
	if groups == nil {
		return nil
	}

	partDtos := make([]*dto.QuestionGroupResponse, 0, len(groups.Items))
	for _, group := range groups.Items {

		partDtos = append(partDtos, &dto.QuestionGroupResponse{
			ID:          group.QuestionGroupID,
			Name:        group.Name,
			Description: group.Description,
			PartID:      group.PartID,
			PlanType:    group.PlanType,
			GroupType:   group.GroupType,
		})
	}

	return &dto.PaginatedGroupResponse{
		Items:       partDtos,
		TotalItems:  groups.TotalItems,
		TotalPages:  groups.TotalPages,
		CurrentPage: groups.CurrentPage,
		PageSize:    groups.PageSize,
	}
}
