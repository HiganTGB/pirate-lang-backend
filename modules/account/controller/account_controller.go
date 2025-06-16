package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"pirate-lang-go/core/utils"
	"pirate-lang-go/modules/account/dto"
)

func (controller *AccountController) LockUser(c echo.Context) error {
	ctx := c.Request().Context()
	userIDStr := c.Param("userId")
	userID, err := uuid.Parse(userIDStr)
	if err == nil {
		requestData := new(dto.LockUserRequest)
		if err := c.Bind(requestData); err != nil {
			return controller.BadRequest("Invalid request data", err)
		}
		err = controller.accountService.LockUser(ctx, requestData, userID)
		if err != nil {
			return controller.BadRequest("Failed to lock user", err)
		}
		return controller.SuccessResponse(c, nil, "Lock user successfully")
	}
	return controller.BadRequest("Invalid user ID format", err)
}

func (controller *AccountController) UnlockUser(c echo.Context) error {
	ctx := c.Request().Context()
	userIDStr := c.Param("userId")
	userID, err := uuid.Parse(userIDStr)
	if err == nil {
		requestData := new(dto.UnlockUserRequest)
		if err := c.Bind(requestData); err != nil {
			return controller.BadRequest("Invalid request data", err)
		}
		err = controller.accountService.UnlockUser(ctx, requestData, userID)
		if err != nil {
			return controller.BadRequest("Failed to unlock user", err)
		}
		return controller.SuccessResponse(c, nil, "Unlock user successfully")
	}
	return controller.BadRequest("Invalid user ID format", err)
}

func (controller *AccountController) DeleteUser(c echo.Context) error {
	return nil
}

func (controller *AccountController) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	pageNumber := utils.ToNumberWithDefault(c.QueryParam("pageNumber"), 1)
	pageSize := utils.ToNumberWithDefault(c.QueryParam("pageSize"), 20)

	resultGetUsers, err := controller.accountService.GetUsers(ctx, pageNumber, pageSize)
	if err != nil {
		return controller.NotFound(err.Message)
	}

	return controller.SuccessResponse(c, resultGetUsers, "Get users successfully")

}
func (controller *AccountController) GetDetailUser(c echo.Context) error {
	ctx := c.Request().Context()
	userIDStr := c.Param("userId")
	userID, err := uuid.Parse(userIDStr)
	if err == nil {
		profile, err := controller.accountService.GetManagerProfile(ctx, userID)
		if err != nil {
			return controller.BadRequest("Error get profile", err)
		}
		return controller.SuccessResponse(c, profile, "Get Profile successfully")
	}
	return controller.BadRequest("Invalid user ID format", err)
}
