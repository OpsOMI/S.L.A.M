package private

import (
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Router) RoomRoutes() {
	r.routes["/join"] = r.HandleJoin
	r.routes["/room/create"] = r.HandleCreate
}

func (r *Router) HandleJoin(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	if len(cmd.Args) != 1 {
		return apperrors.NewNotification("usage: /join [roomcode]")
	}

	messages, err := r.api.Users().Join(req, cmd.Args[0])
	if err != nil {
		return err
	}

	r.store.SetRoom(cmd.Args[0])
	r.terminal.SetMessages(messages)

	return apperrors.NewNotification("Joined Successfully: " + cmd.Args[0])
}

func (r *Router) HandleCreate(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	if len(cmd.Args) != 1 {
		return apperrors.NewNotification("usage: /room/create [isSecure]:[True/False]")
	}

	if cmd.Args[0] == "help" {
		return apperrors.NewNotification("usage: /room/create [isSecure]:[True/False]")
	}

	isSecureStr := strings.ToLower(cmd.Args[0])
	var isSecure bool
	switch isSecureStr {
	case "true":
		isSecure = true
	case "false":
		isSecure = false
	default:
		return apperrors.NewError("invalid argument: use 'True' or 'False'")
	}

	code, err := r.api.Rooms().Create(req, isSecure)
	if err != nil {
		return err
	}

	return apperrors.NewNotification("Room Created Successfully, Code: " + code)
}
