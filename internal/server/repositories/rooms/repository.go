package rooms

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"
	"github.com/OpsOMI/S.L.A.M/pkg/txmanagerpkg"
	"github.com/google/uuid"
)

type IRoomsRepository interface {
	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*rooms.Room, error)

	GetByCode(
		ctx context.Context,
		code string,
	) (*rooms.Room, error)

	GetByCodeAndOwnerID(
		ctx context.Context,
		ownerID uuid.UUID,
		code string,
	) (*rooms.Room, error)

	GetByOwnerID(
		ctx context.Context,
		ownerID uuid.UUID,
		lim, off int32,
	) (*rooms.Rooms, error)

	Create(
		ctx context.Context,
		ownerID uuid.UUID,
		code, hashedPassword string,
	) (*uuid.UUID, error)

	DeleteByID(
		ctx context.Context,
		id uuid.UUID,
	) error

	IsExistByID(
		ctx context.Context,
		id uuid.UUID,
	) (*bool, error)

	IsExistByCode(
		ctx context.Context,
		code string,
	) (*bool, error)

	IsExistByOwnerID(
		ctx context.Context,
		ownerID uuid.UUID,
	) (*bool, error)
}

type repository struct {
	queries   *pgqueries.Queries
	txManager txmanagerpkg.ITXManager
	mappers   domains.IMapper
}

func NewRepository(
	queries *pgqueries.Queries,
	mappers domains.IMapper,
	txManager txmanagerpkg.ITXManager,
) IRoomsRepository {
	return &repository{
		txManager: txManager,
		queries:   queries,
		mappers:   mappers,
	}
}
