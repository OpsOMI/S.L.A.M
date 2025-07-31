package apperrors

import (
	"fmt"
	"time"
)

// ErrorSource represents the source layer where the error originated.
type ErrorSource string

const (
	SourceRepo    ErrorSource = "repository" // Error originating from the repository layer
	SourceService ErrorSource = "service"    // Error originating from the service layer
	SourceDomain  ErrorSource = "domain"     // Error originating from the domain layer
)

// AppError represents an application-specific error with additional metadata.
type AppError struct {
	Code      string //  Frobidden, OK,
	Message   string
	Source    ErrorSource
	Cause     error
	Timestamp time.Time
}

// Error formats the error message, including the source, status code, message, and underlying cause if any.
func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s %s: %v", e.Source, e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s %s", e.Source, e.Code, e.Message)
}

// Unwrap returns the underlying error, supporting Go 1.13+ error unwrapping.
func (e *AppError) Unwrap() error {
	return e.Cause
}

// New creates a new AppError with the given status code, message, optional cause, and source.
func New(code string, message string, cause error, source ErrorSource) error {
	return &AppError{
		Code:      code,
		Message:   message,
		Cause:     cause,
		Source:    source,
		Timestamp: time.Now(),
	}
}
