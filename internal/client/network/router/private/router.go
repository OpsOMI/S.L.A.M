package private

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/client/network/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/types"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type Router struct {
	api      api.IAPI
	terminal *terminal.Terminal
	store    *store.SessionStore
	routes   map[string]types.RouteFunc
}

func NewRouter(
	terminal *terminal.Terminal,
	store *store.SessionStore,
	api api.IAPI,
) Router {
	r := Router{
		api:      api,
		store:    store,
		terminal: terminal,
		routes:   make(map[string]types.RouteFunc),
	}
	r.RoomRoutes()

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
	req.JwtToken = r.store.JWT
	req.Scope = "private"
	r.store.ParseJWT()

	if handler, ok := r.routes[command.Name]; ok {
		return handler(command, req)
	}

	return fmt.Errorf("unknown public command: %s", command.Name)
}
