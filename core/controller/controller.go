package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response types
type (
	SuccessResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}

	ErrorResponse struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details any    `json:"details,omitempty"`
	}

	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ValidationResponse struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Errors  []ValidationError `json:"errors"`
	}
)

// Response handler interface and implementation
type BaseController interface {
	BadRequest(c *gin.Context, message string, details ...any)
	InternalServerError(c *gin.Context, message string, details ...any)
	NotFound(c *gin.Context, message string, details ...any)
	Unauthorized(c *gin.Context, message string, details ...any)
	Forbidden(c *gin.Context, message string, details ...any)
	SuccessResponse(c *gin.Context, data any, message string)
	ErrorResponse(c *gin.Context, httpStatus int, err error)
	ValidationFailedResponse(c *gin.Context, message string, errors []ValidationError) // Thêm hàm này
}

type responseHandler struct{}

func NewBaseController() BaseController {
	return &responseHandler{}
}

// Success response functions
func NewSuccessResponse(data any, message string) *SuccessResponse {
	return &SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// Error response functions
func NewErrorResponse(code int, message string, details ...interface{}) *ErrorResponse {
	err := &ErrorResponse{
		Status:  "error", // Added status field for consistency
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err

}

// Validation functions
func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

// HTTP Error handlers
func (h *responseHandler) BadRequest(c *gin.Context, message string, details ...any) {
	errResp := NewErrorResponse(http.StatusBadRequest, message, details...)
	c.JSON(http.StatusBadRequest, errResp)
	c.Abort()
}

func (h *responseHandler) InternalServerError(c *gin.Context, message string, details ...any) {
	errResp := NewErrorResponse(http.StatusInternalServerError, message, details...)
	c.JSON(http.StatusInternalServerError, errResp)
	c.Abort()
}

func (h *responseHandler) NotFound(c *gin.Context, message string, details ...any) {
	errResp := NewErrorResponse(http.StatusNotFound, message, details...)
	c.JSON(http.StatusNotFound, errResp)
	c.Abort()
}

func (h *responseHandler) Unauthorized(c *gin.Context, message string, details ...any) {
	errResp := NewErrorResponse(http.StatusUnauthorized, message, details...)
	c.JSON(http.StatusUnauthorized, errResp)
	c.Abort()
}

func (h *responseHandler) Forbidden(c *gin.Context, message string, details ...any) {
	errResp := NewErrorResponse(http.StatusForbidden, message, details...)
	c.JSON(http.StatusForbidden, errResp)
	c.Abort()
}

func (h *responseHandler) SuccessResponse(c *gin.Context, data any, message string) {
	c.JSON(http.StatusOK, NewSuccessResponse(data, message))
}

func (h *responseHandler) ErrorResponse(c *gin.Context, httpStatus int, err error) {
	errResp := NewErrorResponse(httpStatus, err.Error())
	c.JSON(httpStatus, errResp)
	c.Abort()
}

// ValidationFailedResponse handles validation errors.
func (h *responseHandler) ValidationFailedResponse(c *gin.Context, message string, errors []ValidationError) {
	response := ValidationResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	}
	c.JSON(http.StatusBadRequest, response) //  400 for validation errors
	c.Abort()
}
