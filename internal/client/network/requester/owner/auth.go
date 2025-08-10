package owner

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Requester) AuthRoutes() {
	r.routes["/register"] = r.HandleRegister
}

func (r *Requester) HandleRegister(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	req.RequestID = commons.RequestIDRegister
	if err := r.api.Users().Register(req); err != nil {
		return err
	}

	return nil
}
