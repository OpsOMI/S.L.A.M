package controller

import (
	"net"
	"strings"
	"time"

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
	router := router.NewRouter(api, store, terminal)

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
	c.terminal.Render()
	c.terminal.SetConnected(c.conn != nil)

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
		input, err := c.terminal.Prompt()
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
			c.terminal.ClearOutput()
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
			// IN THE BACKGROUND WE WILL READ MSG FROM SERVER. THIS MESSAGES WILL ADDED TO THE TERMINAL
			// THIS IS WHERE THE CHAT IS GOING TO HAPPEN
			// NEED MESSAGE STRUCT! SENDER NICKNAME + CONENT
			c.terminal.PrintMessage(input) // This is for know. This messages will recived from db.
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
