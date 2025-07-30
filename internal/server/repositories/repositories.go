package repositories

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains"
	"github.com/OpsOMI/S.L.A.M/pkg/txmanagerpkg"
)

type IRepositories interface {
}

type repositories struct {
}

func NewRepositories(
	q *pgqueries.Queries,
	txManager txmanagerpkg.ITXManager,
	mappers domains.IMapper,
) IRepositories {
	return &repositories{}
}
