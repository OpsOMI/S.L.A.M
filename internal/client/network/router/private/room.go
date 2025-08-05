package private

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Router) RoomRoutes() {
	r.routes["/join"] = r.HandleJoin
}

func (r *Router) HandleJoin(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Hata")
	}

	messages, err := r.api.Users().Join(req, cmd.Args[0])
	if err != nil {
		return err
	}

	r.store.SetRoom(cmd.Args[0])
	r.terminal.SetMessages(messages)
	r.terminal.PrintNotification("Joined Successfully: " + cmd.Args[0])

	return nil
}
