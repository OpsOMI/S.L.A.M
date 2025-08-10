package private

import (
	"strconv"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (r *Requester) RoomRoutes() {
	r.routes["/join"] = r.HandleJoin
	r.routes["/hidden"] = r.HandleHidden
	r.routes["/room/create"] = r.HandleCreate
	r.routes["/room/list"] = r.HandleList
}

func (r *Requester) HandleJoin(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	req.RequestID = commons.RequestIDJoinRoom
	if len(cmd.Args) != 1 {
		return apperrors.NewNotification("usage: /join [roomcode]")
	}

	if err := r.api.Users().Join(req, cmd.Args[0]); err != nil {
		return err
	}

	return nil
}

func (r *Requester) HandleCreate(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	req.RequestID = commons.RequestIDCreateRoom

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

	if err := r.api.Rooms().Create(req, isSecure); err != nil {
		return err
	}

	return nil
}

func (r *Requester) HandleList(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	req.RequestID = commons.RequestIDListRoom
	var page, limit int32
	if len(cmd.Args) != 0 {
		if cmd.Args[0] == "help" {
			return apperrors.NewNotification("usage: /room/list [page:(optional)] [limit:(optional)]")
		}

		p, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return apperrors.NewError("page must be an integer value.")
		}
		page = int32(p)

		if len(cmd.Args) > 1 {
			lim, err := strconv.Atoi(cmd.Args[1])
			if err != nil {
				return apperrors.NewError("page must be an integer value.")
			}
			limit = int32(lim)
		}
	}

	if err := r.api.Rooms().List(req, page, limit); err != nil {
		return err
	}

	return nil
}

func (r *Requester) HandleHidden(
	cmd parser.Command,
	req *request.ClientRequest,
) error {
	r.terminal.SetRooms(nil)

	return apperrors.NewNotification("Rooms were hidden!")
}
