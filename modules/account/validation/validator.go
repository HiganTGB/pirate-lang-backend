package validator

import (
	"pirate-lang-go/core/utils"
	"pirate-lang-go/core/validation"
	"pirate-lang-go/modules/account/dto"
)

func ValidateCreateAccount(dataRequest dto.CreateAccountRequest) *validation.ValidationResult {
	result := validation.NewValidationResult()

	// Validate email
	switch {
	case utils.IsEmpty(dataRequest.Email):
		result.AddError("email", "Email is required")
	case !utils.IsValidEmail(dataRequest.Email):
		result.AddError("email", "Invalid email format")
	case !utils.IsValidEmailDomain(dataRequest.Email):
		result.AddError("email", "Invalid email domain")
	}

	// Validate password and confirmation
	if len(dataRequest.Password) < 8 {
		result.AddError("password", "Password must be at least 8 characters")
	}
	if dataRequest.Password != dataRequest.ConfirmPassword {
		result.AddError("confirm_password", "Password and confirmation do not match")
	}

	// Validate username
	if utils.IsEmpty(dataRequest.Username) {
		result.AddError("username", "Username is required")
	}

	return result
}
func ValidateLogin(dataRequest dto.LoginRequest) *validation.ValidationResult {
	result := validation.NewValidationResult()

	// Validate email
	switch {
	case utils.IsEmpty(dataRequest.Email):
		result.AddError("email", "Email is required")
	case !utils.IsValidEmail(dataRequest.Email):
		result.AddError("email", "Invalid email format")
	case !utils.IsValidEmailDomain(dataRequest.Email):
		result.AddError("email", "Invalid email domain")
	}

	// Validate password
	if len(dataRequest.Password) < 8 {
		result.AddError("password", "Password must be at least 8 characters")
	}

	return result
}
func ValidateChangePassword(dataRequest dto.ChangePasswordRequest) *validation.ValidationResult {
	result := validation.NewValidationResult()
	// Validate old password
	if len(dataRequest.CurrentPassword) < 8 {
		result.AddError("old_password", "Old password must be at least 8 characters")
	}
	// Validate new password and confirmation
	if len(dataRequest.NewPassword) < 8 {
		result.AddError("new_password", "New password must be at least 8 characters")
	}
	if dataRequest.NewPassword != dataRequest.ConfirmPassword {
		result.AddError("confirm_password", "New password and confirmation do not match")
	}

	return result
}
