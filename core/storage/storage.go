package storage

import (
	"context"
	"fmt"
	"io"
	"mime"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"pirate-lang-go/core/logger"
)

var (
	instance *Storage
	once     sync.Once
)

type Storage struct {
	client         *minio.Client
	avatarBucket   string
	audioBucket    string
	questionBucket string
	endpoint       string
	useSSL         bool
}

// Constants for bucket names
const (
	AvatarBucket   = "avatars"
	AudioBucket    = "audio-files"
	QuestionBucket = "question-images"
)

// NewMinIOService initializes and returns a new MinIO service instance
func NewStorageClient(addr, access string, secret string, ssl bool) (*Storage, error) {
	var serviceErr error
	once.Do(func() {
		client, err := minio.New(addr, &minio.Options{
			Creds:  credentials.NewStaticV4(access, secret, ""),
			Secure: ssl,
		})
		if err != nil {
			serviceErr = fmt.Errorf("failed to initialize MinIO client: %w", err)
			return
		}

		// Test the connection (optional but good practice)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err = client.ListBuckets(ctx) // Try to list buckets to verify connection
		if err != nil {
			serviceErr = fmt.Errorf("failed to connect to MinIO: %w", err)
			return
		}
		logger.Info("Successfully connected to MinIO.")

		// Initialize and ensure buckets exist
		bucketsToCreate := []string{AvatarBucket, AudioBucket, QuestionBucket}
		for _, b := range bucketsToCreate {
			found, err := client.BucketExists(ctx, b)
			if err != nil {
				logger.Error(fmt.Sprintf("Warning: Failed to check if MinIO bucket '%s' exists: %v", b, err))
			}
			if !found {
				logger.Info(fmt.Sprintf("MinIO bucket '%s' does not exist. Attempting to create...", b))
				err = client.MakeBucket(ctx, b, minio.MakeBucketOptions{})
				if err != nil {
					serviceErr = fmt.Errorf("failed to create MinIO bucket '%s': %w", b, err)
					return // Propagate error and stop initialization
				}
				logger.Info(fmt.Sprintf("MinIO bucket '%s' created successfully.", b))
			}
		}

		instance = &Storage{
			client:         client,
			avatarBucket:   AvatarBucket,
			audioBucket:    AudioBucket,
			questionBucket: QuestionBucket,
			endpoint:       addr,
			useSSL:         ssl,
		}
	})

	return instance, serviceErr
}

// getContentType guesses the content type from file extension
func getContentType(filename string) string {
	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		return "application/octet-stream" // Default if cannot guess
	}
	return contentType
}

// generateObjectName generates a unique object name using UUID
func generateObjectName(prefix string, originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	uniqueID := uuid.New().String()
	if prefix != "" {
		return fmt.Sprintf("%s/%s%s", prefix, uniqueID, ext)
	}
	return fmt.Sprintf("%s%s", uniqueID, ext)
}

// buildObjectURL constructs a public URL for an object
func (s *Storage) buildObjectURL(bucket, objectName string) string {
	// You might need to adjust this URL format based on your MinIO deployment
	// (e.g., if behind a proxy, CDN, or using a virtual-hosted style access)
	// For local MinIO, this usually works.
	scheme := "http"
	if s.useSSL {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s/%s/%s", scheme, s.endpoint, bucket, objectName)
}

// --- Generic Upload/Replace/Delete Helper Functions ---

// uploadFile uploads a file to a specific bucket with a given object name
func (s *Storage) uploadFile(ctx context.Context, bucket, objectName string, file io.Reader, fileSize int64, contentType string) (minio.UploadInfo, error) {
	uploadInfo, err := s.client.PutObject(ctx, bucket, objectName, file, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return minio.UploadInfo{}, fmt.Errorf("failed to upload object '%s' to bucket '%s': %w", objectName, bucket, err)
	}
	logger.Info(fmt.Sprintf("Successfully uploaded object '%s' to bucket '%s'. Size: %d", objectName, bucket, uploadInfo.Size))
	return uploadInfo, nil
}

// deleteFile removes an object from a specific bucket
func (s *Storage) deleteFile(ctx context.Context, bucket, objectName string) error {
	err := s.client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete object '%s' from bucket '%s': %w", objectName, bucket, err)
	}
	logger.Info(fmt.Sprintf("Successfully deleted object '%s' from bucket '%s'.", objectName, bucket))
	return nil
}

// --- Specific Functions for Avatar, Audio, Question Images ---

// UploadAvatar uploads a new avatar image for a user
// Returns the new object name and its URL
func (s *Storage) UploadAvatar(ctx context.Context, userID uuid.UUID, file io.Reader, fileSize int64, filename string) (string, string, error) {
	// Object name for avatar can be user_id.ext or a unique name under user_id/
	// For simplicity, let's use userID as part of the object name prefix.
	objectName := generateObjectName(userID.String(), filename)
	contentType := getContentType(filename)

	_, err := s.uploadFile(ctx, s.avatarBucket, objectName, file, fileSize, contentType)
	if err != nil {
		return "", "", err
	}
	return objectName, s.buildObjectURL(s.avatarBucket, objectName), nil
}

// ReplaceAvatar replaces an existing avatar image for a user
// oldObjectName is the full path to the old avatar in MinIO (e.g., "userID/uuid.jpg")
func (s *Storage) ReplaceAvatar(ctx context.Context, userID uuid.UUID, oldObjectName string, file io.Reader, fileSize int64, filename string) (string, string, error) {
	// First, delete the old avatar if it exists
	if oldObjectName != "" {
		err := s.deleteFile(ctx, s.avatarBucket, oldObjectName)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to delete old avatar '%s': %v", oldObjectName, err))
			// Decide if this should block the new upload or just log a warning
			// For now, we'll log and proceed with new upload.
		}
	}

	// Then, upload the new avatar
	return s.UploadAvatar(ctx, userID, file, fileSize, filename)
}

// DeleteAvatar deletes a user's avatar image
func (s *Storage) DeleteAvatar(ctx context.Context, objectName string) error {
	return s.deleteFile(ctx, s.avatarBucket, objectName)
}

// UploadQuestionAudio uploads a new audio file
func (s *Storage) UploadQuestionAudio(ctx context.Context, file io.Reader, fileSize int64, filename string) (string, string, error) {
	objectName := generateObjectName("", filename) // No specific prefix for audio example
	contentType := getContentType(filename)

	_, err := s.uploadFile(ctx, s.audioBucket, objectName, file, fileSize, contentType)
	if err != nil {
		return "", "", err
	}
	return objectName, s.buildObjectURL(s.audioBucket, objectName), nil
}

// ReplaceQuestionAudio replaces an existing audio file
func (s *Storage) ReplaceQuestionAudio(ctx context.Context, oldObjectName string, file io.Reader, fileSize int64, filename string) (string, string, error) {
	if oldObjectName != "" {
		err := s.deleteFile(ctx, s.audioBucket, oldObjectName)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to delete old audio '%s': %v", oldObjectName, err))
		}
	}
	return s.UploadQuestionAudio(ctx, file, fileSize, filename)
}

// DeleteQuestionAudio deletes an audio file
func (s *Storage) DeleteQuestionAudio(ctx context.Context, objectName string) error {
	return s.deleteFile(ctx, s.audioBucket, objectName)
}

// UploadQuestionImage uploads a new image for a question
func (s *Storage) UploadQuestionImage(ctx context.Context, questionID uuid.UUID, file io.Reader, fileSize int64, filename string) (string, string, error) {
	// Object name for question image can be questionID/uuid.jpg
	objectName := generateObjectName(questionID.String(), filename)
	contentType := getContentType(filename)

	_, err := s.uploadFile(ctx, s.questionBucket, objectName, file, fileSize, contentType)
	if err != nil {
		return "", "", err
	}
	return objectName, s.buildObjectURL(s.questionBucket, objectName), nil
}

// ReplaceQuestionImage replaces an existing question image
func (s *Storage) ReplaceQuestionImage(ctx context.Context, questionID uuid.UUID, oldObjectName string, file io.Reader, fileSize int64, filename string) (string, string, error) {
	if oldObjectName != "" {
		err := s.deleteFile(ctx, s.questionBucket, oldObjectName)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to delete old question image '%s': %v", oldObjectName, err))
		}
	}
	return s.UploadQuestionImage(ctx, questionID, file, fileSize, filename)
}

// DeleteQuestionImage deletes a question image
func (s *Storage) DeleteQuestionImage(ctx context.Context, objectName string) error {
	return s.deleteFile(ctx, s.questionBucket, objectName)
}
