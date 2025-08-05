package users

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/utils"
	"github.com/google/uuid"
)

type IUsersMapper interface {
	One(
		dbModel *pgqueries.User,
	) *User

	OneWithClient(
		dbModel *pgqueries.UserLoginRow,
	) *User

	OneWithPrivateCode(
		dbModel *pgqueries.GetUserFullInfoRow,
	) *User

	Many(
		dbModels []pgqueries.User,
		count int64,
	) *Users

	CreateUser(
		nickname, privateCode, username, password, role string,
	) pgqueries.CreateUserParams

	ChangeNickname(
		id uuid.UUID,
		nickname string,
	) pgqueries.ChangeNicknameParams
}

type mapper struct {
	utils   utils.IUtilMapper
	clients clients.IClientsMapper
}

func NewMapper(
	utils utils.IUtilMapper,
	clients clients.IClientsMapper,
) *mapper {
	return &mapper{
		utils:   utils,
		clients: clients,
	}
}

func (m *mapper) One(
	dbModel *pgqueries.User,
) *User {
	if dbModel == nil {
		return nil
	}

	return &User{
		ID:          dbModel.ID,
		Role:        dbModel.Role,
		Username:    dbModel.Username,
		Password:    dbModel.Password,
		Nickname:    dbModel.Nickname,
		PrivateCode: dbModel.PrivateCode,
		CreatedAt:   dbModel.CreatedAt.Time,
	}
}

func (m *mapper) OneWithClient(
	dbModel *pgqueries.UserLoginRow,
) *User {
	if dbModel == nil {
		return nil
	}

	userModel := m.One(&dbModel.User)
	clientModel := m.clients.One(&dbModel.Client)

	userModel.Clients = clientModel

	return userModel
}

func (m *mapper) OneWithPrivateCode(
	dbModel *pgqueries.GetUserFullInfoRow,
) *User {
	if dbModel == nil {
		return nil
	}

	userModel := m.One(&dbModel.User)
	clientModel := m.clients.One(&dbModel.Client)

	userModel.Clients = clientModel

	return userModel
}

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
	nickname, privateCode, username, password, role string,
) pgqueries.CreateUserParams {
	return pgqueries.CreateUserParams{
		Nickname:    nickname,
		PrivateCode: privateCode,
		Username:    username,
		Password:    password,
		Role:        role,
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
