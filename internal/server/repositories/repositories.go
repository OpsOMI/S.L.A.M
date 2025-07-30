package repositories

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories/users"
	"github.com/OpsOMI/S.L.A.M/pkg/txmanagerpkg"
)

type IRepositories interface {
	Users() users.IUserRepository
}

type repositories struct {
	users users.IUserRepository
}

func NewRepositories(
	queries *pgqueries.Queries,
	mappers domains.IMapper,
	txManager txmanagerpkg.ITXManager,
) IRepositories {
	user := users.NewRepository(queries, mappers, txManager)

	return &repositories{
		users: user,
	}
}

func (r *repositories) Users() users.IUserRepository {
	return r.users
}
