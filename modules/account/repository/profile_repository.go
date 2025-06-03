package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"pirate-lang-go/core/logger"
	"pirate-lang-go/internal/database"
	"pirate-lang-go/modules/account/entity"
	"strings"
	"time"
)

func (r *AccountRepository) CreateProfile(ctx context.Context, profile *entity.UserProfile) error {
	var (
		userId      uuid.UUID
		fullName    sql.NullString
		birthday    sql.NullTime
		gender      sql.NullString
		phoneNumber sql.NullString
		address     sql.NullString
		bio         sql.NullString
	)
	userId = profile.UserId
	fullName = sql.NullString{String: profile.FullName, Valid: true}
	birthday = sql.NullTime{Time: *profile.Birthday, Valid: true}
	gender = sql.NullString{String: strings.ToLower(profile.Gender), Valid: true}
	phoneNumber = sql.NullString{String: profile.PhoneNumber, Valid: true}
	if profile.Address != "" {
		address = sql.NullString{String: profile.Address, Valid: true}
	}
	if profile.Bio != "" {
		bio = sql.NullString{String: profile.Bio, Valid: true}
	}
	params := database.CreateUserProfileParams{
		UserID:      userId,
		FullName:    fullName,
		Birthday:    birthday,
		Gender:      gender,
		PhoneNumber: phoneNumber,
		Address:     address,
		Bio:         bio,
	}
	err := r.Queries.CreateUserProfile(ctx, params)
	if err != nil {
		logger.Error("AccountRepository:CreateProfile:Error when create userProfile", "error", err)
		return err
	}
	return err
}
func (r *AccountRepository) GetProfile(ctx context.Context, userId uuid.UUID) (*entity.UserProfile, error) {
	dbProfile, err := r.Queries.GetUserProfile(ctx, userId)
	if err != nil {
		logger.Error("AccountRepository:GetProfile:Error when get userProfile", "error", err)
		return nil, err
	}
	var birthdayPtr *time.Time
	if dbProfile.Birthday.Valid {
		validTime := dbProfile.Birthday.Time
		birthdayPtr = &validTime
	}
	profile := &entity.UserProfile{
		UserId:      userId,
		FullName:    dbProfile.FullName.String,
		Birthday:    birthdayPtr,
		Gender:      dbProfile.Gender.String,
		PhoneNumber: dbProfile.PhoneNumber.String,
		Address:     dbProfile.Address.String,
		AvatarUrl:   dbProfile.AvatarUrl.String,
		Bio:         dbProfile.Bio.String,
	}
	return profile, err

}
func (r *AccountRepository) UpdateProfile(ctx context.Context, profile *entity.UserProfile) error {
	var (
		userId      uuid.UUID
		fullName    sql.NullString
		birthday    sql.NullTime
		gender      sql.NullString
		phoneNumber sql.NullString
		address     sql.NullString
		bio         sql.NullString
	)
	userId = profile.UserId
	fullName = sql.NullString{String: profile.FullName, Valid: true}
	birthday = sql.NullTime{Time: *profile.Birthday, Valid: true}
	gender = sql.NullString{String: strings.ToLower(profile.Gender), Valid: true}
	phoneNumber = sql.NullString{String: profile.PhoneNumber, Valid: true}
	if profile.Address != "" {
		address = sql.NullString{String: profile.Address, Valid: true}
	}
	if profile.Bio != "" {
		bio = sql.NullString{String: profile.Bio, Valid: true}
	}
	params := database.UpdateUserProfileParams{
		UserID:      userId,
		FullName:    fullName,
		Birthday:    birthday,
		Gender:      gender,
		PhoneNumber: phoneNumber,
		Address:     address,
		Bio:         bio,
	}
	err := r.Queries.UpdateUserProfile(ctx, params)
	if err != nil {
		logger.Error("AccountRepository:CreateProfile:Error when create userProfile", "error", err)
		return err
	}
	return err
}
func (r *AccountRepository) UpdateAvatar(ctx context.Context, updateAvatarUrl string, userID uuid.UUID) error {
	var (
		avatarUrl sql.NullString
	)
	if updateAvatarUrl != "" {
		avatarUrl = sql.NullString{String: updateAvatarUrl, Valid: true}
	}
	params := database.UpdateUserAvatarParams{
		UserID:    userID,
		AvatarUrl: avatarUrl,
	}
	err := r.Queries.UpdateUserAvatar(ctx, params)
	if err != nil {
		logger.Error("AccountRepository:UpdateAvatar:Error when create userAvatar", "error", err)
		return err
	}
	return err
}
func (r *AccountRepository) GetAvatar(ctx context.Context, userID uuid.UUID) (string, error) {
	avatar, err := r.Queries.GetUserAvatar(ctx, userID)
	if err != nil {
		logger.Error("AccountRepository:UpdateAvatar:Error when create userAvatar", "error", err)
		return "", err
	}
	return avatar.String, err
}
