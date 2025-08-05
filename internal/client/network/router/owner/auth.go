package owner

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Router) AuthRoutes() {
	r.routes["/register"] = r.HandleRegister
}

func (r *Router) HandleRegister(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	id, err := r.api.Users().Register(req)
	if err != nil {
		return err
	}

	fmt.Println(id, "SADA")

	return nil
}
