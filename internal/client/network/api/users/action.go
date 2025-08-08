package users

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/message"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
	"github.com/google/uuid"
)

func (s *module) Login(
	req *request.ClientRequest,
) (*string, error) {
	reader := bufio.NewReader(os.Stdin)

	username, err := utils.Read(reader, "Username")
	if err != nil {
		return nil, err
	}

	password, err := utils.ReadPassword("Password")
	if err != nil {
		return nil, err
	}

	payload := users.LoginReq{
		Username: username,
		Password: password,
	}

	baseResp, err := utils.SendRequest(s.conn, req, payload)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckBaseResponse(baseResp); err != nil {
		return nil, err
	}

	var data users.LoginResp
	if err := utils.LoadData(baseResp.Data, &data); err != nil {
		return nil, err
	}

	return &data.Token, nil
}

func (s *module) Register(
	req *request.ClientRequest,
) (*uuid.UUID, error) {
	reader := bufio.NewReader(os.Stdin)

	nickname, err := utils.Read(reader, "Nickname")
	if err != nil {
		return nil, err
	}

	username, err := utils.Read(reader, "Username")
	if err != nil {
		return nil, err
	}

	password, err := utils.ReadPassword("Password")
	if err != nil {
		return nil, err
	}

	payload := users.RegisterReq{
		Nickname: nickname,
		Username: username,
		Password: password,
	}

	baseResp, err := utils.SendRequest(s.conn, req, payload)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckBaseResponse(baseResp); err != nil {
		return nil, err
	}

	var data users.RegisterResp
	if err := utils.LoadData(baseResp.Data, &data); err != nil {
		return nil, err
	}

	return &data.ID, nil
}

func (s *module) Join(
	req *request.ClientRequest,
	roomCode string,
) (*message.MessagesReps, error) {
	password, err := utils.ReadPassword("Password")
	if err != nil {
		return nil, err
	}

	roomCode = strings.TrimSpace(roomCode)

	payload := rooms.JoinReq{
		RoomCode: roomCode,
		Password: password,
	}

	baseResp, err := utils.SendRequest(s.conn, req, payload)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckBaseResponse(baseResp); err != nil {
		return nil, err
	}

	var data message.MessagesReps
	if err := utils.LoadData(baseResp.Data, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *module) SendMessage(
	req *request.ClientRequest,
	content string,
) error {
	content = strings.TrimSpace(content)

	payload := message.MessageReq{
		Content: content,
	}

	baseResp, err := utils.SendRequest(s.conn, req, payload)
	if err != nil {
		return err
	}

	if baseResp.Code != "OK" && baseResp.Message != "" {
		return fmt.Errorf("server error: %s ", baseResp.Message)
	}

	return apperrors.NewNotification("Sent!")
}
