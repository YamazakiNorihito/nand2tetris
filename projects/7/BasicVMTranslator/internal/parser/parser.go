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
	Advance() error
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
	lineLength       int

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
		lineLength:       len(lines),
		commandType:      config.C_UNKNOWN,
		currentLineIndex: -1,
	}, nil
}

func (p *parser) HasMoreLines() bool {
	return p.currentLineIndex+1 < len(p.lines)
}

func (p *parser) Advance() error {
	if !p.HasMoreLines() {
		return fmt.Errorf("no more lines")
	}
	p.resetInstructionParts()
	p.currentLineIndex++
	p.currentLine = p.lines[p.currentLineIndex]

	parts := strings.Fields(p.currentLine)
	if len(parts) == 0 {
		return fmt.Errorf("empty command line")
	}

	command := parts[0]

	if config.IsArithmeticCommand(command) {
		p.commandType = config.C_ARITHMETIC
		p.arg1 = command
		return nil
	}

	if len(parts) != 3 {
		return fmt.Errorf("invalid command format for non-arithmetic command: %s", p.currentLine)
	}

	p.arg1 = parts[1]
	val, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("invalid argument for %s: %w", command, err)
	}
	p.arg2 = val

	switch command {
	case "push":
		p.commandType = config.C_PUSH
	case "pop":
		p.commandType = config.C_POP
	default:
		return fmt.Errorf("unknown command: %s", command)
	}

	return nil
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
