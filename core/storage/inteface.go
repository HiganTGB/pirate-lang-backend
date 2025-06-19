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
	// Audio Operations
	UploadAudio(ctx context.Context, id uuid.UUID, file io.Reader, fileSize int64, filename string, folder string) (string, string, error)
	UploadTranscriptAudio(ctx context.Context, id uuid.UUID, file io.Reader, fileSize int64, filename string, folder string, lang string) (string, string, error)
	UploadImage(ctx context.Context, id uuid.UUID, file io.Reader, fileSize int64, filename string, folder string) (string, string, error)
}
