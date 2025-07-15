package parser

import (
	"fmt"
	"io"
	"ny/nand2tetris/basicvmtranslator/internal/config"
	"strconv"
	"strings"
)

type IParser interface {
	HasMoreLines() bool
	Advance() (bool, error)
	CommandType() config.CommandType
	Arg1() string
	Arg2() int
	RawLineLength() int
	CurrentLineIndex() int
}

type parser struct {
	lines            []string
	currentLineIndex int
	currentLine      string
	rawLineLength    int

	commandType config.CommandType
	arg1        string
	arg2        int
}

func NewParser(r io.Reader) (IParser, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return &parser{
			lines:            []string{},
			currentLineIndex: -1,
			commandType:      config.C_UNKNOWN,
		}, nil
	}

	rawLines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	var lines []string
	for _, line := range rawLines {
		if commentIndex := strings.Index(line, "//"); commentIndex != -1 {
			line = line[:commentIndex]
		}
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}

	return &parser{
		lines:            lines,
		rawLineLength:    len(rawLines),
		commandType:      config.C_UNKNOWN,
		currentLineIndex: -1,
	}, nil
}

func (p *parser) HasMoreLines() bool {
	return p.currentLineIndex+1 < len(p.lines)
}

func (p *parser) Advance() (bool, error) {
	if !p.HasMoreLines() {
		return false, nil
	}
	p.resetInstructionParts()
	p.currentLineIndex++
	p.currentLine = p.lines[p.currentLineIndex]

	parts := strings.Fields(p.currentLine)
	command := parts[0]

	commandType, commandArgLength := config.GetCommandInfo(command)
	if commandType == config.C_UNKNOWN {
		return true, fmt.Errorf("unknown command: %s", command)
	}

	if len(parts)-1 != commandArgLength {
		return true, fmt.Errorf("invalid number of arguments for command %s: expected %d, got %d", command, commandArgLength, len(parts)-1)
	}

	p.commandType = commandType
	if commandArgLength >= 1 {
		p.arg1 = parts[1]
	} else {
		p.arg1 = command
	}

	if commandArgLength == 2 {
		val, err := strconv.Atoi(parts[2])
		if err != nil {
			return true, fmt.Errorf("invalid argument for %s: %w", command, err)
		}
		p.arg2 = val
	} else {
		p.arg2 = 0
	}
	return true, nil
}

func (p *parser) CommandType() config.CommandType {
	return p.commandType
}
func (p *parser) Arg1() string {
	return p.arg1
}
func (p *parser) Arg2() int {
	return p.arg2
}

func (p *parser) resetInstructionParts() {
	p.currentLine = ""
	p.commandType = config.C_UNKNOWN
	p.arg1 = ""
	p.arg2 = 0
}

func (p *parser) RawLineLength() int {
	return p.rawLineLength
}

func (p *parser) CurrentLineIndex() int {
	return p.currentLineIndex
}
