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
