package assembler

import (
	"fmt"
	"io"
	"ny/nand2tetris/fullassembler/internal/code"
	"ny/nand2tetris/fullassembler/internal/parser"
	"ny/nand2tetris/fullassembler/internal/symboltable"
	"strconv"
)

type Assembler struct {
	code   code.ICode
	parser parser.IParser
}

func NewAssembler(c code.ICode, p parser.IParser) *Assembler {
	return &Assembler{
		code:   c,
		parser: p,
	}
}

func (a *Assembler) Assemble(writer io.Writer) error {
	st, err := a.buildLabelSymbolTable()
	if err != nil {
		return err
	}

	return a.generateCode(writer, st)
}

func (a *Assembler) buildLabelSymbolTable() (symboltable.ISymbolTable, error) {
	a.parser.Reset()
	st := symboltable.NewSymbolTable()
	romAddress := 0
	for a.parser.HasMoreLines() {
		if err := a.parser.Advance(); err != nil {
			return nil, err
		}

		switch a.parser.InstructionType() {
		case parser.A_INSTRUCTION, parser.C_INSTRUCTION:
			romAddress++
		case parser.L_INSTRUCTION:
			symbol := a.parser.Symbol()
			if st.Contains(symbol) {
				return nil, fmt.Errorf("label '%s' is already defined", symbol)
			}
			st.AddEntry(symbol, romAddress)
		}
	}
	return st, nil
}

func (a *Assembler) generateCode(writer io.Writer, st symboltable.ISymbolTable) error {
	a.parser.Reset()
	for a.parser.HasMoreLines() {
		if err := a.parser.Advance(); err != nil {
			return err
		}

		instructionType := a.parser.InstructionType()

		switch instructionType {
		case parser.A_INSTRUCTION:
			symbol := a.parser.Symbol()
			var aRegisterValue int
			if value, err := strconv.Atoi(symbol); err == nil {
				// It's a number (e.g., @1234)
				aRegisterValue = value
			} else {
				// It's a symbol (e.g., @sum)
				if !st.Contains(symbol) {
					st.AddVariable(symbol)
				}
				aRegisterValue = st.GetAddress(symbol)
			}
			binary := fmt.Sprintf("%016b", uint16(aRegisterValue))
			if _, err := fmt.Fprintln(writer, binary); err != nil {
				return err
			}
		case parser.C_INSTRUCTION:
			binary := "111" + a.code.Comp(a.parser.Comp()) + a.code.Dest(a.parser.Dest()) + a.code.Jump(a.parser.Jump())
			if _, err := fmt.Fprintln(writer, binary); err != nil {
				return err
			}
		case parser.L_INSTRUCTION:
			// do nothing
		}
	}
	return nil
}
