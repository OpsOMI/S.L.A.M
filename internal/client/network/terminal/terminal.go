package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

type Terminal struct {
	reader    *bufio.Reader
	messages  []string
	output    Notification
	label     string
	connected bool
	height    int
	width     int
}

type Notification struct {
	Code    string // Error, Information
	Message string
}

func NewTerminal() *Terminal {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24 // default fallback
		width = 80
	}

	return &Terminal{
		reader:    bufio.NewReader(os.Stdin),
		messages:  []string{},
		connected: true,
		height:    height,
		width:     width,
	}
}

func (t *Terminal) moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func (t *Terminal) Render() {
	t.ClearScreen()

	// Welcome message
	welcome := "Welcome to S.L.A.M Client"
	col := max((t.width-len(welcome))/2, 1)
	t.moveCursor(1, col)
	fmt.Print(welcome)

	// Connected Dot
	dotCol := t.width
	t.moveCursor(1, dotCol)

	if t.connected {
		fmt.Print("\033[32m⬤\033[0m")
	} else {
		fmt.Print("\033[31m⬤\033[0m")
	}

	// Start messages
	msgStartRow := 3
	maxMessages := t.height - msgStartRow - 2
	if len(t.messages) > maxMessages {
		t.messages = t.messages[len(t.messages)-maxMessages:]
	}

	for i, msg := range t.messages {
		t.moveCursor(msgStartRow+i, 1)
		fmt.Print(msg)
	}

	// Print outputMessage (error or notification) on line above prompt (height-1)
	if t.output.Message != "" {
		t.moveCursor(t.height-1, 1)

		if t.output.Code == "error" {
			fmt.Printf("\033[31m%s\033[0m", t.output.Message)
		} else {
			fmt.Printf("\033[32m%s\033[0m", t.output.Message)
		}
	}

	// Print prompt at bottom line (this requires promptLabel to be stored)
	if t.label != "" {
		t.moveCursor(t.height, 1)
		fmt.Print(t.label)
	}
}

func (t *Terminal) Prompt() (string, error) {
	t.moveCursor(t.height, 1)
	fmt.Print(t.label)

	input, err := t.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// Setter
func (t *Terminal) SetConnected(status bool) {
	t.connected = status
	t.Render()
}

func (t *Terminal) SetPromptLabel(label, nickname string) {
	var promptLabel string
	if nickname != "" {
		promptLabel = "\033[31m[" + nickname + "]\033[0m " + label + " "
	} else {
		promptLabel = "\033[34m[Unknown]\033[0m " + label + " "
	}

	t.label = promptLabel
}

// Clears
func (t *Terminal) ClearScreen() {
	fmt.Print("\033[2J") // Clear entire screen
	fmt.Print("\033[H")  // Move cursor to top-left
}

func (t *Terminal) ClearOutput() {
	t.output.Code = ""
	t.output.Message = ""
	t.Render()
}

// Prints
func (t *Terminal) PrintMessage(msg string) {
	t.messages = append(t.messages, msg)
	// Crop messages if they exceed available screen space (height - 2 lines for error + prompt)
	maxMessages := t.height - 2
	if len(t.messages) > maxMessages {
		t.messages = t.messages[len(t.messages)-maxMessages:]
	}
	t.Render()
}

func (t *Terminal) PrintError(msg string) {
	t.output.Message = msg
	t.output.Code = "error"
	t.Render()

	go func() {
		time.Sleep(1 * time.Second)
		t.ClearOutput()
	}()
}

func (t *Terminal) PrintNotification(msg string) {
	t.output.Message = msg
	t.output.Code = "info"
	t.Render()

	go func() {
		time.Sleep(1 * time.Second)
		t.ClearOutput()
	}()
}
