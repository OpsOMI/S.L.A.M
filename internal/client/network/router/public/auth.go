package public

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Router) AuthRoutes() {
	r.routes["/login"] = r.HandleLogin
}

func (r *Router) HandleLogin(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	jwt, err := r.api.Users().Login(req)
	if err != nil {
		return err
	}

	// Logged In.
	r.store.SetToken(jwt)
	r.store.ParseJWT()
	r.terminal.Render()

	return apperrors.NewNotification("Login Successful")
}
