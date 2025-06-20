package parser

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	A_INSTRUCTION = "A_INSTRUCTION"
	C_INSTRUCTION = "C_INSTRUCTION"
	L_INSTRUCTION = "L_INSTRUCTION"
)

type IParser interface {
	CurrentLine() string
	Reset()
	HasMoreLines() bool
	Advance() error
	InstructionType() string
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}

type parser struct {
	lines            []string
	currentLineIndex int
	currentLine      string
	instructionType  string
	symbol           string
	dest             string
	comp             string
	jump             string
}

var (
	aCommandRegex = regexp.MustCompile(`^@([A-Za-z_.$:][A-Za-z0-9_.$:]*|\d+)$`)
	lCommandRegex = regexp.MustCompile(`^\(([A-Za-z_.$:][A-Za-z0-9_.$:]*)\)$`)
	cCommandRegex = regexp.MustCompile(
		`^(?:(?P<dest>[AMD]{1,3})=)?(?P<comp>0|1|-1|D|A|M|!D|!A|!M|-D|-A|-M|D\+1|A\+1|M\+1|D-1|A-1|M-1|D\+A|D\+M|D-A|D-M|A-D|M-D|D&A|D&M|D\|A|D\|M)(?:;(?P<jump>JGT|JEQ|JGE|JLT|JNE|JLE|JMP))?$`,
	)
)

func NewParser(assemblyFilePath string) (IParser, error) {
	content, err := os.ReadFile(assemblyFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': %w", assemblyFilePath, err)
	}

	if len(content) == 0 {
		return &parser{lines: []string{}}, nil
	}

	lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	return &parser{
		lines:            lines,
		currentLineIndex: 0,
	}, nil
}

func (p *parser) CurrentLine() string {
	return p.currentLine
}

func (p *parser) Reset() {
	p.currentLineIndex = 0
}

func (p *parser) HasMoreLines() bool {
	return p.currentLineIndex < len(p.lines)
}

func (p *parser) Advance() error {
	p.resetInstructionParts()
	for p.currentLineIndex < len(p.lines) {
		line := p.lines[p.currentLineIndex]
		p.currentLineIndex++

		line = strings.TrimSpace(line)

		if commentIndex := strings.Index(line, "//"); commentIndex >= 0 {
			line = strings.TrimSpace(line[:commentIndex])
		}

		if len(line) == 0 || strings.HasPrefix(line, "//") {
			continue
		}

		p.currentLine = line
		return p.parseInstruction(line)
	}
	return nil
}

func (p *parser) resetInstructionParts() {
	p.currentLine = ""
	p.instructionType = ""
	p.symbol = ""
	p.dest = ""
	p.comp = ""
	p.jump = ""
}

func (p *parser) parseInstruction(line string) error {
	if match := aCommandRegex.FindStringSubmatch(line); match != nil {
		p.instructionType = A_INSTRUCTION
		p.symbol = match[1]
		return nil
	}

	if match := lCommandRegex.FindStringSubmatch(line); match != nil {
		p.instructionType = L_INSTRUCTION
		p.symbol = match[1]
		return nil
	}

	if match := cCommandRegex.FindStringSubmatch(line); match != nil {
		p.instructionType = C_INSTRUCTION

		groupNames := cCommandRegex.SubexpNames()
		for i, name := range groupNames {
			if i == 0 || name == "" {
				continue
			}
			switch name {
			case "dest":
				p.dest = match[i]
			case "comp":
				p.comp = match[i]
			case "jump":
				p.jump = match[i]
			}
		}
		return nil
	}

	return fmt.Errorf("invalid instruction format on line %d: '%s'", p.currentLineIndex, line)
}

func (p *parser) InstructionType() string { return p.instructionType }

func (p *parser) Symbol() string { return p.symbol }

func (p *parser) Dest() string { return p.dest }

func (p *parser) Comp() string { return p.comp }

func (p *parser) Jump() string { return p.jump }
