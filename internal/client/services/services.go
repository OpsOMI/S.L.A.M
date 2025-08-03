package services

import (
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/services/users"
)

type IServices interface {
	Users() users.IUserService
}

type services struct {
	users users.IUserService
}

func NewServices(
	conn net.Conn,
	logger logger.ILogger,
) IServices {
	users := users.NewService(conn)

	return &services{
		users: users,
	}
}

func (s *services) Users() users.IUserService {
	return s.users
}
