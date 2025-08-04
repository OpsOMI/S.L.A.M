package public

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/client/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/types"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type Router struct {
	api    api.IAPI
	store  *store.SessionStore
	routes map[string]types.RouteFunc
}

func NewRouter(
	store *store.SessionStore,
	api api.IAPI,
) Router {
	r := Router{
		api:    api,
		store:  store,
		routes: make(map[string]types.RouteFunc),
	}

	r.AuthRoutes()

	return r
}

func (r *Router) Supports(name string) bool {
	_, exists := r.routes[name]
	return exists
}

func (r *Router) Route(
	command parser.Command,
	req *request.ClientRequest,
) error {
	req.JwtToken = ""
	req.Scope = "public"

	if handler, ok := r.routes[command.Name]; ok {
		return handler(command, req)
	}

	return fmt.Errorf("unknown public command: %s", command.Name)
}
