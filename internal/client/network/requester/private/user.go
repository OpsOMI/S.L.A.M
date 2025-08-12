package private

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Requester) UserRoutes() {
	r.routes["/me"] = r.Me
}

func (r *Requester) Me(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	req.RequestID = commons.RequestIDMe

	if err := r.api.Users().Me(req); err != nil {
		return err
	}

	return nil
}
