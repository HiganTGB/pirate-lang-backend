package validation

import (
	"github.com/google/uuid"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/core/validation"
	"pirate-lang-go/modules/library/dto"
)

var ValidParagraphTypes = map[string]bool{
	"READING":   true,
	"LISTENING": true,
	"SPEAKING":  true,
	"WRITING":   true,
}

var ValidPlan = map[string]bool{
	"SUBSCRIPTION": true,
	"FREE":         true,
}
var ValidGroup = map[string]bool{
	"MULTIPLE_CHOICE":        true,
	"MULTIPLE_CHOICE_HIDDEN": true,
	"ESSAY":                  true,
}
var ValidLang = map[string]bool{
	"vn":  true,
	"eng": true,
}
var ValidExamTypes = map[string]bool{
	"TOEIC L&R":    true,
	"TOEIC S&W":    true,
	"TOEIC Bridge": true,
	"General":      true,
}

func ValidateCreateExam(dataRequest *dto.CreateExamRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()

	// Validate ExamTitle
	if utils.IsEmpty(dataRequest.ExamTitle) {
		result.AddError("exam_title", "Exam title is required")
	}

	// Validate DurationMinutes
	if dataRequest.DurationMinutes <= 0 {
		result.AddError("duration_minutes", "Duration minutes must be a positive number")
	}

	// Validate ExamType
	if utils.IsEmpty(dataRequest.ExamType) {
		result.AddError("exam_type", "Exam type is required")
	} else {
		if !ValidExamTypes[dataRequest.ExamType] {
			result.AddError("exam_type", "Exam type must be one of 'Practice', 'MockTest', 'Diagnostic', 'Placement'")
		}
	}

	// Validate Scores (optional: assuming scores can be 0 or positive)
	// If scores must be strictly positive, change <= 0 to < 0
	if dataRequest.MaxListeningScore < 0 {
		result.AddError("max_listening_score", "Max Listening Score cannot be negative")
	}
	if dataRequest.MaxReadingScore < 0 {
		result.AddError("max_reading_score", "Max Reading Score cannot be negative")
	}
	if dataRequest.MaxSpeakingScore < 0 {
		result.AddError("max_speaking_score", "Max Speaking Score cannot be negative")
	}
	if dataRequest.MaxWritingScore < 0 {
		result.AddError("max_writing_score", "Max Writing Score cannot be negative")
	}

	return result
}

func ValidateUpdateExam(dataRequest *dto.UpdateExamRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()

	// Validate ExamTitle (optional for update, depending on business logic)
	if utils.IsEmpty(dataRequest.ExamTitle) {
		result.AddError("exam_title", "Exam title is required for update")
	}

	// Validate DurationMinutes
	if dataRequest.DurationMinutes <= 0 {
		result.AddError("duration_minutes", "Duration minutes must be a positive number")
	}

	// Validate ExamType
	if utils.IsEmpty(dataRequest.ExamType) {
		result.AddError("exam_type", "Exam type is required for update")
	} else {
		if !ValidExamTypes[dataRequest.ExamType] {
			result.AddError("exam_type", "Exam type must be one of 'Practice', 'MockTest', 'Diagnostic', 'Placement'")
		}
	}

	// Validate Scores
	if dataRequest.MaxListeningScore < 0 {
		result.AddError("max_listening_score", "Max Listening Score cannot be negative")
	}
	if dataRequest.MaxReadingScore < 0 {
		result.AddError("max_reading_score", "Max Reading Score cannot be negative")
	}
	if dataRequest.MaxSpeakingScore < 0 {
		result.AddError("max_speaking_score", "Max Speaking Score cannot be negative")
	}
	if dataRequest.MaxWritingScore < 0 {
		result.AddError("max_writing_score", "Max Writing Score cannot be negative")
	}

	return result
}
func ValidateLang(lang string) bool {
	return !ValidLang[lang]
}

func ValidateCreateExamPart(dataRequest *dto.CreateExamPartRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(dataRequest.PartTitle) {
		result.AddError("part_title", "Part title is required")
	}

	if dataRequest.ToeicPartNumber < 0 {
		result.AddError("toeic_part_number", "Toeic Part Number cannot be negative")
	}

	if !dataRequest.IsPracticeComponent {
		if !dataRequest.ExamID.Valid {
			result.AddError("exam_id", "Exam ID is required when IsPracticeComponent is false")
		}

		if dataRequest.PartOrder <= 0 {
			result.AddError("part_order", "Part order must be a positive number when IsPracticeComponent is false")
		}
	} else {
		if utils.IsEmpty(dataRequest.PlanType) {
			result.AddError("plan_type", "Plan type is required")
		} else {
			if !ValidPlan[dataRequest.PlanType] {
				result.AddError("plan_type", "Plan type must be one of 'SUBSCRIPTION' or 'FREE'")
			}
		}

	}

	return result
}
func ValidateUpdateExamPart(dataRequest *dto.UpdateExamPartRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(dataRequest.PartTitle) {
		result.AddError("part_title", "Part title is required")
	}

	if dataRequest.ToeicPartNumber < 0 {
		result.AddError("toeic_part_number", "Toeic Part Number cannot be negative")
	}

	if !dataRequest.IsPracticeComponent {
		if !dataRequest.ExamID.Valid {
			result.AddError("exam_id", "Exam ID is required when IsPracticeComponent is false")
		}

		if dataRequest.PartOrder <= 0 {
			result.AddError("part_order", "Part order must be a positive number when IsPracticeComponent is false")
		}
	} else {
		if utils.IsEmpty(dataRequest.PlanType) {
			result.AddError("plan_type", "Plan type is required")
		} else {
			if !ValidPlan[dataRequest.PlanType] {
				result.AddError("plan_type", "Plan type must be one of 'SUBSCRIPTION' or 'FREE'")
			}
		}

	}

	return result
}
func ValidateCreateParagraph(dataRequest *dto.CreateParagraphRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(dataRequest.ParagraphContent) {
		result.AddError("paragraph_content", "Paragraph content is required")
	}

	if utils.IsEmpty(dataRequest.Title) {
		result.AddError("title", "Title is required")
	}

	if dataRequest.PartID == uuid.Nil {
		result.AddError("part_id", "Part ID is required")
	}

	if dataRequest.ParagraphOrder <= 0 {
		result.AddError("paragraph_order", "Paragraph order must be a positive number")
	}

	if utils.IsEmpty(dataRequest.ParagraphType) {
		result.AddError("paragraph_type", "Paragraph type is required")
	} else {
		if !ValidParagraphTypes[dataRequest.ParagraphType] {
			result.AddError("paragraph_type", "Paragraph type must be one of 'READING', 'LISTENING', 'SPEAKING', 'WRITING'")
		}
	}

	return result
}
func ValidateUpdateParagraph(dataRequest *dto.UpdateParagraphRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()

	if utils.IsEmpty(dataRequest.ParagraphContent) {
		result.AddError("paragraph_content", "Paragraph content is required")
	}

	if utils.IsEmpty(dataRequest.Title) {
		result.AddError("title", "Title is required")
	}

	if dataRequest.PartID == uuid.Nil {
		result.AddError("part_id", "Part ID is required")
	}

	if dataRequest.ParagraphOrder <= 0 {
		result.AddError("paragraph_order", "Paragraph order must be a positive number")
	}

	if utils.IsEmpty(dataRequest.ParagraphType) {
		result.AddError("paragraph_type", "Paragraph type is required")
	} else {
		if !ValidParagraphTypes[dataRequest.ParagraphType] {
			result.AddError("paragraph_type", "Paragraph type must be one of 'READING', 'LISTENING', 'SPEAKING', 'WRITING'")
		}
	}

	return result
}
