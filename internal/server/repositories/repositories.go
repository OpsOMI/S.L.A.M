package repositories

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories/clients"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories/messages"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories/users"
	"github.com/OpsOMI/S.L.A.M/pkg/txmanagerpkg"
)

type IRepositories interface {
	Users() users.IUserRepository
	Clients() clients.IClientsRepository
	Rooms() rooms.IRoomsRepository
	Messages() messages.IMessagesRepository
}

type repositories struct {
	users    users.IUserRepository
	clients  clients.IClientsRepository
	rooms    rooms.IRoomsRepository
	messages messages.IMessagesRepository
}

func NewRepositories(
	queries *pgqueries.Queries,
	mappers domains.IMapper,
	txManager txmanagerpkg.ITXManager,
) IRepositories {
	user := users.NewRepository(queries, mappers, txManager)
	clients := clients.NewRepository(queries, mappers, txManager)
	rooms := rooms.NewRepository(queries, mappers, txManager)
	messages := messages.NewRepository(queries, mappers, txManager)

	return &repositories{
		users:    user,
		clients:  clients,
		rooms:    rooms,
		messages: messages,
	}
}

func (r *repositories) Users() users.IUserRepository {
	return r.users
}

func (r *repositories) Clients() clients.IClientsRepository {
	return r.clients
}

func (r *repositories) Rooms() rooms.IRoomsRepository {
	return r.rooms
}

func (r *repositories) Messages() messages.IMessagesRepository {
	return r.messages
}
