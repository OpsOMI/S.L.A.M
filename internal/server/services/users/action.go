package users

import (
	"context"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/serviceerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/users"
	"github.com/google/uuid"
)

func (s *service) Login(
	ctx context.Context,
	username, password string,
) (*users.User, error) {
	domainModel, err := s.repositories.Users().Login(ctx, username)
	if err != nil {
		if strings.Contains(err.Error(), "not_found") {
			return nil, serviceerrors.BadRequest(users.ErrInvalidCredentials)
		}
		return nil, err
	}

	ok, err := s.packages.Hasher().CompareArgon2(domainModel.Password, password)
	if err != nil {
		return nil, serviceerrors.Internal(users.ErrHashCompareFailed, err)
	}
	if !ok {
		return nil, serviceerrors.BadRequest(users.ErrInvalidCredentials)
	}

	return domainModel, nil
}

func (s *service) GetByID(
	ctx context.Context,
	id string,
) (*users.User, error) {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return nil, err
	}

	domainModel, err := s.repositories.Users().GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return domainModel, nil
}

func (s *service) GetByUsername(
	ctx context.Context,
	username string,
) (*users.User, error) {
	domainModel, err := s.repositories.Users().GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return domainModel, nil
}

func (s *service) GetByNickname(
	ctx context.Context,
	nickname string,
) (*users.User, error) {
	domainModel, err := s.repositories.Users().GetByUsername(ctx, nickname)
	if err != nil {
		return nil, err
	}

	return domainModel, nil
}

func (s *service) GetByPrivateCode(
	ctx context.Context,
	privateCode string,
) (*users.User, error) {
	domainModel, err := s.repositories.Users().GetByUsername(ctx, privateCode)
	if err != nil {
		return nil, err
	}

	return domainModel, nil
}

func (s *service) CreateUser(
	ctx context.Context,
	nickname, privateCode, username, password, role string,
) (*uuid.UUID, error) {
	if ok, err := s.repositories.Users().IsExistByNickname(ctx, nickname); err != nil {
		return nil, err
	} else if *ok {
		return nil, serviceerrors.Conflict(users.ErrNicknameBeingUsed)
	}

	if ok, err := s.repositories.Users().IsExistByUsername(ctx, nickname); err != nil {
		return nil, err
	} else if *ok {
		return nil, serviceerrors.Conflict(users.ErrUsernameBeingUsed)
	}

	domainModel := users.New(nickname, privateCode, username, password, role)
	if err := domainModel.ValidateCreate(); err != nil {
		return nil, err
	}

	hashedPassword, err := s.packages.Hasher().HashArgon2(domainModel.Password)
	if err != nil {
		return nil, serviceerrors.Internal(users.ErrHashingFailed, err)
	}
	domainModel.Password = hashedPassword

	id, err := s.repositories.Users().CreateUser(ctx, domainModel)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *service) ChangeNickname(
	ctx context.Context,
	id, nickname string,
) error {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return err
	}
	model, err := s.repositories.Users().GetByID(ctx, uid)
	if err != nil {
		return err
	}

	model.Nickname = nickname
	if err := model.ValidateForUpdate(); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteByID(
	ctx context.Context,
	id string,
) error {
	uid, err := s.utils.Parse().ParseRequiredUUID(id)
	if err != nil {
		return err
	}

	if err := s.repositories.Users().DeleteByID(ctx, uid); err != nil {
		return err
	}

	return nil
}
