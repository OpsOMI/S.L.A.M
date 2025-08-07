package apperrors

import "fmt"

type AppError struct {
	Code    string // "Error", "Notification",
	Message string
}

const (
	CodeError        = "Error"
	CodeNotification = "Notification"
)

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewError(message string) *AppError {
	return &AppError{
		Code:    "Error",
		Message: message,
	}
}

func NewNotification(message string) *AppError {
	return &AppError{
		Code:    "Notification",
		Message: message,
	}
}

func New(code, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
