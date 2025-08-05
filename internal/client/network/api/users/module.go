package users

import (
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
	"github.com/google/uuid"
)

type IUserModule interface {
	Login(
		req *request.ClientRequest,
	) (string, error)

	Register(
		req *request.ClientRequest,
	) (uuid.UUID, error)
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
