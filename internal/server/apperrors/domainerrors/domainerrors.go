package domainerrors

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
)

// BadRequest returns a 400 error from the domain layer
func BadRequest(message string) error {
	return apperrors.New(
		commons.ResponseIDJustMessage,
		commons.StatusBadRequest,
		message,
		nil,
		apperrors.SourceDomain,
	)
}

// NotFound returns a 404 error from the domain layer
func NotFound(message string) error {
	return apperrors.New(
		commons.ResponseIDJustMessage,
		commons.StatusNotFound,
		message,
		nil,
		apperrors.SourceDomain,
	)
}
