package rooms

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (s *module) List(
	req *request.ClientRequest,
	page, limit int32,
) (*rooms.RoomsResp, error) {
	payload := rooms.MyRoomReq{
		Page:  page,
		Limit: limit,
	}

	baseResp, err := utils.SendRequest(s.conn, req, payload)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckBaseResponse(baseResp); err != nil {
		return nil, err
	}

	var data rooms.RoomsResp
	if err := utils.LoadData(baseResp.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *module) Create(
	req *request.ClientRequest,
	isSecure bool,
) (*string, error) {
	var password string
	if isSecure {

		password, err := utils.ReadPassword("Password")
		if err != nil {
			return nil, err
		}

		confirmPassword, err := utils.ReadPassword("Confirm Password")
		if err != nil {
			return nil, err
		}

		if password != confirmPassword {
			return nil, apperrors.NewNotification("Passwords do not match!")
		}
	}

	payload := rooms.CreateReq{
		Password: password,
	}

	baseResp, err := utils.SendRequest(s.conn, req, payload)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckBaseResponse(baseResp); err != nil {
		return nil, err
	}

	var data rooms.CreateResp
	if err := utils.LoadData(baseResp.Data, &data); err != nil {
		return nil, err
	}

	return &data.Code, nil
}
