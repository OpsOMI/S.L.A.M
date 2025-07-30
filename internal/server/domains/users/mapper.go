package users

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/utils"
	"github.com/google/uuid"
)

type IUsersMapper interface {
	One(
		dbModel *pgqueries.User,
	) *User

	Many(
		dbModels []pgqueries.User,
		count int64,
	) *Users

	CreateUser(
		username, nickname, password string,
	) pgqueries.CreateUserParams

	ChangeNickname(
		id uuid.UUID,
		nickname string,
	) pgqueries.ChangeNicknameParams
}

type mapper struct {
	utils utils.IUtilMapper
}

// NewMapper creates a new instance of userAuthMapper.
func NewMapper(
	utilsMapper utils.IUtilMapper,
) *mapper {
	return &mapper{
		utils: utilsMapper,
	}
}

// One maps a DB model (TUsersAuth) to a domain model (UserAuth).
func (m *mapper) One(
	dbModel *pgqueries.User,
) *User {
	if dbModel == nil {
		return nil
	}

	return &User{
		ID:          dbModel.ID,
		Username:    dbModel.Username,
		Password:    dbModel.Password,
		Nickname:    dbModel.Nickname,
		PrivateCode: dbModel.PrivateCode,
		CreatedAt:   dbModel.CreatedAt.Time,
	}
}

// Many maps a list of DB models (TUsersAuth) to domain models with total count.
func (m *mapper) Many(
	dbModels []pgqueries.User,
	count int64,
) *Users {
	var appModels []User
	for _, dbModel := range dbModels {
		appModels = append(appModels, *m.One(&dbModel))
	}

	return &Users{
		Items:      appModels,
		TotalCount: count,
	}
}

func (m *mapper) CreateUser(
	username, nickname, password string,
) pgqueries.CreateUserParams {
	return pgqueries.CreateUserParams{
		Username: username,
		Nickname: nickname,
		Password: password,
	}
}

func (m *mapper) ChangeNickname(
	id uuid.UUID,
	nickname string,
) pgqueries.ChangeNicknameParams {
	return pgqueries.ChangeNicknameParams{
		ID:       id,
		Nickname: *m.utils.FromStrToPtrStr(nickname),
	}
}
