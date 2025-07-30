package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/connection"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/tokenstore"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/controllers/private"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/controllers/public"
)

type Controller struct {
	listener    net.Listener
	logger      logger.ILogger
	config      config.Configs
	tokenstore  tokenstore.ITokenStore
	connmanager *connection.ConnectionManager
}

func NewController(
	listener net.Listener,
	logger logger.ILogger,
	config config.Configs,
) *Controller {
	tokenstore := tokenstore.NewJWTManager(config.Server.Jwt.Issuer, config.Server.Jwt.Secret)
	connmanager := connection.NewConnectionManager()

	return &Controller{
		listener:    listener,
		logger:      logger,
		config:      config,
		tokenstore:  tokenstore,
		connmanager: connmanager,
	}
}

func (c *Controller) Start() error {
	for {
		conn, err := c.listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		fmt.Println("New connection:", conn.RemoteAddr())

		go c.HandleConnection(conn)
	}
}

func (c *Controller) HandleConnection(conn net.Conn) {
	defer conn.Close()

	c.logger.Info("New connection accepted: " + conn.RemoteAddr().String())

	if err := response.Success(conn, map[string]string{"message": "Welcome to SLAM!"}); err != nil {
		c.logger.Error("Failed to send welcome message: " + err.Error())
		return
	}

	public := public.NewController(c.logger)
	private := private.NewController(c.logger, c.tokenstore)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var msg request.ClientMessage
		err := json.Unmarshal([]byte(line), &msg)
		if err != nil {
			c.logger.Error("Invalid JSON format from " + conn.RemoteAddr().String() + ": " + err.Error())

			_ = response.Error(conn, "invalid JSON format")
			continue
		}

		switch msg.Scope {
		case "public":
			public.Route(conn, msg.Command, msg.Payload)
		case "private":
			private.Route(conn, msg.JwtToken, msg.Command, msg.Payload)
		case "owner":
			// owner.Route(conn, msg.Command, msg.Payload)
		default:
			_ = response.Error(conn, fmt.Sprintf("invalid scope: %s", msg.Scope))
			continue
		}

		c.logger.Info("Command received from " + conn.RemoteAddr().String() + ": " + msg.Command)
	}

	if err := scanner.Err(); err != nil {
		c.logger.Error("Connection error for " + conn.RemoteAddr().String() + ": " + err.Error())
	}
}
