package controller

import (
	"net"
	"strings"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
	"github.com/OpsOMI/S.L.A.M/internal/client/infrastructure/network"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/router"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
)

type Controller struct {
	conn     net.Conn
	config   config.Configs
	logger   logger.ILogger
	terminal *terminal.Terminal
	parser   parser.IParser
	router   router.Router
	store    *store.SessionStore
}

func NewController(
	conn net.Conn,
	logger logger.ILogger,
	config config.Configs,
) *Controller {
	terminal := terminal.NewTerminal()
	parser := parser.NewParser()
	api := api.NewAPI(conn, logger)
	store := store.NewSessionStore()
	router := router.NewRouter(api, store, terminal)

	return &Controller{
		conn:     conn,
		config:   config,
		logger:   logger,
		terminal: terminal,
		parser:   parser,
		router:   router,
		store:    store,
	}
}

func (c *Controller) Run() {
	c.terminal.Render()
	c.terminal.SetConnected(c.conn != nil)
	c.terminal.ClearScreen()

	for {
		if !c.checkConnection() {
			c.terminal.SetConnected(false)
			if c.conn != nil {
				c.conn.Close()
				c.conn = nil
			}
		} else {
			c.terminal.SetConnected(true)
		}

		c.terminal.SetPromptLabel("->", c.store.Nickname)
		c.terminal.Render()

		input, err := c.terminal.Prompt()
		if err != nil {
			c.logger.Error("Error reading input: " + err.Error())
			continue
		}

		switch {
		case input == "/exit" || input == "/quit":
			c.logger.Info("User exited the client.")
			return

		case input == "/clear":
			c.terminal.ClearScreen()
			continue
		case input == "/reconnect":
			if err = c.Reconnect(); err != nil {
				c.logger.Warn("Failed to reconnect to the server: " + err.Error())
				c.terminal.PrintError("Could not reconnect to the server. Please check your connection.")
				continue
			}

			c.terminal.PrintNotification("Successfully reconnected to the server.")
			c.logger.Info("Reconnected to the server successfully.")

		case strings.HasPrefix(input, "/"):
			command, err := c.parser.Parse(input)
			if err != nil {
				c.logger.Warn("Invalid command syntax: " + err.Error())
				c.terminal.PrintError("Invalid command syntax.")
				continue
			}

			if err := c.router.Route(command); err != nil {
				c.terminal.PrintError(err.Error())
			}

		default:
			if input != "" {
				c.terminal.PrintMessage("You", input)
			}
		}
	}
}

func (c *Controller) checkConnection() bool {
	if c.conn == nil {
		return false
	}

	// Set a short read deadline to avoid blocking
	_ = c.conn.SetReadDeadline(time.Now().Add(1 * time.Millisecond))

	one := make([]byte, 1)
	_, err := c.conn.Read(one)

	// Reset the read deadline to default (no deadline)
	_ = c.conn.SetReadDeadline(time.Time{})

	if err != nil {
		// If the error is a timeout, it means no data was received but the connection is still alive
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return true
		}
		// Any other error indicates the connection is dead
		return false
	}

	return true
}

func (c *Controller) Reconnect() error {
	conn, err := network.Reconnect(c.conn, c.config)
	if err != nil {
		return err
	}

	c.conn = conn

	api := api.NewAPI(c.conn, c.logger)
	c.router = router.NewRouter(api, c.store, c.terminal)

	return nil
}
