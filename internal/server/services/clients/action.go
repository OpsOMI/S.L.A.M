package clients

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
)

func (s *service) GetByID(
	ctx context.Context,
	id string,
) (*clients.Client, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	domainModel, err := s.repositories.Clients().GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return domainModel, nil
}

func (s *service) GetByClientKey(
	ctx context.Context,
	clientKey string,
) (*clients.Client, error) {
	key, err := s.utils.Parse().ParseRequiredUUID(clientKey)
	if err != nil {
		return nil, err
	}

	domainModel, err := s.repositories.Clients().GetByClientKey(ctx, key)
	if err != nil {
		return nil, err
	}

	return domainModel, nil
}

func (s *service) GetByUserID(
	ctx context.Context,
	userID string,
) (*clients.Clients, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(userID)
	if err != nil {
		return nil, err
	}

	return s.repositories.Clients().GetByUserID(ctx, uid)
}

func (s *service) RevokeByID(
	ctx context.Context,
	id string,
) error {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return err
	}

	return s.repositories.Clients().RevokeByID(ctx, uid)
}

func (s *service) DeleteByID(
	ctx context.Context,
	id string,
) error {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return err
	}

	return s.repositories.Clients().DeleteByID(ctx, uid)
}

func (s *service) IsExistByID(
	ctx context.Context,
	id string,
) (*bool, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	return s.repositories.Clients().IsExistByID(ctx, uid)
}

func (s *service) IsExistByClientKey(
	ctx context.Context,
	clientKey string,
) (*bool, error) {
	key, err := s.utils.Parse().ParseRequiredUUID(clientKey)
	if err != nil {
		return nil, err
	}

	return s.repositories.Clients().IsExistByClientKey(ctx, key)
}

func (s *service) IsRevoked(
	ctx context.Context,
	id string,
) (*bool, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	return s.repositories.Clients().IsRevoked(ctx, uid)
}
