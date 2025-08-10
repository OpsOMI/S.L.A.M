package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/message"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"golang.org/x/term"
)

type Terminal struct {
	reader          *bufio.Reader
	messages        []Messages
	rooms           []Rooms
	output          Notification
	label           string
	connected       bool
	currentRoomCode string
	height          int
	width           int
}

type Messages struct {
	SenderNickname string
	Content        string
}

type Notification struct {
	Code    string // Error, Information
	Message string
}

type Rooms struct {
	Code     string
	IsLocked bool
}

func NewTerminal() *Terminal {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24 // default fallback
		width = 80
	}

	return &Terminal{
		reader:    bufio.NewReader(os.Stdin),
		messages:  make([]Messages, 0),
		rooms:     make([]Rooms, 0),
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

	// Welcome message with colored "S.L.A.M"
	welcomePrefix := "Welcome to "
	welcomeColored := "\033[34mS.L.A.M\033[0m"
	welcomeSuffix := " Client"
	welcome := welcomePrefix + welcomeColored + welcomeSuffix
	col := max((t.width-len("Welcome to S.L.A.M Client"))/2, 1)
	t.moveCursor(1, col)
	fmt.Print(welcome)

	// Joined Room with green colored code
	if t.currentRoomCode != "" {
		maskedCode := t.currentRoomCode[:2] + "****"
		roomCodeMsg := "In Room " + maskedCode
		col = max((t.width-len(roomCodeMsg))/2, 1)
		t.moveCursor(2, col)
		fmt.Printf("\033[32m%s\033[0m", roomCodeMsg)
	}

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
	start := 0
	if len(t.messages) > maxMessages {
		start = len(t.messages) - maxMessages
	}
	visibleMessages := t.messages[start:]

	// 20 is the length of the room messages.
	msgLastColumn := t.width - 20
	roomsColStart := msgLastColumn + 5

	row := msgStartRow
	for _, message := range visibleMessages {
		msgLabel := message.SenderNickname + ": "
		msgContent := message.Content
		labelLen := len(msgLabel)

		firstLine := true
		for {
			var segment string
			if firstLine {
				end := min(len(msgContent), msgLastColumn-labelLen)
				segment = msgLabel + msgContent[:end]
				msgContent = msgContent[end:]
			} else {
				end := min(len(msgContent), msgLastColumn-labelLen)
				segment = strings.Repeat(" ", labelLen) + msgContent[:end]
				msgContent = msgContent[end:]
			}

			t.moveCursor(row, 1)
			fmt.Print(segment)
			if len(segment) < msgLastColumn {
				fmt.Print(strings.Repeat(" ", msgLastColumn-len(segment)))
			}

			row++
			firstLine = false
			if len(msgContent) == 0 {
				break
			}
		}
	}

	// Room
	row = msgStartRow
	for _, room := range t.rooms {
		var icon, color string
		if room.IsLocked {
			icon = "[LOCK]"
			color = "\033[31m"
		} else {
			icon = "[OPEN]"
			color = "\033[32m"
		}

		t.moveCursor(row, roomsColStart)
		fmt.Printf("%s%s %s\033[0m", color, icon, room.Code)
		row++
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
		t.PrintLabel()
	}
}

func (t *Terminal) Prompt() (string, error) {
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

func (t *Terminal) SetMessages(
	messages *message.MessagesReps,
) {
	t.messages = nil
	for _, m := range messages.Items {
		t.messages = append(t.messages, Messages{
			SenderNickname: m.SenderNickname,
			Content:        m.Content,
		})
	}
	t.Render()
}

func (t *Terminal) SetRooms(
	rooms *rooms.RoomsResp,
) {
	t.rooms = nil

	if rooms == nil {
		t.Render()
		return
	}

	for _, r := range rooms.Items {
		t.rooms = append(t.rooms, Rooms{
			Code:     r.Code,
			IsLocked: r.IsLocked,
		})
	}
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

func (t *Terminal) SetCurrentRoom(code string) {
	t.currentRoomCode = code
}

// Clears
func (t *Terminal) ClearScreen() {
	fmt.Print("\033[2J") // Clear entire screen
	fmt.Print("\033[H")  // Move cursor to top-left
}

func (t *Terminal) ClearLine(line int) {
	// Save current cursor position
	fmt.Print("\0337") // or \033[s

	// Move to the specified line and clear it
	fmt.Printf("\033[%d;1H", line)
	fmt.Print("\033[2K")

	// Redraw the input label (e.g., ":")
	t.PrintLabel()

	// Restore previous cursor position
	fmt.Print("\0338") // or \033[u
}

func (t *Terminal) ClearOutput() {
	t.output.Code = ""
	t.output.Message = ""
	t.Render()
}

// Prints
func (t *Terminal) Print(err error) {
	if appErr, ok := err.(*apperrors.AppError); ok {
		switch appErr.Code {
		case "Error":
			t.PrintError(appErr.Message)
		case "Notification":
			t.PrintNotification(appErr.Message)
		default:
			t.PrintError(appErr.Message) // default fallback
		}
	} else {
		t.PrintError(err.Error())
	}
}

func (t *Terminal) PrintError(msg string) {
	t.output.Message = msg
	t.output.Code = "error"
	t.Render()

	go func() {
		time.Sleep(1 * time.Second)
		t.ClearLine(t.height - 1)
	}()
}

func (t *Terminal) PrintNotification(msg string) {
	t.output.Message = msg
	t.output.Code = "info"
	t.Render()

	go func() {
		time.Sleep(1 * time.Second)
		t.ClearLine(t.height - 1)
	}()
}

func (t *Terminal) PrintLabel() {
	t.moveCursor(t.height, 1)
	fmt.Print(t.label)
}

func (t *Terminal) PrintMessage(nickname, content string) {
	t.messages = append(t.messages, Messages{
		SenderNickname: nickname,
		Content:        content,
	})
	t.Render()
}
