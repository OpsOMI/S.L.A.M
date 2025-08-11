package public

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Requester) AuthRoutes() {
	r.routes["/login"] = r.HandleLogin
}

func (r *Requester) HandleLogin(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	req.RequestID = commons.RequestIDLogin
	if r.store.JWT != "" {
		return apperrors.NewNotification("Already logged in")
	}

	if err := r.api.Users().Login(req, r.cfg.ClientKey); err != nil {
		return err
	}

	return nil
}
