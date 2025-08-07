package api

import (
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/api/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/api/users"
)

type IAPI interface {
	Users() users.IUserModule
	Rooms() rooms.IRoomModule
}

type apis struct {
	users users.IUserModule
	rooms rooms.IRoomModule
}

func NewAPI(
	conn net.Conn,
	logger logger.ILogger,
) IAPI {
	users := users.NewModule(conn)
	rooms := rooms.NewModule(conn)

	return &apis{
		users: users,
		rooms: rooms,
	}
}

func (s *apis) Users() users.IUserModule {
	return s.users
}

func (s *apis) Rooms() rooms.IRoomModule {
	return s.rooms
}
