package private

import (
	"fmt"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Router) RoomRoutes() {
	r.routes["/join"] = r.HandleJoin
	r.routes["/room/create"] = r.HandleCreate
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

func (r *Router) HandleCreate(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: /room/create [isSecure]:[True/False]") // eksik ya da fazla argüman
	}

	if cmd.Args[0] == "help" {
		return fmt.Errorf("usage: /room/create [isSecure]:[True/False]")
	}

	isSecureStr := strings.ToLower(cmd.Args[0])
	var isSecure bool
	switch isSecureStr {
	case "true":
		isSecure = true
	case "false":
		isSecure = false
	default:
		return fmt.Errorf("invalid argument: use 'True' or 'False'")
	}

	code, err := r.api.Rooms().Create(req, isSecure)
	if err != nil {
		return err
	}
	r.terminal.PrintNotification("Room Created Successfully, Code: = " + code)

	return nil
}
