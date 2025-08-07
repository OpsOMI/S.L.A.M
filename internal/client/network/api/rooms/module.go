package rooms

import (
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type IRoomModule interface {
	MyRooms(
		req *request.ClientRequest,
		page, limit int32,
	) (*rooms.RoomsResp, error)

	Create(
		req *request.ClientRequest,
		isSecure bool,
	) (string, error)
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
