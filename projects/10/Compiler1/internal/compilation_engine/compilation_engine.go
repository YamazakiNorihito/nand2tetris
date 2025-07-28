package compilation_engine

import (
	"encoding/xml"
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
	Tokens "ny/nand2tetris/compiler1/internal/tokens"
)

type ICompilationEngine interface {
	CompileClass() error
}

type CompilationEngine struct {
	tokens         []Tokens.IToken
	xmlEncoder     *xml.Encoder
	index          int
	componentStack *component.ComponentStack
}

func New(tokens []Tokens.IToken, xmlEncoder *xml.Encoder) (ICompilationEngine, error) {
	return &CompilationEngine{tokens: tokens, xmlEncoder: xmlEncoder, index: 0, componentStack: component.NewComponentStack()}, nil
}

func writeFlush(xmlEncoder *xml.Encoder, c *component.Component) error {
	xmlEncoder.Indent("", "  ")
	if err := xmlEncoder.Encode(c); err != nil {
		return fmt.Errorf("failed to encode XML: %w", err)
	}
	return nil
}
