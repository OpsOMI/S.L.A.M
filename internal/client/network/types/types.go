package types

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
)

type RouteFunc func(cmd parser.Command, req *request.ClientRequest) error
