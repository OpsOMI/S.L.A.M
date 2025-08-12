package users

import (
	"bufio"
	"os"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/message"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (s *module) Login(
	req *request.ClientRequest,
	clientKey string,
) error {
	reader := bufio.NewReader(os.Stdin)

	username, err := utils.Read(reader, "Username")
	if err != nil {
		return err
	}

	password, err := utils.ReadPassword("Password")
	if err != nil {
		return err
	}

	payload := users.LoginReq{
		ClientKey: clientKey,
		Username:  username,
		Password:  password,
	}

	if err := utils.SendRequest(s.conn, req, payload); err != nil {
		return err
	}

	return nil
}

func (s *module) Register(
	req *request.ClientRequest,
) error {
	reader := bufio.NewReader(os.Stdin)

	nickname, err := utils.Read(reader, "Nickname")
	if err != nil {
		return err
	}

	username, err := utils.Read(reader, "Username")
	if err != nil {
		return err
	}

	password, err := utils.ReadPassword("Password")
	if err != nil {
		return err
	}

	payload := users.RegisterReq{
		Nickname: nickname,
		Username: username,
		Password: password,
	}

	if err := utils.SendRequest(s.conn, req, payload); err != nil {
		return err
	}

	return nil
}

func (s *module) Join(
	req *request.ClientRequest,
	roomCode string,
) error {
	password, err := utils.ReadPassword("Password")
	if err != nil {
		return err
	}

	roomCode = strings.TrimSpace(roomCode)

	payload := rooms.JoinReq{
		RoomCode: roomCode,
		Password: password,
	}

	if err := utils.SendRequest(s.conn, req, payload); err != nil {
		return err
	}

	return nil
}

func (s *module) SendMessage(
	req *request.ClientRequest,
	content string,
) error {
	content = strings.TrimSpace(content)

	payload := message.MessageReq{
		Content: content,
	}

	if err := utils.SendRequest(s.conn, req, payload); err != nil {
		return err
	}

	return nil
}

func (s *module) Online(
	req *request.ClientRequest,
) error {
	if err := utils.SendRequest(s.conn, req, nil); err != nil {
		return err
	}

	return nil
}

func (s *module) Me(
	req *request.ClientRequest,
) error {
	if err := utils.SendRequest(s.conn, req, nil); err != nil {
		return err
	}

	return nil
}
