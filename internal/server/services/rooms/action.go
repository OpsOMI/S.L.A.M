package rooms

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/serviceerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"
	"github.com/google/uuid"
)

func (s *service) GetByID(
	ctx context.Context,
	id string,
) (*rooms.Room, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	return s.repositories.Rooms().GetByID(ctx, uid)
}

func (s *service) GetByCode(
	ctx context.Context,
	code string,
) (*rooms.Room, error) {
	return s.repositories.Rooms().GetByCode(ctx, code)
}

func (s *service) GetByOwnerID(
	ctx context.Context,
	ownerID string,
) (*rooms.Rooms, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(ownerID)
	if err != nil {
		return nil, err
	}

	return s.repositories.Rooms().GetByOwnerID(ctx, uid)
}

func (s *service) Create(
	ctx context.Context,
	ownerID, password string,
) (*uuid.UUID, error) {
	owner, err := s.users.GetByID(ctx, ownerID)
	if err != nil {
		return nil, err
	}

	var hashedPassword string
	if password != "" {
		hashedPassword, err = s.packages.Hasher().HashArgon2(password)
		if err != nil {
			return nil, serviceerrors.Internal(rooms.ErrPasswordHashFailed, err)
		}
	}

	roomCode, err := s.packages.Hasher().Generate6CharPrivateCode()
	if err != nil {
		return nil, serviceerrors.Internal(rooms.ErrCodeGenerateFailed, err)
	}

	id, err := s.repositories.Rooms().Create(ctx, owner.ID, roomCode, hashedPassword)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *service) DeleteByID(
	ctx context.Context,
	id string,
) error {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return err
	}

	return s.repositories.Rooms().DeleteByID(ctx, uid)
}

func (s *service) IsExistByID(
	ctx context.Context,
	id string,
) (*bool, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	return s.repositories.Rooms().IsExistByID(ctx, uid)
}

func (s *service) IsExistByCode(
	ctx context.Context,
	code string,
) (*bool, error) {
	return s.repositories.Rooms().IsExistByCode(ctx, code)
}

func (s *service) IsExistByOwnerID(
	ctx context.Context,
	ownerID string,
) (*bool, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(ownerID)
	if err != nil {
		return nil, err
	}

	return s.repositories.Rooms().IsExistByOwnerID(ctx, uid)
}

func (s *service) JoinRoom(
	ctx context.Context,
	code, password string,
) (*rooms.Room, error) {
	room, err := s.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	if room.Password != "" {
		ok, err := s.packages.Hasher().CompareArgon2(room.Password, password)
		if err != nil {
			return nil, serviceerrors.Internal(rooms.ErrPasswordHashFailed, err)
		}
		if !ok {
			return nil, serviceerrors.Internal(rooms.ErrInvalidCredentials, err)
		}

		return room, nil
	}

	return room, nil
}
