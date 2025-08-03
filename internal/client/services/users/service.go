package users

import (
	"net"
)

type IUserService interface {
}

type service struct {
	conn net.Conn
}

func NewService(
	conn net.Conn,
) IUserService {
	return &service{
		conn: conn,
	}
}
