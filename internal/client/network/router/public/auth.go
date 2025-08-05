package public

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Router) AuthRoutes() {
	r.routes["/login"] = r.HandleLogin
	r.routes["/me"] = r.HandleMe
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

	return nil
}

func (r *Router) HandleMe(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	fmt.Println("me called")
	return nil
}
