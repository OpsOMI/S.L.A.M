package app

import (
	"encoding/json"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
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

	// Server'dan ilk mesajı oku (Welcome mesajı gibi)
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		logg.Error("Failed to read message from server: " + err.Error())
		return
	}
	message := string(buf[:n])
	logg.Info("Received message from server: " + message)

	clientMsg := request.ClientMessage{
		JwtToken: "",
		Command:  "/ping",
		Payload:  json.RawMessage(`{"message":"Selam Ping Pong"}`),
		Scope:    "public",
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
	logg.Info("Sent ping command to server")

	n, err = conn.Read(buf)
	if err != nil {
		logg.Error("Failed to read response from server: " + err.Error())
		return
	}
	responseMsg := string(buf[:n])
	logg.Info("Received response from server: " + responseMsg)

	time.Sleep(1 * time.Second)
}
