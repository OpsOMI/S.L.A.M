package clients

import (
	"context"
	"database/sql"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/repoerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"github.com/google/uuid"
)

func (r *repository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*clients.Client, error) {
	dbModel, err := r.queries.GetClientByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.NotFound(clients.ErrNotFound)
		}
		return nil, repoerrors.Internal(clients.ErrFetchFailed, err)
	}
	return r.mappers.Clients().One(&dbModel), nil
}

func (r *repository) GetByClientKey(
	ctx context.Context,
	clientKey uuid.UUID,
) (*clients.Client, error) {
	dbModel, err := r.queries.GetClientByClientKey(ctx, clientKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.NotFound(clients.ErrNotFound)
		}
		return nil, repoerrors.Internal(clients.ErrFetchFailed, err)
	}
	return r.mappers.Clients().One(&dbModel), nil
}

func (r *repository) GetByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*clients.Clients, error) {
	dbModels, err := r.queries.GetClientsByUserID(ctx, userID)
	if err != nil {
		return nil, repoerrors.Internal(clients.ErrFetchFailed, err)
	}
	count, err := r.queries.CountClientsByUserID(ctx, userID)
	if err != nil {
		return nil, repoerrors.Internal(clients.ErrCountFailed, err)
	}

	return r.mappers.Clients().Many(dbModels, count), nil
}

func (r *repository) RevokeByID(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := r.queries.RevokeClient(ctx, id)
	if err != nil {
		return repoerrors.Internal(clients.ErrUpdateFailed, err)
	}
	return nil
}

func (r *repository) DeleteByID(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := r.queries.DeleteClient(ctx, id)
	if err != nil {
		return repoerrors.Internal(clients.ErrDeleteFailed, err)
	}
	return nil
}

func (r *repository) IsExistByID(
	ctx context.Context,
	id uuid.UUID,
) (*bool, error) {
	exists, err := r.queries.IsClientExistByID(ctx, id)
	if err != nil {
		return nil, repoerrors.Internal(clients.ErrIsExistsFailed, err)
	}
	return &exists, nil
}

func (r *repository) IsExistByClientKey(
	ctx context.Context,
	clientKey uuid.UUID,
) (*bool, error) {
	exists, err := r.queries.IsClientExistByClientKey(ctx, clientKey)
	if err != nil {
		return nil, repoerrors.Internal(clients.ErrIsExistsFailed, err)
	}
	return &exists, nil
}

func (r *repository) IsRevoked(
	ctx context.Context,
	id uuid.UUID,
) (*bool, error) {
	exists, err := r.queries.IsClientRevoked(ctx, id)
	if err != nil {
		return nil, repoerrors.Internal(clients.ErrIsExistsFailed, err)
	}
	return &exists, nil
}
