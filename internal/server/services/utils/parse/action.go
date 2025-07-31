package parse

import (
	"strconv"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/serviceerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/google/uuid"
)

// ParseOptionalUUID parses a UUID string and allows empty IDs to return an empty UUID.
func (s *service) ParseOptionalUUID(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.UUID{}, nil
	}
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, serviceerrors.BadRequest(
			commons.ErrInvalidID,
		)
	}

	return parsedUUID, nil
}

// ParseRequiredUUID parses a UUID string and returns an error if the ID is invalid or empty.
func (s *service) ParseRequiredUUID(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.UUID{}, serviceerrors.BadRequest(
			commons.ErrInvalidID,
		)
	}
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, serviceerrors.BadRequest(
			commons.ErrInvalidID,
		)
	}

	return parsedUUID, nil
}

func (s *service) Pagination(page, limit string, defaultLimit int) (lim int32, offset int32) {
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	limInt, err := strconv.Atoi(limit)
	if err != nil || limInt <= 0 {
		limInt = defaultLimit
	}

	lim = int32(limInt)
	offset = int32(pageInt-1) * lim

	return
}
