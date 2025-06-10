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
	imageBucket    string
	audioBucket    string
	questionBucket string
	endpoint       string
	useSSL         bool
}

// Constants for bucket names
const (
	ImageBucket    = "images"
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

		// Test the connection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err = client.ListBuckets(ctx) // Try to list buckets to verify connection
		if err != nil {
			serviceErr = fmt.Errorf("failed to connect to MinIO: %w", err)
			return
		}
		logger.Info("Successfully connected to MinIO.")

		// Initialize and ensure buckets exist
		bucketsToCreate := []string{ImageBucket, AudioBucket, QuestionBucket}
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
			imageBucket:    ImageBucket,
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
func buildObjectName(folder string, originalFilename string, name string) string {
	ext := filepath.Ext(originalFilename)
	uniqueID := uuid.New().String()
	if folder != "" {
		return fmt.Sprintf("%s/%s%s", folder, name, ext)
	}
	return fmt.Sprintf("%s%s", uniqueID, ext)
}

// buildObjectURL constructs a public URL for an object
func (s *Storage) buildObjectURL(bucket, objectName string) string {
	scheme := "http"
	if s.useSSL {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s/%s/%s", scheme, s.endpoint, bucket, objectName)
}

// BuildAvatarURL constructs a public URL for an object
func (s *Storage) BuildAvatarURL(objectName string) string {
	return s.buildObjectURL(ImageBucket, objectName)
}

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

// UploadAvatar uploads a new avatar image for a user
// Returns the new object name and its URL
func (s *Storage) UploadAvatar(ctx context.Context, userID uuid.UUID, file io.Reader, fileSize int64, filename string) (string, string, error) {

	objectName := buildObjectName("avatars", filename, userID.String())
	contentType := getContentType(filename)
	_, err := s.uploadFile(ctx, s.imageBucket, objectName, file, fileSize, contentType)
	if err != nil {
		return "", "", err
	}
	return objectName, s.buildObjectURL(s.imageBucket, objectName), nil
}

func (s *Storage) getObject(ctx context.Context, bucket, objectName string) (*minio.Object, minio.ObjectInfo, error) {
	object, err := s.client.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		errResponse := minio.ToErrorResponse(err)
		if errResponse.Code == "NoSuchKey" {
			logger.Warn(fmt.Sprintf("Object '%s' not found in bucket '%s'.", objectName, bucket))
			return nil, minio.ObjectInfo{}, fmt.Errorf("object '%s' not found: %w", objectName, err)
		}
		logger.Error(fmt.Sprintf("Failed to get object '%s' from bucket '%s': %v", objectName, bucket, err))
		return nil, minio.ObjectInfo{}, fmt.Errorf("failed to get object '%s': %w", objectName, err)
	}

	objectInfo, err := object.Stat()
	if err != nil {

		_ = object.Close()
		logger.Error(fmt.Sprintf("Failed to stat object '%s' from bucket '%s': %v", objectName, bucket, err))
		return nil, minio.ObjectInfo{}, fmt.Errorf("failed to get object info for '%s': %w", objectName, err)
	}

	logger.Info(fmt.Sprintf("Successfully retrieved object '%s' from bucket '%s'. Size: %d", objectName, bucket, objectInfo.Size))
	return object, objectInfo, nil
}
