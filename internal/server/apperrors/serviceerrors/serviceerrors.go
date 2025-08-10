package serviceerrors

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
)

// ServiceUnavailable returns a 503 error from the service layer
func ServiceUnavailable(message string) error {
	return apperrors.New(commons.ResponseIDJustMessage, commons.StatusServiceUnavailable, message, nil, apperrors.SourceService)
}

// BadRequest returns a 400 error from the service layer
func BadRequest(message string) error {
	return apperrors.New(commons.ResponseIDJustMessage, commons.StatusBadRequest, message, nil, apperrors.SourceService)
}

// Unauthorized returns a 401 error from the service layer
func Unauthorized(message string) error {
	return apperrors.New(commons.ResponseIDJustMessage, commons.StatusUnauthorized, message, nil, apperrors.SourceService)
}

// Forbidden returns a 403 error from the service layer
func Forbidden(message string) error {
	return apperrors.New(commons.ResponseIDJustMessage, commons.StatusForbidden, message, nil, apperrors.SourceService)
}

// Internal returns a 500 error from the service layer with an underlying error
func Internal(message string, err error) error {
	return apperrors.New(commons.ResponseIDJustMessage, commons.StatusInternalServerError, message, err, apperrors.SourceService)
}

// Conflict returns a 409 error from the service layer
func Conflict(message string) error {
	return apperrors.New(commons.ResponseIDJustMessage, commons.StatusConflict, message, nil, apperrors.SourceService)
}

// TooManyRequests returns a 429 error from the service layer
func TooManyRequests(message string) error {
	return apperrors.New(commons.ResponseIDJustMessage, commons.StatusTooManyRequests, message, nil, apperrors.SourceService)
}
