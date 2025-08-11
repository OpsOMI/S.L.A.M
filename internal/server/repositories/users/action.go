package users

import (
	"context"
	"database/sql"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/repoerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/users"
	"github.com/google/uuid"
)

func (r *repository) Login(
	ctx context.Context,
	username string,
) (*users.User, error) {
	dbModel, err := r.queries.UserLogin(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.NotFound(users.ErrNotFound)
		}
		return nil, repoerrors.Internal(users.ErrFetchFailed, err)
	}
	return r.mappers.Users().OneWithClient(&dbModel), nil
}

func (r *repository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*users.User, error) {
	dbModel, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.NotFound(users.ErrNotFound)
		}
		return nil, repoerrors.Internal(users.ErrFetchFailed, err)
	}
	return r.mappers.Users().One(&dbModel), nil
}

func (r *repository) GetByUsername(
	ctx context.Context,
	username string,
) (*users.User, error) {
	dbModel, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.NotFound(users.ErrNotFound)
		}
		return nil, repoerrors.Internal(users.ErrFetchFailed, err)
	}
	return r.mappers.Users().One(&dbModel), nil
}

func (r *repository) GetByNickname(
	ctx context.Context,
	nickname string,
) (*users.User, error) {
	dbModel, err := r.queries.GetUserByNickname(ctx, nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.NotFound(users.ErrNotFound)
		}
		return nil, repoerrors.Internal(users.ErrFetchFailed, err)
	}
	return r.mappers.Users().One(&dbModel), nil
}

func (r *repository) CreateUser(
	ctx context.Context,
	domainModel users.User,
) (*uuid.UUID, *string, error) {
	var userID uuid.UUID
	var cKey string

	if err := r.txManager.RunInTx(ctx, func(tx *sql.Tx) error {
		qtx := r.queries.WithTx(tx)

		userParams := r.mappers.Users().CreateUser(
			domainModel.Nickname,
			domainModel.Username,
			domainModel.Password,
			domainModel.Role,
		)

		uid, err := qtx.CreateUser(ctx, userParams)
		if err != nil {
			return repoerrors.Internal(users.ErrCreateFailed, err)
		}

		clientKey := uuid.New()
		clientParams := r.mappers.Clients().CreateClient(uid, clientKey)
		if _, err = qtx.CreateClient(ctx, clientParams); err != nil {
			return repoerrors.Internal(clients.ErrCreateFailed, err)
		}
		cKey = clientKey.String()
		userID = uid

		return nil
	}); err != nil {
		return nil, nil, err
	}

	return &userID, &cKey, nil
}

func (r *repository) ChangeNickname(
	ctx context.Context,
	id uuid.UUID,
	nickname string,
) error {
	params := r.mappers.Users().ChangeNickname(id, nickname)
	if err := r.queries.ChangeNickname(ctx, params); err != nil {
		return repoerrors.Internal(users.ErrUpdateFailed, err)
	}
	return nil
}

func (r *repository) BanUser(
	ctx context.Context,
	id uuid.UUID,
) error {
	if err := r.queries.BanUser(ctx, id); err != nil {
		return repoerrors.Internal(users.ErrUpdateFailed, err)
	}
	return nil
}

func (r *repository) DeleteByID(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := r.queries.DeleteUser(ctx, id)
	if err != nil {
		return repoerrors.Internal(users.ErrDeleteFailed, err)
	}
	return nil
}

func (r *repository) IsExistByID(
	ctx context.Context,
	id uuid.UUID,
) (*bool, error) {
	exists, err := r.queries.IsUserExistByID(ctx, id)
	if err != nil {
		return nil, repoerrors.Internal(users.ErrIsExistsFailed, err)
	}
	return &exists, nil
}

func (r *repository) IsExistByUsername(
	ctx context.Context,
	username string,
) (bool, error) {
	exists, err := r.queries.IsUserExistByUsername(ctx, username)
	if err != nil {
		return false, repoerrors.Internal(users.ErrIsExistsFailed, err)
	}
	return exists, nil
}

func (r *repository) IsExistByNickname(
	ctx context.Context,
	nickname string,
) (*bool, error) {
	exists, err := r.queries.IsUserExistByNickname(ctx, nickname)
	if err != nil {
		return nil, repoerrors.Internal(users.ErrIsExistsFailed, err)
	}
	return &exists, nil
}
