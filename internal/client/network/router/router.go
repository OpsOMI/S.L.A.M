package router

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/client/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/router/public"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
)

type Router struct {
	public public.Router
}

func NewRouter(
	api api.IAPI,
	store *store.SessionStore,
) Router {
	public := public.NewRouter(
		store,
		api,
	)

	return Router{
		public: public,
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

	return nil
}
