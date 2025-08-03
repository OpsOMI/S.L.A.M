package controller

import (
	"fmt"
	"net"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/api"
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
	c.terminal.ShowWelcome()

	for {
		input, err := c.terminal.Prompt(">: ")
		if err != nil {
			c.logger.Error("Error reading input: " + err.Error())
			continue
		}

		if input == "exit" || input == "quit" {
			c.logger.Info("User exited the client.")
			break
		}

		if strings.HasPrefix(input, "/") {
			// This is where the clients trying commands.
			command, err := c.parser.Parse(input)
			if err != nil {
				c.logger.Warn("Invalid command syntax, ignoring parse error: " + err.Error())
				continue
			}

			c.logger.Info("User typed command: " + input)
			fmt.Println(command)

			if err := c.router.Route(command); err != nil {
				c.logger.Warn("Unknown command or routing error: " + err.Error())
			}
		} else {
			// This is where the client trying to send a message to another client.
			c.logger.Info("User typed message: " + input)
			fmt.Println(input)
		}
	}
}
