package rooms

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (s *module) List(
	req *request.ClientRequest,
	page, limit int32,
) error {
	payload := rooms.ListRoomReq{
		Page:  page,
		Limit: limit,
	}

	if err := utils.SendRequest(s.conn, req, payload); err != nil {
		return err
	}

	return nil
}

func (s *module) Create(
	req *request.ClientRequest,
	isSecure bool,
) error {
	var password string
	if isSecure {

		password, err := utils.ReadPassword("Password")
		if err != nil {
			return err
		}
		fmt.Println()

		confirmPassword, err := utils.ReadPassword("Confirm Password")
		if err != nil {
			return err
		}

		if password != confirmPassword {
			return apperrors.NewNotification("Passwords do not match!")
		}
	}

	payload := rooms.CreateReq{
		Password: password,
	}

	if err := utils.SendRequest(s.conn, req, payload); err != nil {
		return err
	}

	return nil
}
