package rooms

import (
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type IRoomModule interface {
	List(
		req *request.ClientRequest,
		page, limit int32,
	) error

	Create(
		req *request.ClientRequest,
		isSecure bool,
	) error

	Clean(
		req *request.ClientRequest,
		roomCode string,
	) error
}

type module struct {
	conn net.Conn
}

func NewModule(
	conn net.Conn,
) IRoomModule {
	return &module{
		conn: conn,
	}
}
