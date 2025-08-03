package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Terminal struct {
	reader *bufio.Reader
}

func NewTerminal() *Terminal {
	return &Terminal{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (t *Terminal) ShowWelcome() {
	fmt.Println("===================================")
	fmt.Println("      Welcome to S.L.A.M Client    ")
	fmt.Println("===================================")
	fmt.Println()
}

func (t *Terminal) Prompt(label string) (string, error) {
	fmt.Print(label)
	input, err := t.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func (t *Terminal) PrintError(msg string) {
	fmt.Println("\033[31m" + msg + "\033[0m")
}
