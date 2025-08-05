package domains

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/clients"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/users"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/utils"
)

type IMapper interface {
	Common() commons.ICommonMapper
	Utils() utils.IUtilMapper
	Users() users.IUsersMapper
	Clients() clients.IClientsMapper
	Rooms() rooms.IRoomsMapper
	Messages() messages.IMessagesMapper
}

type mappers struct {
	common   commons.ICommonMapper
	utils    utils.IUtilMapper
	users    users.IUsersMapper
	clients  clients.IClientsMapper
	rooms    rooms.IRoomsMapper
	messages messages.IMessagesMapper
}

func NewMappers() IMapper {
	common := commons.NewMapper()
	utils := utils.NewMapper()
	clients := clients.NewMapper(utils)
	users := users.NewMapper(utils, clients)
	rooms := rooms.NewMapper(utils)
	messages := messages.NewMapper(utils)

	return &mappers{
		common:   common,
		utils:    utils,
		users:    users,
		clients:  clients,
		rooms:    rooms,
		messages: messages,
	}
}

// Common returns the common domain mapper.
func (m *mappers) Common() commons.ICommonMapper {
	return m.common
}

// Utils returns the util domain mapper.
func (m *mappers) Utils() utils.IUtilMapper {
	return m.utils
}

func (m *mappers) Users() users.IUsersMapper {
	return m.users
}

func (m *mappers) Clients() clients.IClientsMapper {
	return m.clients
}

func (m *mappers) Rooms() rooms.IRoomsMapper {
	return m.rooms
}

func (m *mappers) Messages() messages.IMessagesMapper {
	return m.messages
}
