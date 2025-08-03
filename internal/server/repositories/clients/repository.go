package clients

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"github.com/OpsOMI/S.L.A.M/pkg/txmanagerpkg"
	"github.com/google/uuid"
)

type IClientsRepository interface {
	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*clients.Client, error)

	GetByClientKey(
		ctx context.Context,
		clientKey uuid.UUID,
	) (*clients.Client, error)

	GetByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) (*clients.Clients, error)

	RevokeByID(
		ctx context.Context,
		id uuid.UUID,
	) error

	DeleteByID(
		ctx context.Context,
		id uuid.UUID,
	) error

	IsExistByID(
		ctx context.Context,
		id uuid.UUID,
	) (*bool, error)

	IsExistByClientKey(
		ctx context.Context,
		clientKey uuid.UUID,
	) (*bool, error)

	IsRevoked(
		ctx context.Context,
		id uuid.UUID,
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
) IClientsRepository {
	return &repository{
		txManager: txManager,
		queries:   queries,
		mappers:   mappers,
	}
}
