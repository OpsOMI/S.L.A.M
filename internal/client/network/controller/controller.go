package controller

import (
	"bufio"
	"encoding/json"
	"net"
	"strings"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
	"github.com/OpsOMI/S.L.A.M/internal/client/infrastructure/network"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/api"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/requester"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/responder"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

type Controller struct {
	// Mutablable connection and related structures
	conn      net.Conn
	api       api.IAPI
	requester requester.Requesters
	responder responder.Responder

	// Immutable Deps
	logger   logger.ILogger
	terminal *terminal.Terminal
	parser   parser.IParser
	store    *store.SessionStore
	config   config.Configs

	// Channels
	done      chan struct{}
	responses chan response.BaseResponse
	inputChan chan string
}

func NewController(
	conn net.Conn,
	logger logger.ILogger,
	config config.Configs,
) *Controller {
	api := api.NewAPI(conn, logger)
	parser := parser.NewParser()
	terminal := terminal.NewTerminal()
	store := store.NewSessionStore()
	requester := requester.NewRequesters(api, store, terminal, &config)
	responder := responder.NewResponder(store, terminal)

	return &Controller{
		api:       api,
		conn:      conn,
		store:     store,
		config:    config,
		logger:    logger,
		parser:    parser,
		terminal:  terminal,
		requester: requester,
		responder: responder,
		done:      make(chan struct{}),
		inputChan: make(chan string),
		responses: make(chan response.BaseResponse, 200),
	}
}

func (c *Controller) Run() {
	c.terminal.Render()
	c.terminal.SetConnected(c.conn != nil)
	c.terminal.ClearScreen()

	if c.conn != nil {
		c.ListenServerMessages()
	}

	go c.responder.Listen(c.responses)

	for {
		select {
		case <-c.done:
			c.cleanup()
			return
		default:
		}

		if !c.isConnected() {
			c.terminal.SetConnected(false)
		} else {
			c.terminal.SetConnected(true)
		}

		input := c.HandleUserInput()
		c.handleInput(input)
	}
}

func (c *Controller) HandleUserInput() string {
	c.terminal.SetPromptLabel("->", c.store.Nickname)
	c.terminal.Render()

	input, err := c.terminal.Prompt()
	if err != nil {
		c.logger.Error("Error reading input: " + err.Error())
	}

	return input
}

func (c *Controller) handleInput(input string) {
	switch {
	case input == "/exit" || input == "/quit":
		c.logger.Info("User exited the client.")
		c.terminal.ClearScreen()
		close(c.done)
		return

	case input == "/clear":
		c.terminal.ClearScreen()
		return

	case input == "/logout":
		c.store.Logout()
		c.terminal.SetMessages(nil)
		c.terminal.SetRooms(nil)
		c.terminal.SetCurrentRoom("")
		c.terminal.Render()

	case input == "/reconnect":
		if err := c.Reconnect(); err != nil {
			c.logger.Warn("Reconnect failed: " + err.Error())
			c.terminal.PrintError("Could not reconnect to the server.")
			return
		}
		c.terminal.PrintNotification("Reconnected successfully.")
		c.logger.Info("Reconnected successfully.")

	case strings.HasPrefix(input, "/"):
		command, err := c.parser.Parse(input)
		if err != nil {
			c.logger.Warn("Invalid command: " + err.Error())
			c.terminal.PrintError("Invalid command syntax.")
			return
		}
		if err := c.requester.SendRequest(command); err != nil {
			c.terminal.Print(err)
		}

	default:
		if input != "" && c.store.Room != "" {
			if err := c.api.Users().SendMessage(&request.ClientRequest{
				RequestID: commons.RequestIDSendMessage,
				JwtToken:  c.store.JWT,
				Scope:     "private",
				Command:   "/send",
			}, input,
			); err != nil {
				c.terminal.Print(err)
			}
			c.terminal.PrintMessage("You", input)
		}
	}
}

func (c *Controller) isConnected() bool {
	if c.conn == nil {
		return false
	}

	_ = c.conn.SetReadDeadline(time.Now().Add(1 * time.Millisecond))

	one := make([]byte, 1)
	_, err := c.conn.Read(one)

	_ = c.conn.SetReadDeadline(time.Time{})

	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return true
		}
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
	c.api = api.NewAPI(c.conn, c.logger)
	c.requester = requester.NewRequesters(c.api, c.store, c.terminal, &c.config)

	return nil
}

func (c *Controller) ListenServerMessages() {
	go func() {
		reader := bufio.NewReader(c.conn)

		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}

				c.logger.Warn("Server message read error: " + err.Error())
				break
			}

			var serverResp response.BaseResponse
			if err := json.Unmarshal([]byte(msg), &serverResp); err != nil {
				c.logger.Warn("Invalid server message format: " + err.Error())
				continue
			}

			c.responses <- serverResp
		}
	}()
}

func (c *Controller) cleanup() {
	if c.conn != nil {
		c.conn.Close()
	}
	c.logger.Info("Client gracefully shut down.")
}
