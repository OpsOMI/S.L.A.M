package public

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
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
	cfg      *config.Configs
}

func NewRequester(
	cfg *config.Configs,
	terminal *terminal.Terminal,
	store *store.SessionStore,
	api api.IAPI,
) Requester {
	r := Requester{
		api:      api,
		cfg:      cfg,
		store:    store,
		terminal: terminal,
		routes:   make(map[string]commons.RouteFunc),
	}

	r.AuthRoutes()

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
	req.JwtToken = ""
	req.Scope = "public"

	if handler, ok := r.routes[command.Name]; ok {
		return handler(command, req)
	}

	return apperrors.NewError("unknown public command:" + command.Name)
}
