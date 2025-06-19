package entity

import (
	"github.com/google/uuid"
	"pirate-lang-go/core/entity"
	"time"
)

type Part struct {
	ID          uuid.UUID `db:"part_id"`
	Skill       string    `db:"skill"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Sequence    int       `db:"sequence"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
type PaginatedParts = entity.Pagination[*Part]

type QuestionGroup struct {
	QuestionGroupID    uuid.UUID `db:"question_group_id" `
	Name               string    `db:"name" `
	Description        string    `db:"description" `
	ContextTextContent string    `db:"context_text_content"`
	ContextAudioURL    string    `db:"context_audio_url" `
	ContextImageURL    string    `db:"context_image_url"`
	PartID             uuid.UUID `db:"part_id" `
	GroupType          string    `db:"group_type" `
	CreatedAt          time.Time `db:"created_at" `
	UpdatedAt          time.Time `db:"updated_at"`
	PlanType           string    `db:"plan_type"`
	Version            int       `db:"version" `
	IsLocked           bool      `db:"is_locked" `
	LockedAt           time.Time `db:"locked_at"`
	LockReason         string    `db:"lock_reason" `
	UnlockedAt         time.Time `db:"unlocked_at" `
	UnlockReason       string    `db:"unlock_reason"`
}
type PaginatedQuestionGroup = entity.Pagination[*QuestionGroup]
