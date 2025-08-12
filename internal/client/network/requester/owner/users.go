package owner

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Requester) UserRoutes() {
	r.routes["/online"] = r.HandleOnline
}

func (r *Requester) HandleOnline(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	req.RequestID = commons.RequestIDOnline
	if err := r.api.Users().Online(req); err != nil {
		return err
	}

	return nil
}
