package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"pirate-lang-go/core/constants"
	"pirate-lang-go/core/errors"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/account/dto"
	"pirate-lang-go/modules/account/mapper"
	"time"
)

func (s *AccountService) CreateAccount(ctx context.Context, requestData *dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errors.AppError) {

	existingUser, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, requestData.Email, requestData.Username, uuid.Nil)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to check existing user", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "", err)
	}

	if existingUser != nil {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(requestData.Password)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to hash password", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Convert DTO to entity
	user := mapper.ToUserEntity(requestData)
	user.Password = hashedPassword

	// Save to database
	createdUser, err := s.repo.CreateAccount(ctx, user)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to create account", "error", err, "email", requestData.Email)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Generate access token (expires in 1 day)
	accessToken, err := utils.GenerateToken(createdUser.ID, createdUser.Email, createdUser.UserName, constants.AccessTokenExpiry)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to generate access token", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Generate refresh token (expires in 7 days)
	refreshToken, err := utils.GenerateToken(createdUser.ID, createdUser.Email, createdUser.UserName, constants.RefreshTokenExpiry)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to generate refresh token", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Prepare response
	response := &dto.CreateAccountResponse{
		Username:     createdUser.UserName,
		Email:        createdUser.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}
func (s *AccountService) Login(ctx context.Context, requestData *dto.LoginRequest) (*dto.LoginResponse, *errors.AppError) {

	// Check rate limiting first
	rateKey := fmt.Sprintf("login_rate:%s", requestData.Email)
	rateCount, _ := s.cache.Get(ctx, rateKey).Int()
	if rateCount >= 10 { // Max 10 attempts per minute
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount: reached rate limit", nil)
	}

	// Increment and set rate limit
	s.cache.Incr(ctx, rateKey)
	s.cache.Expire(ctx, rateKey, time.Minute) // Reset after 1 minute

	// Check if user is blocked
	blockKey := fmt.Sprintf("login_blocked:%s", requestData.Email)
	blocked, err := s.cache.IsLoginBlocked(ctx, blockKey)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount: cannot connect redis", err)
	}
	if blocked {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:login blocked", err)
	}

	existingUser, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, requestData.Email, "", uuid.Nil)

	if err != nil {
		logger.Error("AccountService:Login:Failed to check existing user", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:User not found", err)
	}

	// Check credentials
	if existingUser == nil {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:User not found", err)
	}
	if existingUser.IsLocked {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:User is locker", err)
	}
	if !utils.ComparePassword(existingUser.Password, requestData.Password) {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:Wrong password", err)
	}

	// Generate access token (expires in 1 day)
	accessToken, err := utils.GenerateToken(existingUser.ID, existingUser.Email, existingUser.UserName, 24*time.Duration(time.Hour))
	if err != nil {
		logger.Error("AccountService:Login:Failed to generate access token", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:cannot connect redis generate", err)
	}
	// Generate refresh token (expires in 7 days)
	refreshToken, err := utils.GenerateToken(existingUser.ID, existingUser.Email, existingUser.UserName, 7*24*time.Duration(time.Hour))
	if err != nil {
		logger.Error("AccountService:Login:Failed to generate refresh token", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:cannot connect redis refresh", err)
	}
	// Prepare response
	response := &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}
func (s *AccountService) Logout(ctx context.Context, token string) *errors.AppError {
	// Add token to blacklist with expiry matching the token's expiry
	claims, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AccountService:Logout:Failed to validate token", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	// Calculate remaining time until token expiry
	expiryTime := time.Until(time.Unix(claims.ExpiresAt.Unix(), 0)).Seconds()
	if expiryTime <= 0 {
		return nil // Token already expired
	}

	// TODO: Add to blacklist
	//s.cache.AddToBlacklist(ctx, token, time.Duration(expiryTime)*time.Second)
	return nil
}
func (s *AccountService) ChangePassword(ctx context.Context, token string, requestData *dto.ChangePasswordRequest) *errors.AppError {
	// Validate and parse token
	claims, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AccountService:ChangePassword:Failed to validate token", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Get user from database
	user, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, "", "", claims.UserID)
	if err != nil {
		logger.Error("AccountService:ChangePassword:Failed to get user", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	if user == nil {
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Verify current password
	if !utils.ComparePassword(user.Password, requestData.CurrentPassword) {
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Hash the new password
	hashedPassword, err := utils.HashPassword(requestData.NewPassword)
	if err != nil {
		logger.Error("AccountService:ChangePassword:Failed to hash new password", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Update password in database
	user.Password = hashedPassword
	err = s.repo.UpdatePassword(ctx, user)
	if err != nil {
		logger.Error("AccountService:ChangePassword:Failed to update password", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	return nil
}
