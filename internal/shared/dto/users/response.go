package users

import "github.com/google/uuid"

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
