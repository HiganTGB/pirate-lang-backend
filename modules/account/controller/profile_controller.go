package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/account/dto"
	validator "pirate-lang-go/modules/account/validation"
)

func (controller *AccountController) CreateProfiles(c echo.Context) error {
	ctx := c.Request().Context()
	token, errToken := utils.GetTokenFromHeader(c)
	if errToken != nil {
		return controller.Unauthorized("Unauthorized", errToken)
	}
	// Parse request body
	requestData := new(dto.CreateUserProfile)

	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidator := validator.ValidateCreateUserProfile(requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	err := controller.accountService.CreateProfile(ctx, token, requestData)
	if err != nil {
		return controller.BadRequest("Error create profile", err)
	}
	return controller.SuccessResponse(c, nil, "Create Profile successfully")

}
func (controller *AccountController) UpdateProfiles(c echo.Context) error {
	ctx := c.Request().Context()
	token, errToken := utils.GetTokenFromHeader(c)
	if errToken != nil {
		return controller.Unauthorized("Unauthorized", errToken)
	}
	// Parse request body
	requestData := new(dto.UpdateUserProfile)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidator := validator.ValidateUpdateUserProfile(requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	err := controller.accountService.UpdateProfile(ctx, token, requestData)
	if err != nil {
		return controller.BadRequest("Error update profile", err)
	}
	return controller.SuccessResponse(c, nil, "Update Profile successfully")
}
func (controller *AccountController) GetProfile(c echo.Context) error {
	ctx := c.Request().Context()
	token, errToken := utils.GetTokenFromHeader(c)
	if errToken != nil {
		return controller.Unauthorized("Unauthorized", errToken)
	}
	profile, err := controller.accountService.GetProfile(ctx, token)
	if err != nil {
		return controller.BadRequest("Error get profile", err)
	}
	return controller.SuccessResponse(c, profile, "Get Profile successfully")
}
func (controller *AccountController) UpdateAvatar(c echo.Context) error {
	file, errFile := c.FormFile("avatar")
	if errFile != nil {
		return controller.BadRequest(fmt.Sprintf("Error getting image file: %v", errFile))
	}
	ctx := c.Request().Context()
	token, errToken := utils.GetTokenFromHeader(c)
	if errToken != nil {
		return controller.Unauthorized("Unauthorized", errToken)
	}
	resultUpdateAvatar, err := controller.accountService.UpdateAvatar(ctx, file, token)
	if err != nil {
		return controller.BadRequest(fmt.Sprintf("Error updating avatar: %v", err))
	}
	return controller.SuccessResponse(c, resultUpdateAvatar, "Update Avatar successfully")
}
