package entity

import (
	"github.com/google/uuid"
	"pirate-lang-go/core/entity"
	"time"
)

type Exam struct {
	ExamID            uuid.UUID `db:"exam_id"`
	ExamTitle         string    `db:"exam_title"`
	Description       string    `db:"description"`
	DurationMinutes   int32     `db:"duration_minutes"`
	ExamType          string    `db:"exam_type"`
	MaxListeningScore int32     `db:"max_listening_score"`
	MaxReadingScore   int32     `db:"max_reading_score"`
	MaxSpeakingScore  int32     `db:"max_speaking_score"`
	MaxWritingScore   int32     `db:"max_writing_score"`
	TotalScore        int32     `db:"total_score"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
type PaginatedExams = entity.Pagination[*Exam]

type ExamPart struct {
	PartID              uuid.UUID `json:"part_id"`
	ExamID              uuid.UUID `json:"exam_id"`
	PartTitle           string    `json:"part_title"`
	PartOrder           int32     `json:"part_order"`
	Description         string    `json:"description"`
	IsPracticeComponent bool      `json:"is_practice_component"`
	PlanType            string    `json:"plan_type"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	ToeicPartNumber     int32     `json:"toeic_part_number"`
}
type PaginatedExamPart = entity.Pagination[*ExamPart]
type Paragraph struct {
	ParagraphID      uuid.UUID `json:"paragraph_id"`
	ParagraphContent string    `json:"paragraph_content"`
	Title            string    `json:"title"`
	PartID           uuid.UUID `json:"part_id"`
	ParagraphOrder   int32     `json:"paragraph_order"`
	ParagraphType    string    `json:"paragraph_type"`
	AudioUrl         string    `json:"audio_url"`
	ImageUrl         string    `json:"image_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
type PaginatedParagraph = entity.Pagination[*Paragraph]
