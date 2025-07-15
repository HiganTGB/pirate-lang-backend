package dto

import (
	"github.com/google/uuid"
	"pirate-lang-go/core/entity"
	"time"
)

type AnswerOption struct {
	A *string `json:"A,omitempty"`
	B *string `json:"B,omitempty"`
	C *string `json:"C,omitempty"`
	D *string `json:"D,omitempty"`
}
type ExamResponse struct {
	ExamID            uuid.UUID `json:"exam_id"`
	ExamTitle         string    `json:"exam_title"`
	Description       string    `json:"description"`
	DurationMinutes   int32     `json:"duration_minutes"`
	ExamType          string    `json:"exam_type"`
	MaxListeningScore int32     `json:"max_listening_score"`
	MaxReadingScore   int32     `json:"max_reading_score"`
	MaxSpeakingScore  int32     `json:"max_speaking_score"`
	MaxWritingScore   int32     `json:"max_writing_score"`
	TotalScore        int32     `json:"total_score"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
type PaginatedExamResponse = entity.Pagination[*ExamResponse]
type CreateExamRequest struct {
	ExamTitle         string `json:"exam_title"`
	Description       string `json:"description"`
	DurationMinutes   int32  `json:"duration_minutes"`
	ExamType          string `json:"exam_type"`
	MaxListeningScore int32  `json:"max_listening_score"`
	MaxReadingScore   int32  `json:"max_reading_score"`
	MaxSpeakingScore  int32  `json:"max_speaking_score"`
	MaxWritingScore   int32  `json:"max_writing_score"`
}
type UpdateExamRequest struct {
	ExamTitle         string `json:"exam_title"`
	Description       string `json:"description"`
	DurationMinutes   int32  `json:"duration_minutes"`
	ExamType          string `json:"exam_type"`
	MaxListeningScore int32  `json:"max_listening_score"`
	MaxReadingScore   int32  `json:"max_reading_score"`
	MaxSpeakingScore  int32  `json:"max_speaking_score"`
	MaxWritingScore   int32  `json:"max_writing_score"`
}

type ExamPartResponse struct {
	PartTitle           string    `json:"part_title"`
	PartOrder           int32     `json:"part_order"`
	Description         string    `json:"description"`
	IsPracticeComponent bool      `json:"is_practice_component"`
	PlanType            string    `json:"plan_type"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	ToeicPartNumber     int32     `json:"toeic_part_number"`
}
type PaginatedExamPartResponse = entity.Pagination[*ExamPartResponse]
type CreateExamPartRequest struct {
	ExamID              uuid.NullUUID `json:"exam_id"`
	PartTitle           string        `json:"part_title"`
	PartOrder           int32         `json:"part_order"`
	Description         string        `json:"description"`
	IsPracticeComponent bool          `json:"is_practice_component"`
	PlanType            string        `json:"plan_type"`
	ToeicPartNumber     int32         `json:"toeic_part_number"`
}
type UpdateExamPartRequest struct {
	ExamID              uuid.NullUUID `json:"exam_id"`
	PartTitle           string        `json:"part_title"`
	PartOrder           int32         `json:"part_order"`
	Description         string        `json:"description"`
	IsPracticeComponent bool          `json:"is_practice_component"`
	PlanType            string        `json:"plan_type"`
	ToeicPartNumber     int32         `json:"toeic_part_number"`
}

type CreateParagraphRequest struct {
	ParagraphContent string    `json:"paragraph_content"`
	Title            string    `json:"title"`
	PartID           uuid.UUID `json:"part_id"`
	ParagraphOrder   int32     `json:"paragraph_order"`
	ParagraphType    string    `json:"paragraph_type"`
	AudioUrl         string    `json:"audio_url"`
	ImageUrl         string    `json:"image_url"`
}

type UpdateParagraphRequest struct {
	ParagraphContent string    `json:"paragraph_content"`
	Title            string    `json:"title"`
	PartID           uuid.UUID `json:"part_id"`
	ParagraphOrder   int32     `json:"paragraph_order"`
	ParagraphType    string    `json:"paragraph_type"`
}

type ParagraphResponse struct {
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
type UpdateContentFileResponse struct {
	Filename  string `json:"original_filename"`
	ObjectURL string `json:"object_url"`
}
type QuestionResponse struct {
	QuestionID           uuid.UUID    `json:"question_id"`
	QuestionContent      string       `json:"question_content"`
	QuestionType         string       `json:"question_type"`
	PartID               uuid.UUID    `json:"part_id"`
	ParagraphID          uuid.UUID    `json:"paragraph_id"`
	QuestionOrder        int32        `json:"question_order"`
	AudioUrl             string       `json:"audio_url"`
	ImageUrl             string       `json:"image_url"`
	ToeicQuestionSection string       `json:"toeic_question_section"`
	QuestionNumberInPart int32        `json:"question_number_in_part"`
	AnswerOption         AnswerOption `json:"answer_option"`
	CorrectAnswer        string       `json:"correct_answer"`
	CreatedAt            time.Time    `json:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at"`
}
type CreateQuestionRequest struct {
	QuestionContent      string    `json:"question_content"`
	QuestionType         string    `json:"question_type"`
	PartID               uuid.UUID `json:"part_id"`
	ParagraphID          uuid.UUID `json:"paragraph_id"`
	QuestionOrder        int32     `json:"question_order"`
	ToeicQuestionSection string    `json:"toeic_question_section"`
	QuestionNumberInPart int32     `json:"question_number_in_part"`
	AnswerOption         string    `json:"answer_option"`
	CorrectAnswer        string    `json:"correct_answer"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
type UpdateQuestionRequest struct {
	QuestionContent      string    `json:"question_content"`
	QuestionType         string    `json:"question_type"`
	PartID               uuid.UUID `json:"part_id"`
	ParagraphID          uuid.UUID `json:"paragraph_id"`
	QuestionOrder        int32     `json:"question_order"`
	ToeicQuestionSection string    `json:"toeic_question_section"`
	QuestionNumberInPart int32     `json:"question_number_in_part"`
	AnswerOption         string    `json:"answer_option"`
	CorrectAnswer        string    `json:"correct_answer"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
type PaginatedQuestionResponse = entity.Pagination[*QuestionResponse]
