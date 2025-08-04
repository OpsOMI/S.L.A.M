package utils

import (
	"encoding/json"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/serviceerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
)

func ParseJSON[T any](data json.RawMessage, target *T) error {
	if err := json.Unmarshal(data, target); err != nil {
		return serviceerrors.BadRequest(commons.ErrParseFailed)
	}
	return nil
}
