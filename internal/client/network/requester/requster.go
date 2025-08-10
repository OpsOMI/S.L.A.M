package requester

import (
	"errors"

	"github.com/OpsOMI/S.L.A.M/internal/client/network/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/requester/owner"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/requester/private"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/requester/public"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type Requesters struct {
	public  public.Requester
	private private.Requester
	owner   owner.Requester
}

func NewRequesters(
	api api.IAPI,
	store *store.SessionStore,
	terminal *terminal.Terminal,
) Requesters {
	public := public.NewRequester(
		terminal,
		store,
		api,
	)
	private := private.NewRequester(
		terminal,
		store,
		api,
	)
	owner := owner.NewRequester(
		terminal,
		store,
		api,
	)

	return Requesters{
		public:  public,
		private: private,
		owner:   owner,
	}
}

func (r *Requesters) SendRequest(
	command parser.Command,
) error {
	req := request.ClientRequest{
		Command: command.Name,
	}

	if r.public.Supports(command.Name) {
		return r.public.Route(command, &req)
	}
	if r.private.Supports(command.Name) {
		return r.private.Route(command, &req)
	}
	if r.owner.Supports(command.Name) {
		return r.owner.Route(command, &req)
	}

	return errors.New("Unknown command: " + command.Name)
}
