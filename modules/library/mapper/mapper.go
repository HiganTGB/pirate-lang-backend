package mapper

import (
	"encoding/json"
	"fmt"
	"pirate-lang-go/modules/library/dto"
	"pirate-lang-go/modules/library/entity"
)

func ToCreateExamEntity(req *dto.CreateExamRequest) *entity.Exam {
	if req == nil {
		return nil
	}
	return &entity.Exam{
		ExamTitle:         req.ExamTitle,
		Description:       req.Description,
		DurationMinutes:   int32(req.DurationMinutes),
		ExamType:          req.ExamType,
		MaxListeningScore: int32(req.MaxListeningScore),
		MaxReadingScore:   int32(req.MaxReadingScore),
		MaxSpeakingScore:  int32(req.MaxSpeakingScore),
		MaxWritingScore:   req.MaxWritingScore,
		TotalScore:        int32(req.MaxListeningScore + req.MaxReadingScore + req.MaxSpeakingScore + req.MaxWritingScore), // Example calculation
	}
}

func ToUpdateExamEntity(req *dto.UpdateExamRequest) *entity.Exam {
	if req == nil {
		return nil
	}

	return &entity.Exam{
		ExamTitle:         req.ExamTitle,
		Description:       req.Description,
		DurationMinutes:   req.DurationMinutes,
		ExamType:          req.ExamType,
		MaxListeningScore: req.MaxListeningScore,
		MaxReadingScore:   req.MaxReadingScore,
		MaxSpeakingScore:  req.MaxSpeakingScore,
		MaxWritingScore:   req.MaxWritingScore,
		TotalScore:        req.MaxListeningScore + req.MaxReadingScore + req.MaxSpeakingScore + req.MaxWritingScore,
	}
}

func ToExamResponse(exam *entity.Exam) *dto.ExamResponse {
	if exam == nil {
		return nil
	}
	return &dto.ExamResponse{
		ExamID:            exam.ExamID,
		ExamTitle:         exam.ExamTitle,
		Description:       exam.Description,
		DurationMinutes:   int32(exam.DurationMinutes),
		ExamType:          exam.ExamType,
		MaxListeningScore: int32(exam.MaxListeningScore),
		MaxReadingScore:   int32(exam.MaxReadingScore),
		MaxSpeakingScore:  int32(exam.MaxSpeakingScore),
		MaxWritingScore:   int32(exam.MaxWritingScore),
		TotalScore:        int32(exam.TotalScore),
		CreatedAt:         exam.CreatedAt,
		UpdatedAt:         exam.UpdatedAt,
	}
}

func ToPaginatedExamsResponse(exams *entity.PaginatedExams) *dto.PaginatedExamResponse {
	if exams == nil {
		return nil
	}

	examDTOs := make([]*dto.ExamResponse, 0, len(exams.Items))
	for _, exam := range exams.Items {
		examDTOs = append(examDTOs, &dto.ExamResponse{
			ExamID:            exam.ExamID,
			ExamTitle:         exam.ExamTitle,
			Description:       exam.Description,
			DurationMinutes:   int32(exam.DurationMinutes),
			ExamType:          exam.ExamType,
			MaxListeningScore: int32(exam.MaxListeningScore),
			MaxReadingScore:   int32(exam.MaxReadingScore),
			MaxSpeakingScore:  int32(exam.MaxSpeakingScore),
			MaxWritingScore:   int32(exam.MaxWritingScore),
			TotalScore:        int32(exam.TotalScore),
			CreatedAt:         exam.CreatedAt,
			UpdatedAt:         exam.UpdatedAt,
		})
	}

	return &dto.PaginatedExamResponse{
		Items:       examDTOs,
		TotalItems:  exams.TotalItems,
		TotalPages:  exams.TotalPages,
		CurrentPage: exams.CurrentPage,
		PageSize:    exams.PageSize,
	}
}
func ToCreateExamPartEntity(dto *dto.CreateExamPartRequest) *entity.ExamPart {
	if dto == nil {
		return nil
	}
	return &entity.ExamPart{
		ExamID:              dto.ExamID.UUID,
		PartTitle:           dto.PartTitle,
		PartOrder:           dto.PartOrder,
		Description:         dto.Description,
		IsPracticeComponent: dto.IsPracticeComponent,
		PlanType:            dto.PlanType,
		ToeicPartNumber:     dto.ToeicPartNumber,
	}
}

func ToUpdateExamPartEntity(dto *dto.UpdateExamPartRequest) *entity.ExamPart {
	if dto == nil {
		return nil
	}
	return &entity.ExamPart{
		ExamID:              dto.ExamID.UUID,
		PartTitle:           dto.PartTitle,
		PartOrder:           dto.PartOrder,
		Description:         dto.Description,
		IsPracticeComponent: dto.IsPracticeComponent,
		PlanType:            dto.PlanType,
		ToeicPartNumber:     dto.ToeicPartNumber,
	}
}
func ToExamPartResponse(entity *entity.ExamPart) *dto.ExamPartResponse {
	if entity == nil {
		return nil
	}
	return &dto.ExamPartResponse{
		PartTitle:           entity.PartTitle,
		PartOrder:           entity.PartOrder,
		Description:         entity.Description,
		IsPracticeComponent: entity.IsPracticeComponent,
		PlanType:            entity.PlanType,
		ToeicPartNumber:     entity.ToeicPartNumber,
	}
}
func ToPaginatedExamPartsResponse(parts *entity.PaginatedExamPart) *dto.PaginatedExamPartResponse {
	if parts == nil {
		return nil
	}

	examDTOs := make([]*dto.ExamPartResponse, 0, len(parts.Items))
	for _, exam := range parts.Items {
		examDTOs = append(examDTOs, &dto.ExamPartResponse{
			PartTitle:           exam.PartTitle,
			PartOrder:           exam.PartOrder,
			Description:         exam.Description,
			IsPracticeComponent: exam.IsPracticeComponent,
			PlanType:            exam.PlanType,
			ToeicPartNumber:     exam.ToeicPartNumber,
			CreatedAt:           exam.CreatedAt,
			UpdatedAt:           exam.UpdatedAt,
		})
	}

	return &dto.PaginatedExamPartResponse{
		Items:       examDTOs,
		TotalItems:  parts.TotalItems,
		TotalPages:  parts.TotalPages,
		CurrentPage: parts.CurrentPage,
		PageSize:    parts.PageSize,
	}
}
func ToCreateParagraphEntity(dto *dto.CreateParagraphRequest) *entity.Paragraph {
	if dto == nil {
		return nil
	}
	return &entity.Paragraph{
		ParagraphContent: dto.ParagraphContent,
		Title:            dto.Title,
		PartID:           dto.PartID,
		ParagraphOrder:   dto.ParagraphOrder,
		ParagraphType:    dto.ParagraphType,
		AudioUrl:         dto.AudioUrl,
		ImageUrl:         dto.ImageUrl,
	}
}

func ToUpdateParagraphEntity(dto *dto.UpdateParagraphRequest) *entity.Paragraph {
	if dto == nil {
		return nil
	}
	return &entity.Paragraph{
		ParagraphContent: dto.ParagraphContent,
		Title:            dto.Title,
		PartID:           dto.PartID,
		ParagraphOrder:   dto.ParagraphOrder,
		ParagraphType:    dto.ParagraphType,
	}
}

func ToParagraphResponse(ent *entity.Paragraph) *dto.ParagraphResponse {
	if ent == nil {
		return nil
	}
	return &dto.ParagraphResponse{
		ParagraphID:      ent.ParagraphID,
		ParagraphContent: ent.ParagraphContent,
		Title:            ent.Title,
		PartID:           ent.PartID,
		ParagraphOrder:   ent.ParagraphOrder,
		ParagraphType:    ent.ParagraphType,
		AudioUrl:         ent.AudioUrl,
		ImageUrl:         ent.ImageUrl,
		CreatedAt:        ent.CreatedAt,
		UpdatedAt:        ent.UpdatedAt,
	}
}

func MarshalAnswerOption(opt *dto.AnswerOption) (string, error) {
	jsonData, err := json.Marshal(opt)
	if err != nil {
		return "", fmt.Errorf("failed to marshal AnswerOption to JSON: %w", err)
	}
	return string(jsonData), nil
}

func UnmarshalAnswerOption(jsonString string) (dto.AnswerOption, error) {
	var opt dto.AnswerOption
	if jsonString == "" {
		return opt, nil
	}

	err := json.Unmarshal([]byte(jsonString), &opt)
	if err != nil {
		return dto.AnswerOption{}, fmt.Errorf("failed to unmarshal JSON to AnswerOption: %w", err)
	}
	return opt, nil
}

func ToCreateQuestionEntity(dto *dto.CreateQuestionRequest) *entity.Question {
	if dto == nil {
		return nil
	}
	return &entity.Question{
		QuestionContent:      dto.QuestionContent,
		QuestionType:         dto.QuestionType,
		PartID:               dto.PartID,
		ParagraphID:          dto.ParagraphID,
		QuestionOrder:        dto.QuestionOrder,
		ToeicQuestionSection: dto.ToeicQuestionSection,
		QuestionNumberInPart: dto.QuestionNumberInPart,
		AnswerOption:         dto.AnswerOption,
	}
}
func ToUpdateQuestionEntity(dto *dto.UpdateQuestionRequest) *entity.Question {
	if dto == nil {
		return nil
	}
	return &entity.Question{
		QuestionContent:      dto.QuestionContent,
		QuestionType:         dto.QuestionType,
		PartID:               dto.PartID,
		ParagraphID:          dto.ParagraphID,
		QuestionOrder:        dto.QuestionOrder,
		ToeicQuestionSection: dto.ToeicQuestionSection,
		QuestionNumberInPart: dto.QuestionNumberInPart,
		AnswerOption:         dto.AnswerOption,
	}
}

func ToQuestionResponse(entity *entity.Question) *dto.QuestionResponse {
	if entity == nil {
		return nil
	}

	var answerOption, err = UnmarshalAnswerOption(entity.AnswerOption)
	if err != nil {
		answerOption = dto.AnswerOption{}
	}

	return &dto.QuestionResponse{
		QuestionID:           entity.QuestionID,
		QuestionContent:      entity.QuestionContent,
		ParagraphID:          entity.ParagraphID,
		QuestionNumberInPart: entity.QuestionNumberInPart,
		AnswerOption:         answerOption,
		CreatedAt:            entity.CreatedAt,
		UpdatedAt:            entity.UpdatedAt,
	}
}
func ToPaginatedQuestionResponse(parts *entity.PaginatedQuestion) *dto.PaginatedQuestionResponse {
	if parts == nil {
		return nil
	}

	dtOs := make([]*dto.QuestionResponse, 0, len(parts.Items))
	for _, item := range parts.Items {
		dtOs = append(dtOs, ToQuestionResponse(item))
	}

	return &dto.PaginatedQuestionResponse{
		Items:       dtOs,
		TotalItems:  parts.TotalItems,
		TotalPages:  parts.TotalPages,
		CurrentPage: parts.CurrentPage,
		PageSize:    parts.PageSize,
	}
}
