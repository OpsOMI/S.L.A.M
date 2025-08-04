package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Terminal struct {
	reader       *bufio.Reader
	messages     []string
	errorMessage string // single error message shown above prompt
	height       int
	width        int
}

func NewTerminal() *Terminal {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24 // default fallback
		width = 80
	}

	return &Terminal{
		reader:   bufio.NewReader(os.Stdin),
		messages: []string{},
		height:   height,
		width:    width,
	}
}

func (t *Terminal) ClearScreen() {
	fmt.Print("\033[2J") // Clear entire screen
	fmt.Print("\033[H")  // Move cursor to top-left
}

func (t *Terminal) moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row, col)
}

func (t *Terminal) PrintMessage(msg string) {
	t.messages = append(t.messages, msg)
	// Crop messages if they exceed available screen space (height - 2 lines for error + prompt)
	maxMessages := t.height - 2
	if len(t.messages) > maxMessages {
		t.messages = t.messages[len(t.messages)-maxMessages:]
	}
	t.render()
}

func (t *Terminal) PrintError(msg string) {
	t.errorMessage = "\033[31m" + msg + "\033[0m" // red colored error message
	t.render()
}

func (t *Terminal) ClearError() {
	t.errorMessage = ""
	t.render()
}

func (t *Terminal) render() {
	t.ClearScreen()

	// Center the welcome message horizontally on the top row (row=1)
	welcome := "Welcome to S.L.A.M Client"
	col := max((t.width-len(welcome))/2, 1)

	// Move cursor to top row and centered column, then print welcome message
	t.moveCursor(1, col)
	fmt.Print(welcome)

	// Start printing messages from row 3 to leave space for the welcome banner
	msgStartRow := 3

	// Calculate max number of messages that fit in remaining space
	maxMessages := t.height - msgStartRow - 2 // reserving 2 lines for error message and prompt

	// Crop old messages if there are too many to fit the screen
	if len(t.messages) > maxMessages {
		t.messages = t.messages[len(t.messages)-maxMessages:]
	}

	// Print each message line by line starting at msgStartRow
	for i, msg := range t.messages {
		t.moveCursor(msgStartRow+i, 1)
		fmt.Print(msg)
	}

	// If an error message exists, print it one line above the prompt (second last line)
	if t.errorMessage != "" {
		t.moveCursor(t.height-1, 1)
		fmt.Print(t.errorMessage)
	}
}

func (t *Terminal) Prompt(label, nickname string) (string, error) {
	t.render()

	var promptLabel string
	if nickname != "" {
		promptLabel = "\033[31m[" + nickname + "]\033[0m " + label
	} else {
		promptLabel = "\033[34m[Unknown]\033[0m " + label
	}

	t.moveCursor(t.height, 1)
	fmt.Print(promptLabel)

	input, err := t.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}
