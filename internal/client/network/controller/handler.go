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
	"github.com/OpsOMI/S.L.A.M/internal/client/network/parser"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/router"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/message"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

type Controller struct {
	// Mutablable connection and related structures
	conn   net.Conn
	api    api.IAPI
	router router.Router

	// Immutable Deps
	logger   logger.ILogger
	terminal *terminal.Terminal
	parser   parser.IParser
	store    *store.SessionStore
	config   config.Configs

	// Channels
	done        chan struct{}            // Close to stop everything
	messageChan chan message.MessageResp // Server - ui messages
	inputChan   chan string              // Stdin lines
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
		conn:        conn,
		config:      config,
		logger:      logger,
		terminal:    terminal,
		parser:      parser,
		router:      router,
		store:       store,
		api:         api,
		done:        make(chan struct{}),
		messageChan: make(chan message.MessageResp, 200),
		inputChan:   make(chan string),
	}
}

func (c *Controller) Run() {
	c.terminal.Render()
	c.terminal.SetConnected(c.conn != nil)
	c.terminal.ClearScreen()

	// if c.conn != nil {
	// 	c.ListenServerMessages()
	// }

	go c.handleIncomingMessages()

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
		select {
		case msg := <-c.messageChan:
			roomCode := c.store.GetRoom()
			if roomCode == msg.RoomCode {
				c.terminal.PrintMessage(msg.SenderNickname, msg.Content)
			}
		case <-time.After(2 * time.Second):
			c.terminal.SetConnected(c.conn != nil)
		default:
			c.handleInput(input)
		}
	}
}

// Use this later.
func (c *Controller) HandleUserInputGorutine() {
	for {
		c.terminal.SetPromptLabel("->", c.store.Nickname)
		c.terminal.Render()

		input, err := c.terminal.Prompt()
		if err != nil {
			c.logger.Error("Error reading input: " + err.Error())
			continue
		}
		c.inputChan <- input
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

func (c *Controller) handleIncomingMessages() {
	for msg := range c.messageChan {
		roomCode := c.store.GetRoom()
		if roomCode == msg.RoomCode {
			c.terminal.PrintMessage(msg.SenderNickname, msg.Content)
		}
	}
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
		if err := c.router.Route(command); err != nil {
			c.terminal.Print(err)
		}

	default:
		if input != "" {
			if err := c.api.Users().SendMessage(&request.ClientRequest{
				JwtToken: c.store.JWT,
				Scope:    "private",
				Command:  "/send",
			}, input,
			); err != nil {
				c.logger.Warn("Send error: " + err.Error())
				/*
					FIXME
					You will see that this error message is triggered
					when you receive something from the user in special / commands.
				*/
				c.terminal.Print(err)
			}
			if c.store.GetRoom() != "" {
				c.terminal.PrintMessage("You", input)
			}
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
	c.router = router.NewRouter(c.api, c.store, c.terminal)

	// c.ListenServerMessages()

	return nil
}

func (c *Controller) ListenServerMessages() {
	go func() {
		reader := bufio.NewReader(c.conn)

		for {
			_ = c.conn.SetReadDeadline(time.Now().Add(1 * time.Second))

			msg, err := reader.ReadString('\n')

			_ = c.conn.SetReadDeadline(time.Time{})

			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}

				c.logger.Warn("Server message read error: " + err.Error())
				break
			}

			var serverMsg message.MessageResp
			if err := json.Unmarshal([]byte(msg), &serverMsg); err != nil {
				c.logger.Warn("Invalid server message format: " + err.Error())
				continue
			}

			c.messageChan <- serverMsg
		}
	}()
}

func (c *Controller) cleanup() {
	if c.conn != nil {
		c.conn.Close()
	}
	c.logger.Info("Client gracefully shut down.")
}
