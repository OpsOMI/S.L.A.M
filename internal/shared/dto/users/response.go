package users

import (
	"github.com/OpsOMI/S.L.A.M/internal/shared/store"
	"github.com/google/uuid"
)

type MeResp struct {
	UserID    uuid.UUID `json:"userID"`
	ClientKey uuid.UUID `json:"clientKey"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
}

func ToMeResp(
	info *store.TokenInfo,
) MeResp {
	return MeResp{
		ClientKey: info.ClientID,
		UserID:    info.UserID,
		Username:  info.Username,
		Nickname:  info.Nickname,
	}
}

type LoginResp struct {
	Token string
}

func ToLoginResponse(
	token string,
) LoginResp {
	return LoginResp{
		Token: token,
	}
}

type RegisterResp struct {
	ID uuid.UUID
}

func ToRegisterResponse(
	id uuid.UUID,
) RegisterResp {
	return RegisterResp{
		ID: id,
	}
}
