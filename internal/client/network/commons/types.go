package commons

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type RouteFunc func(cmd parser.Command, req *request.ClientRequest) error
