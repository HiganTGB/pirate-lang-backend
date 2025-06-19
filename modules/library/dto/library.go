package dto

import (
	"github.com/google/uuid"
	"pirate-lang-go/core/dto"
	"time"
)

type PartResponse struct {
	ID          uuid.UUID `json:"id"`
	Skill       string    `json:"skill"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Sequence    int       `json:"sequence"`
	CreatedAt   time.Time `json:"created_at"`
}
type CreatePartRequest struct {
	Skill       string `json:"skill"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
}

type UpdatePartRequest struct {
	Skill       string `json:"skill"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
}
type PaginatedPartResponse = dto.Pagination[*PartResponse]
type CreateQuestionGroupRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PartID      uuid.UUID `json:"part_id"`
	PlanType    string    `json:"plan_type"`
	GroupType   string    `json:"group_type"`
}
type UpdateQuestionGroupRequest struct {
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	PartID             uuid.UUID `json:"part_id"`
	PlanType           string    `json:"plan_type"`
	GroupType          string    `json:"group_type"`
	ContextTextContent string    `json:"context_text_content"`
}
type QuestionGroupResponse struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	PartID             uuid.UUID `json:"part_id"`
	PlanType           string    `json:"plan_type"`
	GroupType          string    `json:"group_type"`
	ContextTextContent string    `json:"context_text_content"`
	ContextAudioUrl    string    `json:"context_audio_url"`
	ContextImageUrl    string    `json:"context_image_url"`
	CreatedAt          time.Time `json:"created_at"`
	IsLocked           bool      `json:"is_locked" `
	LockedAt           time.Time `json:"locked_at"`
	LockReason         string    `json:"lock_reason" `
	UnlockedAt         time.Time `json:"unlocked_at" `
	UnlockReason       string    `json:"unlock_reason"`
}
type PaginatedGroupResponse = dto.Pagination[*QuestionGroupResponse]
type UpdateQuestionContentResponse struct {
	Filename  string `json:"filename"`
	ObjectURL string `json:"object_url"`
}
type UpdateContentFileResponse struct {
	Filename  string `json:"original_filename"`
	ObjectURL string `json:"object_url"`
}
