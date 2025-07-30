package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/tokenstore"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/controllers/public"
)

type Controller struct {
	listener   net.Listener
	logger     logger.ILogger
	config     config.Configs
	tokenstore tokenstore.ITokenManager
	// router   *Router
}

func NewController(
	listener net.Listener,
	logger logger.ILogger,
	config config.Configs,
	// router *Router,
) *Controller {
	tokenstore := tokenstore.NewJWTManager(config.Server.Jwt.Issuer, config.Server.Jwt.Secret)

	return &Controller{
		listener:   listener,
		logger:     logger,
		config:     config,
		tokenstore: tokenstore,
		// router:   router,
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

type ClientMessage struct {
	JwtToken string          `json:"jwt_token"` // JWT token for authentication and authorization
	Command  string          `json:"command"`   // Command to execute, e.g., "/join", "/message"
	Payload  json.RawMessage `json:"payload"`   // Command-specific data in JSON format
	Scope    string          `json:"scope"`     // User scope or role, e.g., "public", "private", "owner"
}

func (c *Controller) HandleConnection(conn net.Conn) {
	defer conn.Close()

	c.logger.Info("New connection accepted: " + conn.RemoteAddr().String())
	fmt.Fprintln(conn, "Welcome to SLAM!")

	public := public.NewController()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var msg ClientMessage
		err := json.Unmarshal([]byte(line), &msg)
		if err != nil {
			c.logger.Error("Invalid JSON format from " + conn.RemoteAddr().String() + ": " + err.Error())
			fmt.Fprintln(conn, `{"error":"invalid JSON format"}`)
			continue
		}

		var routeErr error
		switch msg.Scope {
		case "public":
			routeErr = public.Route(conn, msg.Command, msg.Payload)
		case "private":
			// routeErr = private.Route(conn, msg.Command, msg.Payload)
		case "owner":
			// routeErr = owner.Route(conn, msg.Command, msg.Payload)
		default:
			routeErr = fmt.Errorf("invalid scope: %s", msg.Scope)
		}

		if routeErr != nil {
			c.logger.Error("Routing error for " + conn.RemoteAddr().String() + ": " + routeErr.Error())
			fmt.Fprintf(conn, `{"error":"%s"}`+"\n", routeErr.Error())
			continue
		}

		c.logger.Info("Command received from " + conn.RemoteAddr().String() + ": " + msg.Command)
		fmt.Fprintf(conn, `{"info":"command received","command":"%s"}`+"\n", msg.Command)
	}

	if err := scanner.Err(); err != nil {
		c.logger.Error("Connection error for " + conn.RemoteAddr().String() + ": " + err.Error())
	}
}
