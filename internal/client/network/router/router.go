package router

import (
	"errors"

	"github.com/OpsOMI/S.L.A.M/internal/client/network/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/router/owner"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/router/public"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type Router struct {
	public public.Router
	owner  owner.Router
}

func NewRouter(
	api api.IAPI,
	store *store.SessionStore,
	terminal *terminal.Terminal,
) Router {
	public := public.NewRouter(
		store,
		api,
	)
	owner := owner.NewRouter(
		terminal,
		store,
		api,
	)

	return Router{
		public: public,
		owner:  owner,
	}
}

func (r *Router) Route(
	command parser.Command,
) error {
	req := request.ClientRequest{
		Command: command.Name,
	}

	if r.public.Supports(command.Name) {
		return r.public.Route(command, &req)
	}
	if r.owner.Supports(command.Name) {
		return r.owner.Route(command, &req)
	}

	return errors.New("Unknown command: " + command.Name)
}
