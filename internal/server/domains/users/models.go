package users

import (
	"strings"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/domainerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Username    string
	Password    string
	Nickname    string
	PrivateCode string
	Role        string
	Clients     *clients.Client
	CreatedAt   time.Time
}

type Users struct {
	Items      []User
	TotalCount int64
}

// New constructs a UserAuth instance for creation.
func New(
	nickname, privateCode, username, password, role string,
) User {
	return User{
		Nickname:    strings.TrimSpace(nickname),
		Username:    strings.TrimSpace(username),
		Password:    strings.TrimSpace(password),
		PrivateCode: strings.TrimSpace(privateCode),
		Role:        strings.TrimSpace(role),
	}
}

// ValidateCreate checks domain rules for user registration.
func (u *User) ValidateCreate() error {
	if u.Nickname == "" {
		return domainerrors.BadRequest(ErrNicknameRequired)
	}
	if len(u.Nickname) > 12 {
		return domainerrors.BadRequest(ErrNicknameTooLong)
	}
	if len(u.Nickname) < 3 {
		return domainerrors.BadRequest(ErrNicknameTooShort)
	}

	if u.Username == "" {
		return domainerrors.BadRequest(ErrUsernameRequired)
	}
	if len(u.Username) > 30 {
		return domainerrors.BadRequest(ErrUsernameTooLong)
	}
	if len(u.Username) < 3 {
		return domainerrors.BadRequest(ErrUsernameTooShort)
	}

	if u.Password == "" {
		return domainerrors.BadRequest(ErrPasswordRequired)
	}

	return nil
}

// ValidateForUpdate checks update-specific domain rules.
func (u *User) ValidateForUpdate() error {
	if u.Nickname != "" {
		if len(u.Nickname) > 12 {
			return domainerrors.BadRequest(ErrNicknameTooLong)
		}
		if len(u.Nickname) < 3 {
			return domainerrors.BadRequest(ErrNicknameTooShort)
		}
	}

	return nil
}
