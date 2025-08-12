package owner

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type Requester struct {
	api      api.IAPI
	terminal *terminal.Terminal
	store    *store.SessionStore
	routes   map[string]commons.RouteFunc
}

func NewRequester(
	terminal *terminal.Terminal,
	store *store.SessionStore,
	api api.IAPI,
) Requester {
	r := Requester{
		api:      api,
		store:    store,
		terminal: terminal,
		routes:   make(map[string]commons.RouteFunc),
	}

	r.AuthRoutes()
	r.UserRoutes()

	return r
}

func (r *Requester) Supports(name string) bool {
	_, exists := r.routes[name]
	return exists
}

func (r *Requester) Route(
	command parser.Command,
	req *request.ClientRequest,
) error {
	req.JwtToken = r.store.JWT
	if req.JwtToken == "" {
		return apperrors.NewError("Unauthorized: " + command.Name)
	}
	r.store.ParseJWT()
	if r.store.Role != "owner" {
		return apperrors.NewError("Unauthorized: " + command.Name)
	}
	req.Scope = "owner"

	if handler, ok := r.routes[command.Name]; ok {
		return handler(command, req)
	}

	return apperrors.NewError("Unknown Command:" + command.Name)
}
