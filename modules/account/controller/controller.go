package controller

import "github.com/labstack/echo/v4"

func (controller *AccountController) HelloWorld(c echo.Context) error {
	data := map[string]string{
		"message": "Hello, World!",
	}
	return controller.SuccessResponse(c, data, "Hello World from API")
}
