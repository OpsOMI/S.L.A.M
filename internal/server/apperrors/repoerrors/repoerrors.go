package repoerrors

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
)

// Internal returns a 500 error from the repository layer
func Internal(message string, err error) error {
	return apperrors.New(
		commons.StatusInternalServerError,
		message,
		err,
		apperrors.SourceRepo,
	)
}

// BadRequest returns a 400 error from the repository layer
func BadRequest(message string) error {
	return apperrors.New(
		commons.StatusBadRequest,
		message,
		nil,
		apperrors.SourceRepo,
	)
}

// NotFound returns a 404 error from the repository layer
func NotFound(message string) error {
	return apperrors.New(
		commons.StatusNotFound,
		message,
		nil,
		apperrors.SourceRepo,
	)
}
