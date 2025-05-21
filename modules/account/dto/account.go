package dto

import (
	"github.com/google/uuid"
	"pirate-lang-go/core/dto"
)

type CreateAccountRequest struct {
	Username        string `json:"user_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CreateAccountResponse struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	Id            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	IsSocialLogin bool      `json:"is_social_login"`
	IsLocked      bool      `json:"is_locked"`
}

type PaginatedUsersResponse = dto.Pagination[*UserResponse]
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
