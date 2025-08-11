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
	clientKey, username, password, mode string,
) (*users.User, error) {
	user, err := s.repositories.Users().GetByUsername(ctx, username)
	if err != nil {
		if strings.Contains(err.Error(), "not_found") {
			return nil, serviceerrors.BadRequest(users.ErrInvalidCredentials)
		}
		return nil, err
	}

	if mode == "prod" {
		if user.Role == "banned" {
			return nil, serviceerrors.Forbidden(users.ErrUserBanned)
		}

		client, err := s.clients.GetByClientKey(ctx, clientKey)
		if err != nil {
			// If this condition is met, it likely means the user is trying to use a client executable
			// that is not registered in our database—probably due to tampering or reverse engineering.
			// In other words, they are attempting unauthorized access with a modified or stolen client.
			if strings.Contains(err.Error(), "not_found") {
				if err := s.BanUser(ctx, user.ID.String()); err != nil {
					return nil, err
				}
				return nil, serviceerrors.Forbidden(users.ErrUnregisteredClient)
			}
			return nil, err
		}

		if client.IsRevoked() {
			return nil, serviceerrors.Forbidden(users.ErrClientIsRevoked)
		}

		// If the below if is entered, the user will have stolen or using someone else's flash.
		if client.UserID != user.ID {
			if err := s.clients.RevokeByID(ctx, client.ID.String()); err != nil {
				return nil, err
			}
			return nil, serviceerrors.Forbidden(users.ErrInvalidOrStolenClient)
		}
	}

	ok, err := s.packages.Hasher().CompareArgon2(user.Password, password)
	if err != nil {
		return nil, serviceerrors.Internal(users.ErrHashCompareFailed, err)
	}
	if !ok {
		return nil, serviceerrors.BadRequest(users.ErrInvalidCredentials)
	}

	return user, nil
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
	nickname, username, password, role string,
) (*uuid.UUID, *string, error) {
	if ok, err := s.repositories.Users().IsExistByNickname(ctx, nickname); err != nil {
		return nil, nil, err
	} else if *ok {
		return nil, nil, serviceerrors.Conflict(users.ErrNicknameBeingUsed)
	}

	if ok, err := s.repositories.Users().IsExistByUsername(ctx, nickname); err != nil {
		return nil, nil, err
	} else if ok {
		return nil, nil, serviceerrors.Conflict(users.ErrUsernameBeingUsed)
	}

	privateCode, err := s.packages.Hasher().Generate6CharPrivateCode()
	if err != nil {
		return nil, nil, serviceerrors.Conflict(users.ErrCreatingPrivateCodeFailed)
	}

	domainModel := users.New(nickname, privateCode, username, password, role)
	if err := domainModel.ValidateCreate(); err != nil {
		return nil, nil, err
	}

	hashedPassword, err := s.packages.Hasher().HashArgon2(domainModel.Password)
	if err != nil {
		return nil, nil, serviceerrors.Internal(users.ErrHashingFailed, err)
	}
	domainModel.Password = hashedPassword

	id, clientID, err := s.repositories.Users().CreateUser(ctx, domainModel)
	if err != nil {
		return nil, nil, err
	}

	return id, clientID, nil
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

func (s *service) BanUser(
	ctx context.Context,
	id string,
) error {
	model, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repositories.Users().BanUser(ctx, model.ID)
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

func (s *service) IsExistsByUsername(
	ctx context.Context,
	username string,
) (bool, error) {
	return s.repositories.Users().IsExistByUsername(ctx, username)
}
