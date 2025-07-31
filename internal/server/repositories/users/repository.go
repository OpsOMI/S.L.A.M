package users

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/users"
	"github.com/OpsOMI/S.L.A.M/pkg/txmanagerpkg"
	"github.com/google/uuid"
)

type IUserRepository interface {
	GetByID(
		ctx context.Context,
		id uuid.UUID,
	) (*users.User, error)

	GetByUsername(
		ctx context.Context,
		username string,
	) (*users.User, error)

	GetByNickname(
		ctx context.Context,
		nickname string,
	) (*users.User, error)

	GetByPrivateCode(
		ctx context.Context,
		privateCode string,
	) (*users.User, error)

	CreateUser(
		ctx context.Context,
		domainModel users.User,
	) (*uuid.UUID, error)

	ChangeNickname(
		ctx context.Context,
		id uuid.UUID,
		nickname string,
	) error

	DeleteByID(
		ctx context.Context,
		id uuid.UUID,
	) error

	IsExistByID(
		ctx context.Context,
		id uuid.UUID,
	) (*bool, error)

	IsExistByUsername(
		ctx context.Context,
		username string,
	) (*bool, error)

	IsExistByNickname(
		ctx context.Context,
		nickname string,
	) (*bool, error)

	IsExistByPrivateCode(
		ctx context.Context,
		privateCode string,
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
) IUserRepository {
	return &repository{
		txManager: txManager,
		queries:   queries,
		mappers:   mappers,
	}
}
