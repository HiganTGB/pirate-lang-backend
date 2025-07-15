package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/library/entity"
)

func (r *LibraryRepository) GetQuestionsByParagraph(ctx context.Context, paragraphId uuid.UUID) ([]*entity.Question, error) {
	questionDBs, err := r.Queries.ListQuestionsByParagraphID(ctx, uuid.NullUUID{UUID: paragraphId})
	if err != nil {
		logger.Error("LibraryRepository:UpdateQuestionGroup: failed to get questions from group",
			"group_id", paragraphId,
			"error", err)
		return nil, err
	}
	var questions []*entity.Question
	for _, questionDB := range questionDBs {
		question := &entity.Question{
			QuestionID:           questionDB.QuestionID,
			QuestionContent:      questionDB.QuestionContent,
			PartID:               questionDB.PartID,
			ParagraphID:          paragraphId,
			QuestionOrder:        questionDB.QuestionOrder,
			AudioUrl:             questionDB.AudioUrl.String,
			ImageUrl:             questionDB.ImageUrl.String,
			ToeicQuestionSection: questionDB.ToeicQuestionSection,
			QuestionNumberInPart: questionDB.QuestionNumberInPart.Int32,
			QuestionType:         questionDB.QuestionType,
		}
		questions = append(questions, question)
	}
	return questions, nil
}
func (r *LibraryRepository) GetSeparateQuestionsByPart(ctx context.Context, partId uuid.UUID, pageNumber, pageSize int) (*entity.PaginatedQuestion, error) {
	totalItems, err := r.Queries.GetCountSeparateQuestionsByPartID(ctx, partId)
	if err != nil {
		logger.Error("LibraryRepository.GetExams: failed to get total count of exams",
			"page_number", pageNumber,
			"page_size", pageSize,
			"error", err)
		return nil, err
	}
	offset := (pageNumber - 1) * pageSize
	listParams := database.GetPaginatedSeparateQuestionsByPartIDParams{
		PartID: partId,
		Limit:  int32(pageSize),
		Offset: int32(offset),
	}
	questionDBs, err := r.Queries.GetPaginatedSeparateQuestionsByPartID(ctx, listParams)
	if err != nil {
		logger.Error("LibraryRepository:UpdateQuestionGroup: failed to get questions from group",
			"group_id", partId,
			"error", err)
		return nil, err
	}
	var questions []*entity.Question
	for _, questionDB := range questionDBs {
		question := &entity.Question{
			QuestionID:           questionDB.QuestionID,
			QuestionContent:      questionDB.QuestionContent,
			PartID:               questionDB.PartID,
			ParagraphID:          partId,
			QuestionOrder:        questionDB.QuestionOrder,
			AudioUrl:             questionDB.AudioUrl.String,
			ImageUrl:             questionDB.ImageUrl.String,
			ToeicQuestionSection: questionDB.ToeicQuestionSection,
			QuestionNumberInPart: questionDB.QuestionNumberInPart.Int32,
			QuestionType:         questionDB.QuestionType,
		}
		questions = append(questions, question)
	}
	totalPages := (totalItems + int64(pageSize) - 1) / int64(pageSize)
	return &entity.PaginatedQuestion{
		Items:       questions,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
		PageSize:    pageSize,
	}, nil
}
func (r *LibraryRepository) CreateQuestion(ctx context.Context, questionRequest *entity.Question) (*entity.Question, error) {

	var (
		questionContent      string
		questionType         string
		partID               uuid.UUID
		paragraphID          uuid.NullUUID
		questionOrder        int32
		toeicQuestionSection string
		questionNumberInPart sql.NullInt32
		answerOption         pqtype.NullRawMessage
		correctAnswer        sql.NullString
	)
	questionContent = questionRequest.QuestionContent
	questionType = questionRequest.QuestionType
	partID = questionRequest.PartID
	questionOrder = questionRequest.QuestionOrder
	toeicQuestionSection = questionRequest.ToeicQuestionSection
	paragraphID = uuid.NullUUID{UUID: questionRequest.ParagraphID, Valid: questionRequest.ParagraphID != uuid.Nil}
	answerOption = pqtype.NullRawMessage{RawMessage: []byte(questionRequest.AnswerOption), Valid: questionRequest.AnswerOption != ""}
	correctAnswer = sql.NullString{String: questionRequest.CorrectAnswer, Valid: true}
	questionNumberInPart = sql.NullInt32{Int32: questionRequest.QuestionNumberInPart, Valid: true}
	params := database.CreateQuestionParams{
		QuestionContent:      questionContent,
		QuestionType:         questionType,
		PartID:               partID,
		ParagraphID:          paragraphID,
		QuestionOrder:        questionOrder,
		ToeicQuestionSection: toeicQuestionSection,
		QuestionNumberInPart: questionNumberInPart,
		AnswerOption:         answerOption,
		CorrectAnswer:        correctAnswer,
	}
	questionDB, err := r.Queries.CreateQuestion(ctx, params)
	if err != nil {
		logger.Error("LibraryRepository:CreateQuestion: failed to create question")
	}
	return &entity.Question{
		QuestionID:           questionDB.QuestionID,
		QuestionOrder:        questionDB.QuestionOrder,
		QuestionContent:      questionDB.QuestionContent,
		PartID:               questionDB.PartID,
		ParagraphID:          questionDB.PartID,
		QuestionNumberInPart: questionDB.QuestionNumberInPart.Int32,
		QuestionType:         questionDB.QuestionType,
	}, nil
}

func (r *LibraryRepository) UpdateQuestion(ctx context.Context, questionRequest *entity.Question, questionId uuid.UUID) error {

	var (
		questionContent      string
		questionType         string
		partID               uuid.UUID
		paragraphID          uuid.NullUUID
		questionOrder        int32
		toeicQuestionSection string
		questionNumberInPart sql.NullInt32
		answerOption         pqtype.NullRawMessage
		correctAnswer        sql.NullString
	)
	questionContent = questionRequest.QuestionContent
	questionType = questionRequest.QuestionType
	partID = questionRequest.PartID
	questionOrder = questionRequest.QuestionOrder
	toeicQuestionSection = questionRequest.ToeicQuestionSection
	paragraphID = uuid.NullUUID{UUID: questionRequest.ParagraphID, Valid: questionRequest.ParagraphID != uuid.Nil}
	answerOption = pqtype.NullRawMessage{RawMessage: []byte(questionRequest.AnswerOption), Valid: questionRequest.AnswerOption != ""}
	correctAnswer = sql.NullString{String: questionRequest.CorrectAnswer, Valid: true}
	questionNumberInPart = sql.NullInt32{Int32: questionRequest.QuestionNumberInPart, Valid: true}
	params := database.UpdateQuestionParams{
		QuestionID:           questionId,
		QuestionContent:      questionContent,
		QuestionType:         questionType,
		PartID:               partID,
		ParagraphID:          paragraphID,
		QuestionOrder:        questionOrder,
		ToeicQuestionSection: toeicQuestionSection,
		QuestionNumberInPart: questionNumberInPart,
		AnswerOption:         answerOption,
		CorrectAnswer:        correctAnswer,
	}
	err := r.Queries.UpdateQuestion(ctx, params)
	if err != nil {
		logger.Error("LibraryRepository:CreateQuestion: failed to create question")
		return err
	}
	return nil
}
func (r *LibraryRepository) GetQuestion(ctx context.Context, questionId uuid.UUID) (*entity.Question, error) {

	questionDB, err := r.Queries.GetQuestionByID(ctx, questionId)
	if err != nil {
		logger.Error("LibraryRepository:CreateQuestion: failed to create question")
		return nil, err
	}
	return &entity.Question{
		QuestionID:           questionDB.QuestionID,
		QuestionOrder:        questionDB.QuestionOrder,
		QuestionContent:      questionDB.QuestionContent,
		PartID:               questionDB.PartID,
		ParagraphID:          questionDB.PartID,
		QuestionNumberInPart: questionDB.QuestionNumberInPart.Int32,
		QuestionType:         questionDB.QuestionType,
	}, nil
}
func (r *LibraryRepository) UpdateQuestionAudioUrl(ctx context.Context, url *string, questionId uuid.UUID) error {
	var audioUrl sql.NullString

	if url != nil {
		audioUrl = sql.NullString{String: *url, Valid: true}
	} else {
		audioUrl = sql.NullString{Valid: false}
	}
	params := database.UpdateQuestionAudioURLParams{
		QuestionID: questionId,
		AudioUrl:   audioUrl,
	}
	err := r.Queries.UpdateQuestionAudioURL(ctx, params)
	if err != nil {
		logger.Error("LibraryRepository:CreateQuestion: failed to create question")
	}
	return nil
}
func (r *LibraryRepository) UpdateQuestionImageUrl(ctx context.Context, url *string, questionId uuid.UUID) error {
	var imageUrl sql.NullString

	if url != nil {
		imageUrl = sql.NullString{String: *url, Valid: true}
	} else {
		imageUrl = sql.NullString{Valid: false}
	}
	params := database.UpdateQuestionImageURLParams{
		QuestionID: questionId,
		ImageUrl:   imageUrl,
	}
	err := r.Queries.UpdateQuestionImageURL(ctx, params)
	if err != nil {
		logger.Error("LibraryRepository:CreateQuestion: failed to create question")
	}
	return nil
}
