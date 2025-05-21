package controller

import (
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/account/dto"
	"pirate-lang-go/modules/account/validation"
)

func (controller *AccountController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.CreateAccountRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}

	resultValidator := validator.ValidateCreateAccount(*requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}

	resultCreateAccount, err := controller.accountService.CreateAccount(ctx, requestData)
	if err != nil {
		return controller.InternalServerError("Internal server error", err)
	}

	return controller.SuccessResponse(c, resultCreateAccount, "Create account success")
}
func (controller *AccountController) Login(c echo.Context) error {

	// Parse request body
	requestData := new(dto.LoginRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}

	resultValidator := validator.ValidateLogin(*requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}
	ctx := c.Request().Context()
	resultLogin, err := controller.accountService.Login(ctx, requestData)
	if err != nil {
		return controller.InternalServerError("Internal server error", err)
	}

	return controller.SuccessResponse(c, resultLogin, "Login success")
}
func (controller *AccountController) ChangePassword(c echo.Context) error {
	ctx := c.Request().Context()

	token, errToken := utils.GetTokenFromHeader(c)
	if errToken != nil {
		return controller.Unauthorized("Unauthorized", errToken)
	}

	requestData := new(dto.ChangePasswordRequest)
	if err := c.Bind(requestData); err != nil {
		return controller.BadRequest("Invalid request data", err)
	}

	resultValidator := validator.ValidateChangePassword(*requestData)
	if !resultValidator.Valid {
		return controller.BadRequest("Invalid request data", resultValidator.Errors)
	}

	err := controller.accountService.ChangePassword(ctx, token, nil)
	if err != nil {
		return controller.InternalServerError("Internal server error", err)
	}

	return controller.SuccessResponse(c, nil, "Change password success")
}
func (controller *AccountController) RefreshToken(ctx echo.Context) error {
	// TODO: Implement refresh token
	return nil
}
func (controller *AccountController) Logout(c echo.Context) error {
	ctx := c.Request().Context()

	// Get token from header
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return controller.Unauthorized("Unauthorized", err)
	}

	// Call service to handle logout
	errLogout := controller.accountService.Logout(ctx, token)
	if errLogout != nil {
		return controller.InternalServerError("Failed to logout", errLogout)
	}

	return controller.SuccessResponse(c, nil, "Logout successful")
}
