package storage

import (
	"context"
	"io"

	"github.com/google/uuid"
)

// IStorage defines the interface for file storage operations.
type IStorage interface {
	// Avatar Operations
	BuildAvatarURL(objectName string) string
	UploadAvatar(ctx context.Context, userID uuid.UUID, file io.Reader, fileSize int64, filename string) (objectName string, objectURL string, err error)
	ReplaceAvatar(ctx context.Context, userID uuid.UUID, oldObjectName string, file io.Reader, fileSize int64, filename string) (objectName string, objectURL string, err error)
	DeleteAvatar(ctx context.Context, objectName string) error

	// Audio Operations
	UploadQuestionAudio(ctx context.Context, file io.Reader, fileSize int64, filename string) (objectName string, objectURL string, err error)
	ReplaceQuestionAudio(ctx context.Context, oldObjectName string, file io.Reader, fileSize int64, filename string) (objectName string, objectURL string, err error)
	DeleteQuestionAudio(ctx context.Context, objectName string) error

	// Question Image Operations
	UploadQuestionImage(ctx context.Context, questionID uuid.UUID, file io.Reader, fileSize int64, filename string) (objectName string, objectURL string, err error)
	ReplaceQuestionImage(ctx context.Context, questionID uuid.UUID, oldObjectName string, file io.Reader, fileSize int64, filename string) (objectName string, objectURL string, err error)
	DeleteQuestionImage(ctx context.Context, objectName string) error
}
