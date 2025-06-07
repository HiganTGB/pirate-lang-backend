package dto

import (
	"github.com/google/uuid"
	"pirate-lang-go/core/dto"
	"time"
)

type PartResponse struct {
	ID          uuid.UUID `json:"part_id"`
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
