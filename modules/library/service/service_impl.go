package service

import (
	"context"
	"github.com/google/uuid"
	"mime/multipart"
	"pirate-lang-go/core/cache"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/storage"
	"pirate-lang-go/modules/library/dto"
	"pirate-lang-go/modules/library/repository"
)

type LibraryService struct {
	repo    repository.ILibraryRepository
	cache   cache.ICache
	storage storage.IStorage
}

func NewLibraryService(repo repository.ILibraryRepository, cache cache.ICache, storage storage.IStorage) ILibraryService {

	return &LibraryService{
		repo:    repo,
		cache:   cache,
		storage: storage,
	}
}

type ILibraryService interface {
	GetExams(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedExamResponse, *errors.AppError)
	CreateExam(ctx context.Context, dataRequest *dto.CreateExamRequest) *errors.AppError
	UpdateExam(ctx context.Context, dataRequest *dto.UpdateExamRequest, examId uuid.UUID) *errors.AppError
	GetExam(ctx context.Context, examId uuid.UUID) (*dto.ExamResponse, *errors.AppError)
	CreateExamPart(ctx context.Context, dataRequest *dto.CreateExamPartRequest) *errors.AppError
	UpdateExamPart(ctx context.Context, dataRequest *dto.UpdateExamPartRequest, examPartId uuid.UUID) *errors.AppError
	GetExamPart(ctx context.Context, examPartId uuid.UUID) (*dto.ExamPartResponse, *errors.AppError)
	GetExamParts(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedExamPartResponse, *errors.AppError)
	GetExamPartsByExamId(ctx context.Context, examId uuid.UUID) ([]*dto.ExamPartResponse, *errors.AppError)
	//CreateQuestionGroup(ctx context.Context, req *dto.CreateQuestionGroupRequest) (string, *errors.AppError)
	//UpdateQuestionGroup(ctx context.Context, groupId uuid.UUID, req *dto.UpdateQuestionGroupRequest) *errors.AppError
	//GetQuestionGroups(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedGroupResponse, *errors.AppError)
	//UploadAudioGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError)
	//UploadTranscriptAudioGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID, language string) (*dto.UpdateContentFileResponse, *errors.AppError)
	//UploadImageGroup(ctx context.Context, file *multipart.FileHeader, groupId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError)
	//DeleteAudioGroup(ctx context.Context, groupId uuid.UUID) *errors.AppError
	//GetQuestionsByGroups(ctx context.Context, groupId uuid.UUID) ([]*dto.QuestionResponse, error)
	//CreateQuestion(ctx context.Context, request *dto.CreateQuestionRequest, groupId uuid.UUID) (*dto.QuestionResponse, error)
	//UpdateQuestion(ctx context.Context, request *dto.UpdateQuestionRequest, questionId uuid.UUID) error
	CreateParagraph(ctx context.Context, dataRequest *dto.CreateParagraphRequest) *errors.AppError
	UpdateParagraph(ctx context.Context, dataRequest *dto.UpdateParagraphRequest, paragraphId uuid.UUID) *errors.AppError
	GetParagraph(ctx context.Context, paragraphId uuid.UUID) (*dto.ParagraphResponse, *errors.AppError)
	GetParagraphsByPartId(ctx context.Context, partId uuid.UUID) ([]*dto.ParagraphResponse, *errors.AppError)
	UploadAudioParagraph(ctx context.Context, file *multipart.FileHeader, paragraphId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError)
	UploadTranscriptAudioParagraph(ctx context.Context, file *multipart.FileHeader, paragraphId uuid.UUID, language string) (*dto.UpdateContentFileResponse, *errors.AppError)
	UploadImageParagraph(ctx context.Context, file *multipart.FileHeader, paragraphId uuid.UUID) (*dto.UpdateContentFileResponse, *errors.AppError)
}
