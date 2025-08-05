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

func (r *repository) GetByPrivateCode(
	ctx context.Context,
	privateCode string,
) (*users.User, error) {
	dbModel, err := r.queries.GetUserByPrivateCode(ctx, privateCode)
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
) (*uuid.UUID, error) {
	var userID uuid.UUID

	if err := r.txManager.RunInTx(ctx, func(tx *sql.Tx) error {
		qtx := r.queries.WithTx(tx)

		userParams := r.mappers.Users().CreateUser(
			domainModel.Nickname,
			domainModel.PrivateCode,
			domainModel.Username,
			domainModel.Password,
			domainModel.Role,
		)

		uid, err := qtx.CreateUser(ctx, userParams)
		if err != nil {
			return repoerrors.Internal(users.ErrCreateFailed, err)
		}

		clientParams := r.mappers.Clients().CreateClient(uid, uuid.New())
		if _, err = qtx.CreateClient(ctx, clientParams); err != nil {
			return repoerrors.Internal(clients.ErrCreateFailed, err)
		}

		// roomParams := r.mappers.Rooms().CreateParams(uid, domainModel.PrivateCode)
		// if _, err := qtx.CreateRoom(ctx, roomParams); err != nil {
		// 	return repoerrors.Internal(rooms.ErrCreateFailed, err)
		// }
		userID = uid

		return nil
	}); err != nil {
		return nil, err
	}

	return &userID, nil
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

func (r *repository) IsExistByPrivateCode(
	ctx context.Context,
	privateCode string,
) (*bool, error) {
	exists, err := r.queries.IsUserExistByPrivateCode(ctx, privateCode)
	if err != nil {
		return nil, repoerrors.Internal(users.ErrIsExistsFailed, err)
	}
	return &exists, nil
}
