package app

import (
	"encoding/json"
	"log"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/mappers/users"
)

func Run(cfg *config.Configs) {
	logg, err := logger.NewZapLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logg.Sync()

	logg.Info("Welcome to the S.L.A.M client!")

	logg.Info("Attempting to connect to server...")
	conn, err := network.ConnectToServer(
		cfg.ServerName,
		cfg.ServerHost,
		cfg.ServerPort,
		cfg.TSLCertPath,
		cfg.TimeoutSeconds,
		cfg.ReconnectRetry,
	)
	if err != nil {
		logg.Error("Failed to connect to server: " + err.Error())
		return
	}
	defer conn.Close()
	logg.Info("Successfully connected to server")

	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		logg.Error("Failed to read message from server: " + err.Error())
		return
	}
	message := string(buf[:n])
	logg.Info("Received message from server: " + message)

	// req := users.LoginReq{
	// 	Username: "kaan",
	// 	Password: "kaaaan",
	// }

	req := users.RegisterReq{
		Nickname: "slm",
		Username: "cbmmm",
		Password: "kaaaan",
	}

	payloadBytes, err := json.Marshal(req)
	if err != nil {
		log.Fatal("failed to marshal login payload:", err)
	}

	clientMsg := request.ClientMessage{
		JwtToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOiI3ZWJhNDhlYS02ZDM0LTQ3YTEtYWFmNS00MDliN2MxNzJhZDciLCJ1c2VyX2lkIjoiYzQ1MjZlY2EtY2ZjNy00NTU1LTgwNDgtYmUzMTY5ZmE4ZTYyIiwidXNlcm5hbWUiOiJrYWFuIiwibmlja25hbWUiOiJrYWFuIiwicm9sZSI6InVzZXIiLCJpc3MiOiJTTEFNIiwiZXhwIjoxNzU0MzE2ODg1LCJpYXQiOjE3NTQyMzA0ODV9.gh9H5mg5OwgRkqHAxEhAMNxUPAtqTFp5_XmhZYju6Xo",
		Command:  "/auth/register",
		Payload:  payloadBytes,
		Scope:    "owner",
	}

	data, err := json.Marshal(clientMsg)
	if err != nil {
		logg.Error("Failed to marshal client message: " + err.Error())
		return
	}

	_, err = conn.Write(append(data, '\n'))
	if err != nil {
		logg.Error("Failed to send message to server: " + err.Error())
		return
	}
	logg.Info("Sent /me command to server")

	n, err = conn.Read(buf)
	if err != nil {
		logg.Error("Failed to read response from server: " + err.Error())
		return
	}
	responseMsg := string(buf[:n])
	logg.Info("Received response from server: " + responseMsg)

	time.Sleep(1 * time.Second)
}
