package compilation_engine

import (
	"encoding/xml"
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	symboltable "ny/nand2tetris/compiler/internal/symbol_table"
	Tokens "ny/nand2tetris/compiler/internal/tokens"
	vmwriter "ny/nand2tetris/compiler/internal/vm_writer"
)

var variableKindMemorySegmentMap = map[symboltable.VariableKind]vmwriter.Segment{
	symboltable.STATIC: vmwriter.STATIC,
	symboltable.FIELD:  vmwriter.THIS,
	symboltable.ARG:    vmwriter.ARGUMENT,
	symboltable.VAR:    vmwriter.LOCAL,
}

type subroutineInfo struct {
	name         string
	returnType   string
	functionType functionType
}
type functionType string

const (
	CONSTRUCTOR functionType = "constructor"
	FUNCTION    functionType = "function"
	METHOD      functionType = "method"
)

type ICompilationEngine interface {
	CompileClass() error
}

type CompilationEngine struct {
	tokens         []Tokens.IToken
	xmlEncoder     *xml.Encoder
	index          int
	componentStack *component.ComponentStack

	// set by CompileClass method
	// used by CompileSubroutine to write the function name
	className string

	// set by CompileSubroutine method
	subroutineInfo subroutineInfo

	symbolTable *symboltable.SymbolTable

	vmWriter vmwriter.IVMWriter

	labelCounterIf    int
	labelCounterWhile int
}

func New(tokens []Tokens.IToken, xmlEncoder *xml.Encoder, vmWriter vmwriter.IVMWriter) (ICompilationEngine, error) {

	return &CompilationEngine{
		tokens:            tokens,
		xmlEncoder:        xmlEncoder,
		index:             0,
		componentStack:    component.NewComponentStack(),
		symbolTable:       symboltable.New(),
		vmWriter:          vmWriter,
		subroutineInfo:    subroutineInfo{},
		labelCounterIf:    0,
		labelCounterWhile: 0,
	}, nil
}

func writeFlush(xmlEncoder *xml.Encoder, c *component.Component) error {
	xmlEncoder.Indent("", "  ")
	if err := xmlEncoder.Encode(c); err != nil {
		return fmt.Errorf("failed to encode XML: %w", err)
	}
	return nil
}
