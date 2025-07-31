package parse

import (
	"github.com/google/uuid"
)

type IParseService interface {
	// Pagination parsing
	Pagination(page, limit string, defaultLimit int) (lim int32, offset int32)

	// UUID parsing
	ParseOptionalUUID(id string) (uuid.UUID, error)
	ParseRequiredUUID(id string) (uuid.UUID, error)
}

type service struct {
}

func NewService() IParseService {
	return &service{}
}
