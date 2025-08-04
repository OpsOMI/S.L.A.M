package api

import (
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/api/users"
)

type IAPI interface {
	Users() users.IUserModule
}

type apis struct {
	users users.IUserModule
}

func NewAPI(
	conn net.Conn,
	logger logger.ILogger,
) IAPI {
	users := users.NewModule(conn)

	return &apis{
		users: users,
	}
}

func (s *apis) Users() users.IUserModule {
	return s.users
}
