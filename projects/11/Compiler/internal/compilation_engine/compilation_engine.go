package compilation_engine

import (
	"encoding/xml"
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	symboltable "ny/nand2tetris/compiler/internal/symbol_table"
	Tokens "ny/nand2tetris/compiler/internal/tokens"
)

type ICompilationEngine interface {
	CompileClass() error
}

type CompilationEngine struct {
	tokens         []Tokens.IToken
	xmlEncoder     *xml.Encoder
	index          int
	componentStack *component.ComponentStack

	symbolTable *symboltable.SymbolTable
}

func New(tokens []Tokens.IToken, xmlEncoder *xml.Encoder) (ICompilationEngine, error) {
	return &CompilationEngine{
		tokens:         tokens,
		xmlEncoder:     xmlEncoder,
		index:          0,
		componentStack: component.NewComponentStack(),
		symbolTable:    symboltable.New(),
	}, nil
}

func writeFlush(xmlEncoder *xml.Encoder, c *component.Component) error {
	xmlEncoder.Indent("", "  ")
	if err := xmlEncoder.Encode(c); err != nil {
		return fmt.Errorf("failed to encode XML: %w", err)
	}
	return nil
}
