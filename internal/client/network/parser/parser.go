package parser

import (
	"errors"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

type IParser interface {
	Parse(input string) (Command, error)
}

type parser struct{}

func NewParser() IParser {
	return &parser{}
}

func (p *parser) Parse(input string) (Command, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return Command{}, errors.New("empty input")
	}

	if !strings.HasPrefix(input, "/") {
		return Command{}, errors.New("input does not start with /")
	}

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return Command{}, errors.New("no command found")
	}

	cmd := strings.TrimPrefix(parts[0], "/")
	args := []string{}
	if len(parts) > 1 {
		args = parts[1:]
	}

	return Command{Name: cmd, Args: args}, nil
}
