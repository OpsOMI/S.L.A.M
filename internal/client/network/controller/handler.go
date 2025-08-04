package controller

import (
	"net"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/router"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
)

type Controller struct {
	conn     net.Conn
	logger   logger.ILogger
	terminal terminal.Terminal
	parser   parser.IParser
	router   router.Router
	store    *store.SessionStore
}

func NewController(
	conn net.Conn,
	logger logger.ILogger,
) *Controller {
	terminal := terminal.NewTerminal()
	parser := parser.NewParser()
	api := api.NewAPI(conn, logger)
	store := store.NewSessionStore()
	router := router.NewRouter(api, store)

	return &Controller{
		conn:     conn,
		logger:   logger,
		terminal: *terminal,
		parser:   parser,
		router:   router,
		store:    store,
	}
}

func (c *Controller) Run() {
	for {
		input, err := c.terminal.Prompt("-> ", c.store.Nickname)
		if err != nil {
			c.logger.Error("Error reading input: " + err.Error())
			continue
		}

		if input == "exit" || input == "quit" {
			c.logger.Info("User exited the client.")
			break
		}

		if input == "clear" {
			c.terminal.ClearScreen()
			c.terminal.ClearError()
			continue
		}

		if strings.HasPrefix(input, "/") {
			// This is where the clients trying commands.
			command, err := c.parser.Parse(input)
			if err != nil {
				c.logger.Warn("Invalid command syntax, ignoring parse error: " + err.Error())
				continue
			}

			if err := c.router.Route(command); err != nil {
				c.terminal.PrintError(err.Error())
			}
		} else {
			// This is where the client trying to send a message to another client.
			c.terminal.PrintMessage(input) // This is for know. This messages will recived from db.
		}
	}
}
