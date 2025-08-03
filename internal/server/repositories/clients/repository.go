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
