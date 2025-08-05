package rooms

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"
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

func (s *service) IsIsRoomOrDirectChat(
	ctx context.Context,
	roomOrUserCode string,
) (*clients.Client, bool, error) {
	room, err := s.GetByCode(ctx, roomOrUserCode)
	if err != nil && !strings.Contains(err.Error(), "not_found") {
		return nil, false, err
	}
	if room != nil {
		return nil, true, nil
	}

	reciever, err := s.clients.GetByClientKey(ctx, roomOrUserCode)
	if err != nil {
		return nil, false, fmt.Errorf("no user or room found with code: %s", roomOrUserCode)
	}

	return reciever, false, nil
}
