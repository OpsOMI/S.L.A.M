package users

import (
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type IUserModule interface {
	Login(
		req *request.ClientRequest,
		clientKey string,
	) error

	Register(
		req *request.ClientRequest,
	) error

	Join(
		req *request.ClientRequest,
		roomCode string,
	) error

	SendMessage(
		req *request.ClientRequest,
		content string,
	) error
}

type module struct {
	conn net.Conn
}

func NewModule(
	conn net.Conn,
) IUserModule {
	return &module{
		conn: conn,
	}
}
