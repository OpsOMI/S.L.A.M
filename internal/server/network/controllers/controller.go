package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/infrastructure/connection"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/controllers/owner"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/controllers/private"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/controllers/public"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/server/services"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type Controller struct {
	listener    net.Listener
	logger      logger.ILogger
	config      config.Configs
	store       store.IJwtManager
	connections *connection.ConnectionManager
}

func NewController(
	listener net.Listener,
	logger logger.ILogger,
	config config.Configs,
) *Controller {
	tokenstore := store.NewManager(config.Server.Jwt.Issuer, config.Server.Jwt.Secret)
	connections := connection.NewConnectionManager()

	return &Controller{
		listener:    listener,
		logger:      logger,
		config:      config,
		store:       tokenstore,
		connections: connections,
	}
}
func (c *Controller) Start(
	services services.IServices,
) error {
	public := public.NewController(c.logger, c.store, services, c.connections)
	private := private.NewController(c.logger, c.store, services)
	owner := owner.NewController(c.logger, c.store, services)

	for {
		conn, err := c.listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		fmt.Println("New connection:", conn.RemoteAddr())

		go c.HandleConnection(conn, public, private, owner)
	}
}

func (c *Controller) HandleConnection(
	conn net.Conn,
	public *public.Controller,
	private *private.Controller,
	owner *owner.Controller,
) {
	defer conn.Close()

	c.logger.Info("New connection accepted: " + conn.RemoteAddr().String())

	_ = request.Send(conn, map[string]string{"message": "Welcome to SLAM!"})

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var msg request.ClientRequest
		err := json.Unmarshal([]byte(line), &msg)
		if err != nil {
			c.logger.Error("Invalid JSON from " + conn.RemoteAddr().String() + ": " + err.Error())
			continue
		}

		var routeMsg error
		switch msg.Scope {
		case "public":
			routeMsg = public.Route(conn, msg.Command, msg.Payload)
		case "private":
			routeMsg = private.Route(conn, msg.JwtToken, msg.Command, msg.Payload)
		case "owner":
			routeMsg = owner.Route(conn, msg.JwtToken, msg.Command, msg.Payload)
		default:
			routeMsg = response.Response("status.internal", "Invalid Scope", nil)
		}

		_ = response.Handle(conn, routeMsg)
		c.logger.Info("Command received from " + conn.RemoteAddr().String() + ": " + msg.Command)
	}

	if err := scanner.Err(); err != nil {
		c.logger.Error("Connection error for " + conn.RemoteAddr().String() + ": " + err.Error())
	}
}
