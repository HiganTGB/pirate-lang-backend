package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/account/dto"
	"pirate-lang-go/modules/account/entity"
	"pirate-lang-go/modules/account/mapper"
	"time"
)

func (s *AccountService) CreateProfile(ctx context.Context, token string, requestData *dto.CreateUserProfile) *errors.AppError {

	claims, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AccountService:CreateProfile:Failed to validate token", "error", err)
		return errors.NewAppError(errors.ErrUnauthorized, "AccountService:CreateProfile:Failed to get user", err)
	}
	err = s.repo.CreateProfile(ctx, mapper.ToProfileEntity(requestData, &claims.UserID))
	if err != nil {
		logger.Error("AccountService:CreateProfile:Failed to create profile", "error", err)
		return errors.NewAppError(errors.ErrUniqueViolation, "AccountService:CreateAccount:Failed to create profile", err)
	}
	return nil
}
func (s *AccountService) UpdateProfile(ctx context.Context, token string, requestData *dto.UpdateUserProfile) *errors.AppError {
	claims, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AccountService:CreateProfile:Failed to validate token", "error", err)
		return errors.NewAppError(errors.ErrUnauthorized, "AccountService:CreateProfile:Failed to get user", err)
	}
	err = s.repo.UpdateProfile(ctx, mapper.ToUpdateProfileEntity(requestData, &claims.UserID))
	if err != nil {
		logger.Error("AccountService:CreateProfile:Failed to update profile", "error", err)
		return errors.NewAppError(errors.ErrUniqueViolation, "AccountService:CreateAccount:Failed to update profile", err)
	}
	return nil
}

func (s *AccountService) GetProfile(ctx context.Context, token string) (*dto.ProfileResponse, *errors.AppError) {
	claims, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AccountService:CreateProfile:Failed to get token", "error", err)
		return nil, errors.NewAppError(errors.ErrUnauthorized, "AccountService:CreateProfile:Failed to get user", err)
	}
	profile, err := s.repo.GetProfile(ctx, claims.UserID)
	if err != nil {
		logger.Error("AccountService:GetProfile:Failed to get profile", "error", err)
		return nil, errors.NewAppError(errors.ErrNotFound, "AccountService:GetProfile:Failed to get profile", err)
	}
	response := mapper.ToProfileResponse(profile, s.storage.BuildAvatarURL(profile.AvatarUrl))
	return response, nil
}
func (s *AccountService) GetManagerProfile(ctx context.Context, userId uuid.UUID) (*entity.UserProfile, *errors.AppError) {
	profile, err := s.repo.GetProfile(ctx, userId)
	if err != nil {
		logger.Error("AccountService:GetProfile:Failed to get profile", "error", err)
		return nil, errors.NewAppError(errors.ErrNotFound, "AccountService:GetProfile:Failed to get profile", err)
	}
	return profile, nil
}

func (s *AccountService) UpdateAvatar(ctx context.Context, file *multipart.FileHeader, token string) (*dto.UpdateUserAvatarResponse, *errors.AppError) {
	claims, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AccountService:UpdateAvatar:Failed to validate token", "error", err)
		return nil, errors.NewAppError(errors.ErrUnauthorized, "AccountService:UpdateAvatar:Failed to get user", err)
	}
	rateKey := fmt.Sprintf("uploadAvatar_rate:%s", claims.UserID)
	rateCount, _ := s.cache.Get(ctx, rateKey).Int()
	if rateCount >= 10 { // Max 10 attempts per minute
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:UpdateAvatar: reached rate limit", nil)
	}
	// Increment and set rate limit
	s.cache.Incr(ctx, rateKey)
	s.cache.Expire(ctx, rateKey, time.Minute)

	src, err := file.Open()
	if err != nil {
		logger.Error("AccountService:CreateProfile:Failed to create profile", "error", err)
		return nil, errors.NewAppError(errors.ErrInvalidInput, "AccountService:UpdateAvatar:Failed to read file", err)
	}
	defer src.Close()
	resizedBytes, err := utils.ResizeImage(src)
	resizedReader := bytes.NewReader(resizedBytes)
	resizedSize := int64(len(resizedBytes))
	name, url, err := s.storage.UploadAvatar(ctx, claims.UserID, resizedReader, resizedSize, file.Filename)
	if err != nil {
		logger.Error("AccountService:UpdateAvatar:Failed to update avatar", "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "AccountService:UpdateAvatar:Failed to update avatar", err)
	}

	err = s.repo.UpdateAvatar(ctx, name, claims.UserID)
	if err != nil {
		logger.Error("AccountService:UpdateAvatar:Failed to update avatar", "error", err)
		return nil, errors.NewAppError(errors.ErrInternal, "AccountService:UpdateAvatar:Failed to update avatar", err)
	}
	response := &dto.UpdateUserAvatarResponse{
		Filename:  name,
		ObjectURL: url,
	}
	return response, nil
}
