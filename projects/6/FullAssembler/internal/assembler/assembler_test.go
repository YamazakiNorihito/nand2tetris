package assembler

import (
	"bytes"
	"ny/nand2tetris/fullassembler/internal/code"
	"ny/nand2tetris/fullassembler/internal/parser"
	"os"
	"strings"
	"testing"
)

type mockCode struct {
	destMap map[string]string
	compMap map[string]string
	jumpMap map[string]string
}

func (m *mockCode) Dest(mnemonic string) string { return m.destMap[mnemonic] }
func (m *mockCode) Comp(mnemonic string) string { return m.compMap[mnemonic] }
func (m *mockCode) Jump(mnemonic string) string { return m.jumpMap[mnemonic] }

type mockParser struct {
	hasMoreLinesSeq    []bool
	instructionTypeSeq []string
	symbolSeq          []string
	destSeq            []string
	compSeq            []string
	jumpSeq            []string

	hasMoreLinesIdx    int
	instructionTypeIdx int
	symbolIdx          int
	destIdx            int
	compIdx            int
	jumpIdx            int
}

func (m *mockParser) Reset() {
	m.hasMoreLinesIdx = 0
	m.instructionTypeIdx = 0
	m.symbolIdx = 0
	m.destIdx = 0
	m.compIdx = 0
	m.jumpIdx = 0
}

func (m *mockParser) HasMoreLines() bool {
	if m.hasMoreLinesIdx < len(m.hasMoreLinesSeq) {
		val := m.hasMoreLinesSeq[m.hasMoreLinesIdx]
		m.hasMoreLinesIdx++
		return val
	}
	return false
}

func (m *mockParser) Advance() error { return nil }

func (m *mockParser) InstructionType() string {
	if m.instructionTypeIdx < len(m.instructionTypeSeq) {
		val := m.instructionTypeSeq[m.instructionTypeIdx]
		m.instructionTypeIdx++
		return val
	}
	return ""
}

func (m *mockParser) Symbol() string {
	if m.symbolIdx < len(m.symbolSeq) {
		val := m.symbolSeq[m.symbolIdx]
		m.symbolIdx++
		return val
	}
	return ""
}

func (m *mockParser) Dest() string {
	if m.destIdx < len(m.destSeq) {
		val := m.destSeq[m.destIdx]
		m.destIdx++
		return val
	}
	return ""
}

func (m *mockParser) Comp() string {
	if m.compIdx < len(m.compSeq) {
		val := m.compSeq[m.compIdx]
		m.compIdx++
		return val
	}
	return ""
}

func (m *mockParser) Jump() string {
	if m.jumpIdx < len(m.jumpSeq) {
		val := m.jumpSeq[m.jumpIdx]
		m.jumpIdx++
		return val
	}
	return ""
}

func (m *mockParser) CurrentLine() string { return "" }

func TestAssemble_WritesAInstructionBinary(t *testing.T) {
	mockP := &mockParser{
		hasMoreLinesSeq:    []bool{true, false},
		instructionTypeSeq: []string{parser.A_INSTRUCTION},
		symbolSeq:          []string{"2"},
	}
	assembler := NewAssembler(&mockCode{}, mockP)
	var buf bytes.Buffer

	err := assembler.Assemble(&buf)
	if err != nil {
		t.Fatalf("Assemble() returned an unexpected error: %v", err)
	}

	expected := "0000000000000010\n"
	if buf.String() != expected {
		t.Errorf("Expected output %q, got %q", expected, buf.String())
	}
}

func TestAssemble_WritesCInstructionBinary(t *testing.T) {
	mockP := &mockParser{
		hasMoreLinesSeq:    []bool{true, false},
		instructionTypeSeq: []string{parser.C_INSTRUCTION},
		destSeq:            []string{"D"},
		compSeq:            []string{"A"},
		jumpSeq:            []string{"JGT"},
	}
	mockC := &mockCode{
		destMap: map[string]string{"D": "010"},
		compMap: map[string]string{"A": "0110000"},
		jumpMap: map[string]string{"JGT": "001"},
	}
	assembler := NewAssembler(mockC, mockP)
	var buf bytes.Buffer

	err := assembler.Assemble(&buf)
	if err != nil {
		t.Fatalf("Assemble() returned an unexpected error: %v", err)
	}

	expected := "1110110000010001\n"
	if buf.String() != expected {
		t.Errorf("Expected output %q, got %q", expected, buf.String())
	}
}

func TestAssemble_DoesNotWriteForLInstruction(t *testing.T) {
	mockP := &mockParser{
		hasMoreLinesSeq:    []bool{true, false},
		instructionTypeSeq: []string{parser.L_INSTRUCTION},
		symbolSeq:          []string{"LOOP"},
	}
	assembler := NewAssembler(&mockCode{}, mockP)
	var buf bytes.Buffer

	err := assembler.Assemble(&buf)
	if err != nil {
		t.Fatalf("Assemble() returned an unexpected error: %v", err)
	}

	if buf.Len() != 0 {
		t.Errorf("Expected empty output for L-instruction, but got %q", buf.String())
	}
}

func TestAssemble_ComputesR0Equals2Plus3(t *testing.T) {
	mockP := &mockParser{
		hasMoreLinesSeq: []bool{true, true, true, true, true, true, false},
		instructionTypeSeq: []string{
			parser.A_INSTRUCTION, // @2
			parser.C_INSTRUCTION, // D=A
			parser.A_INSTRUCTION, // @3
			parser.C_INSTRUCTION, // D=D+A
			parser.A_INSTRUCTION, // @0
			parser.C_INSTRUCTION, // M=D
		},
		symbolSeq: []string{"2", "3", "0"},
		destSeq:   []string{"D", "D", "M"},
		compSeq:   []string{"A", "D+A", "D"},
		jumpSeq:   []string{"", "", ""},
	}
	mockC := &mockCode{
		compMap: map[string]string{
			"A":   "0110000",
			"D+A": "0000010",
			"D":   "0001100",
		},
		destMap: map[string]string{
			"D": "010",
			"M": "001",
		},
		jumpMap: map[string]string{
			"": "000",
		},
	}
	assembler := NewAssembler(mockC, mockP)
	var buf bytes.Buffer

	err := assembler.Assemble(&buf)
	if err != nil {
		t.Fatalf("Assemble() returned an unexpected error: %v", err)
	}

	expected := strings.Join([]string{
		"0000000000000010",
		"1110110000010000",
		"0000000000000011",
		"1110000010010000",
		"0000000000000000",
		"1110001100001000",
	}, "\n") + "\n"

	if buf.String() != expected {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expected, buf.String())
	}
}

func TestAssemble_Rect(t *testing.T) {
	asmCode := `
// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/6/rect/Rect.asm

// Draws a rectangle at the top-left corner of the screen.
// The rectangle is 16 pixels wide and R0 pixels high.
// Usage: Before executing, put a value in R0.

   // If (R0 <= 0) goto END else n = R0
   @R0
   D=M
   @END
   D;JLE
   @n
   M=D
   // addr = base address of first screen row
   @SCREEN
   D=A
   @addr
   M=D
(LOOP)
   // RAM[addr] = -1
   @addr
   A=M
   M=-1
   // addr = base address of next screen row
   @addr
   D=M
   @32
   D=D+A
   @addr
   M=D
   // decrements n and loops
   @n
   MD=M-1
   @LOOP
   D;JGT
(END)
   @END
   0;JMP
`
	// Using the real parser and code generator for an integration-style test.
	tmpfile, err := os.CreateTemp("", "rect_*.asm")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(strings.TrimSpace(asmCode)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	p, err := parser.NewParser(tmpfile.Name())
	if err != nil {
		t.Fatalf("NewParser() failed: %v", err)
	}

	c := code.NewCode()
	assembler := NewAssembler(c, p)
	var buf bytes.Buffer

	if err := assembler.Assemble(&buf); err != nil {
		t.Fatalf("Assemble() returned an unexpected error: %v", err)
	}

	expected := strings.Join([]string{
		"0000000000000000", "1111110000010000", "0000000000010111", "1110001100000110",
		"0000000000010000", "1110001100001000", "0100000000000000", "1110110000010000",
		"0000000000010001", "1110001100001000", "0000000000010001", "1111110000100000",
		"1110111010001000", "0000000000010001", "1111110000010000", "0000000000100000",
		"1110000010010000", "0000000000010001", "1110001100001000", "0000000000010000",
		"1111110010011000", "0000000000001010", "1110001100000001", "0000000000010111",
		"1110101010000111",
	}, "\n") + "\n"

	if buf.String() != expected {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expected, buf.String())
	}
}
