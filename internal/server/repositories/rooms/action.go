package rooms

import (
	"context"
	"database/sql"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/repoerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"
	"github.com/google/uuid"
)

func (r *repository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*rooms.Room, error) {
	dbModel, err := r.queries.GetRoomByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.NotFound(rooms.ErrNotFound)
		}
		return nil, repoerrors.Internal(rooms.ErrFetchFailed, err)
	}
	return r.mappers.Rooms().One(&dbModel), nil
}

func (r *repository) GetByCode(
	ctx context.Context,
	code string,
) (*rooms.Room, error) {
	dbModel, err := r.queries.GetRoomByCode(ctx, code)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repoerrors.NotFound(rooms.ErrNotFound)
		}
		return nil, repoerrors.Internal(rooms.ErrFetchFailed, err)
	}
	return r.mappers.Rooms().One(&dbModel), nil
}

func (r *repository) GetByOwnerID(
	ctx context.Context,
	ownerID uuid.UUID,
	lim, off int32,
) (*rooms.Rooms, error) {
	params := r.mappers.Rooms().GetByOwnerID(ownerID, lim, off)
	dbModels, err := r.queries.GetRoomsByOwnerID(ctx, params)
	if err != nil {
		return nil, repoerrors.Internal(rooms.ErrFetchFailed, err)
	}

	count, err := r.queries.CountRoomsByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, repoerrors.Internal(rooms.ErrCountFailed, err)
	}

	return r.mappers.Rooms().Many(dbModels, count), nil
}

func (r *repository) Create(
	ctx context.Context,
	ownerID uuid.UUID,
	code, hashedPassword string,
) (*uuid.UUID, error) {
	params := r.mappers.Rooms().CreateParams(ownerID, code, hashedPassword)
	id, err := r.queries.CreateRoom(ctx, params)
	if err != nil {
		return nil, repoerrors.Internal(rooms.ErrCreateFailed, err)
	}

	return &id, nil
}

func (r *repository) DeleteByID(
	ctx context.Context,
	id uuid.UUID,
) error {
	err := r.queries.DeleteRoomByID(ctx, id)
	if err != nil {
		return repoerrors.Internal(rooms.ErrDeleteFailed, err)
	}
	return nil
}

func (r *repository) IsExistByID(
	ctx context.Context,
	id uuid.UUID,
) (*bool, error) {
	exists, err := r.queries.IsRoomExistByID(ctx, id)
	if err != nil {
		return nil, repoerrors.Internal(rooms.ErrIsExistsFailed, err)
	}
	return &exists, nil
}

func (r *repository) IsExistByCode(
	ctx context.Context,
	code string,
) (*bool, error) {
	exists, err := r.queries.IsRoomExistByCode(ctx, code)
	if err != nil {
		return nil, repoerrors.Internal(rooms.ErrIsExistsFailed, err)
	}
	return &exists, nil
}

func (r *repository) IsExistByOwnerID(
	ctx context.Context,
	ownerID uuid.UUID,
) (*bool, error) {
	exists, err := r.queries.IsRoomExistByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, repoerrors.Internal(rooms.ErrIsExistsFailed, err)
	}
	return &exists, nil
}
