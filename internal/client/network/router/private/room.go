package private

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Router) RoomRoutes() {
	r.routes["/join"] = r.HandleJoin
}

// TODO: MAKE CUSTOM ERROR FOR CLIENT. BEACUSE WE CANNOT CHOOSE NOTIFICATION OR ERROR MSG.
// THIS USAGE SHOULD BE NOTIFICATION.
func (r *Router) HandleJoin(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: /join [roomcode]")
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
