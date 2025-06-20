package parser

import (
	"os"
	"strings"
	"testing"
)

func createTempFile(t *testing.T, lines []string) (string, func()) {
	t.Helper()
	content := strings.Join(lines, "\n")
	tmpfile, err := os.CreateTemp("", "test_assembly_*.asm")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpfile.Close()

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	return tmpfile.Name(), func() {
		os.Remove(tmpfile.Name())
	}
}

func TestNewParser_WithAbsolutePath_ReturnsParser(t *testing.T) {
	path, cleanup := createTempFile(t, []string{"@2", "D=A"})
	defer cleanup()

	p, err := NewParser(path)
	if err != nil {
		t.Fatalf("NewParser returned an error: %v", err)
	}
	if p == nil {
		t.Fatal("NewParser returned nil parser")
	}
}

func TestNewParser_WithNonExistentPath_ReturnsError(t *testing.T) {
	p, err := NewParser("non_existent_file.asm")
	if err == nil {
		t.Fatal("NewParser did not return an error for a non-existent file")
	}
	if p != nil {
		t.Fatal("NewParser returned a non-nil parser for a non-existent file")
	}
}

func TestHasMoreLines_WithOnlyEmptyLines_ReturnsFalse(t *testing.T) {
	path, cleanup := createTempFile(t, []string{})
	defer cleanup()

	p, err := NewParser(path)
	if err != nil {
		t.Fatalf("NewParser returned an error: %v", err)
	}

	if p.HasMoreLines() {
		t.Error("HasMoreLines returned true for empty file")
	}
}

func TestHasMoreLines_WithValidInstructionLines_ReturnsTrue(t *testing.T) {
	lines := []string{
		"@10",
		"D=A",
	}
	path, cleanup := createTempFile(t, lines)
	defer cleanup()

	p, err := NewParser(path)
	if err != nil {
		t.Fatalf("NewParser returned an error: %v", err)
	}

	if !p.HasMoreLines() {
		t.Error("HasMoreLines returned false for file with instructions")
	}
}

func TestAdvance_SetsCurrentLineToFirstInstruction(t *testing.T) {
	tests := []struct {
		name        string
		inputLines  []string
		expected    string
		expectError bool
	}{
		{"Skip comments and empty lines", []string{"// comment", "", "   ", "@1", "D=M"}, "@1", false},
		{"Single instruction", []string{"@2"}, "@2", false},
		{"Only comments/empty, no instruction", []string{"   // only comment", "   ", ""}, "", false},
		{"Only comments, no instruction", []string{"   // only comment1", "   // only comment2"}, "", false},
		{"Inline comment", []string{"   // only comment1", "@3   // only comment2"}, "@3", false},
		{"Invalid instruction", []string{"INVALID_LINE"}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup := createTempFile(t, tt.inputLines)
			defer cleanup()

			p, err := NewParser(path)
			if err != nil {
				t.Fatalf("NewParser returned an error: %v", err)
			}

			err = p.Advance()
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected an error for invalid instruction, but got none")
				}
				return
			} else if err != nil {
				t.Fatalf("Advance returned an unexpected error: %v", err)
			}

			if p.CurrentLine() != tt.expected {
				t.Errorf("Expected current line '%s', got '%s'", tt.expected, p.CurrentLine())
			}
		})
	}
}

func TestAdvance_Twice_SetsCurrentLineToSecondInstruction(t *testing.T) {
	tests := []struct {
		name           string
		inputLines     []string
		expectedFirst  string
		expectedSecond string
	}{
		{"Skip comments, then two instructions", []string{"// comment", "", "   ", "@1", "D=M"}, "@1", "D=M"},
		{"Two A-instructions", []string{"@2", "@3"}, "@2", "@3"},
		{"Comments then instruction, then empty", []string{"   // only comment", "   ", "@5"}, "@5", ""},
		{"Inline comment, then instruction, then comment", []string{"@7   // comment", "M=D", "// end"}, "@7", "M=D"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup := createTempFile(t, tt.inputLines)
			defer cleanup()

			p, err := NewParser(path)
			if err != nil {
				t.Fatalf("NewParser returned an error: %v", err)
			}

			err = p.Advance()
			if err != nil {
				t.Fatalf("First Advance returned an unexpected error: %v", err)
			}
			if p.CurrentLine() != tt.expectedFirst {
				t.Errorf("First Advance: Expected current line '%s', got '%s'", tt.expectedFirst, p.CurrentLine())
			}

			err = p.Advance()
			if err != nil {
				t.Fatalf("Second Advance returned an unexpected error: %v", err)
			}
			if p.CurrentLine() != tt.expectedSecond {
				t.Errorf("Second Advance: Expected current line '%s', got '%s'", tt.expectedSecond, p.CurrentLine())
			}
		})
	}
}

func TestInstructionType_ReturnsCorrectInstructionType(t *testing.T) {
	tests := []struct {
		name         string
		inputLines   []string
		expectedType string
	}{
		{"A-instruction", []string{"@2"}, A_INSTRUCTION},
		{"A-instruction with spaces and comment", []string{"   @15   // set value"}, A_INSTRUCTION},
		{"L-instruction", []string{"(LOOP)"}, L_INSTRUCTION},
		{"L-instruction with spaces", []string{"   (END)   "}, L_INSTRUCTION},
		{"C-instruction D=M", []string{"D=M"}, C_INSTRUCTION},
		{"C-instruction M=D;JGT", []string{"M=D;JGT"}, C_INSTRUCTION},
		{"Pure comment/empty lines", []string{"// pure comment", "", "   "}, ""},
		{"Inline comment A-instruction", []string{"// comment", "@3 // inline comment"}, A_INSTRUCTION},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup := createTempFile(t, tt.inputLines)
			defer cleanup()

			p, err := NewParser(path)
			if err != nil {
				t.Fatalf("NewParser returned an error: %v", err)
			}

			err = p.Advance()
			if err != nil && tt.expectedType != "" {
				t.Fatalf("Advance returned an unexpected error: %v", err)
			}

			if p.InstructionType() != tt.expectedType {
				t.Errorf("Expected instruction type '%s', got '%s'", tt.expectedType, p.InstructionType())
			}
		})
	}
}

func TestSymbol_ReturnsCorrectSymbol(t *testing.T) {
	tests := []struct {
		name           string
		inputLines     []string
		expectedSymbol string
	}{
		{"A-instruction decimal", []string{"@21"}, "21"},
		{"A-instruction symbol", []string{"@LOOP"}, "LOOP"},
		{"A-instruction with spaces and comment", []string{"   @R2   // comment"}, "R2"},
		{"L-instruction", []string{"(END)"}, "END"},
		{"L-instruction with spaces", []string{"   (START)   "}, "START"},
		{"A-instruction with inline comment", []string{"// comment", "@foo // inline"}, "foo"},
		{"C-instruction (no symbol)", []string{"D=M"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup := createTempFile(t, tt.inputLines)
			defer cleanup()

			p, err := NewParser(path)
			if err != nil {
				t.Fatalf("NewParser returned an error: %v", err)
			}

			err = p.Advance()
			if err != nil && tt.expectedSymbol != "" {
				t.Fatalf("Advance returned an unexpected error: %v", err)
			}

			if p.Symbol() != tt.expectedSymbol {
				t.Errorf("Expected symbol '%s', got '%s'", tt.expectedSymbol, p.Symbol())
			}
		})
	}
}

func TestDest_ReturnsCorrectDest(t *testing.T) {
	tests := []struct {
		name         string
		inputLines   []string
		expectedDest string
	}{
		{"D=M", []string{"D=M"}, "D"},
		{"M=D", []string{"M=D"}, "M"},
		{"MD=D+1", []string{"MD=D+1"}, "MD"},
		{"A=M-1", []string{"A=M-1"}, "A"},
		{"D;JGT (no dest)", []string{"D;JGT"}, ""},
		{"D+1 (no dest)", []string{"D+1"}, ""},
		{"A-instruction (no dest)", []string{"@2"}, ""},
		{"L-instruction (no dest)", []string{"(LOOP)"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup := createTempFile(t, tt.inputLines)
			defer cleanup()

			p, err := NewParser(path)
			if err != nil {
				t.Fatalf("NewParser returned an error: %v", err)
			}

			err = p.Advance()
			if err != nil && tt.expectedDest != "" {
				t.Fatalf("Advance returned an unexpected error: %v", err)
			}

			if p.Dest() != tt.expectedDest {
				t.Errorf("Expected dest '%s', got '%s'", tt.expectedDest, p.Dest())
			}
		})
	}
}

func TestComp_ReturnsCorrectComp(t *testing.T) {
	tests := []struct {
		name         string
		inputLines   []string
		expectedComp string
	}{
		{"D=M", []string{"D=M"}, "M"},
		{"M=D", []string{"M=D"}, "D"},
		{"MD=D+1", []string{"MD=D+1"}, "D+1"},
		{"A=M-1", []string{"A=M-1"}, "M-1"},
		{"D+1", []string{"D+1"}, "D+1"},
		{"D;JGT", []string{"D;JGT"}, "D"},
		{"0;JMP", []string{"0;JMP"}, "0"},
		{"A-instruction (no comp)", []string{"@2"}, ""},
		{"L-instruction (no comp)", []string{"(LOOP)"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup := createTempFile(t, tt.inputLines)
			defer cleanup()

			p, err := NewParser(path)
			if err != nil {
				t.Fatalf("NewParser returned an error: %v", err)
			}

			err = p.Advance()
			if err != nil && tt.expectedComp != "" {
				t.Fatalf("Advance returned an unexpected error: %v", err)
			}

			if p.Comp() != tt.expectedComp {
				t.Errorf("Expected comp '%s', got '%s'", tt.expectedComp, p.Comp())
			}
		})
	}
}

func TestJump_ReturnsCorrectJump(t *testing.T) {
	tests := []struct {
		name         string
		inputLines   []string
		expectedJump string
	}{
		{"D=M (no jump)", []string{"D=M"}, ""},
		{"M=D (no jump)", []string{"M=D"}, ""},
		{"MD=D+1 (no jump)", []string{"MD=D+1"}, ""},
		{"A=M-1 (no jump)", []string{"A=M-1"}, ""},
		{"D+1 (no jump)", []string{"D+1"}, ""},
		{"D;JGT", []string{"D;JGT"}, "JGT"},
		{"0;JMP", []string{"0;JMP"}, "JMP"},
		{"A-instruction (no jump)", []string{"@2"}, ""},
		{"L-instruction (no jump)", []string{"(LOOP)"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cleanup := createTempFile(t, tt.inputLines)
			defer cleanup()

			p, err := NewParser(path)
			if err != nil {
				t.Fatalf("NewParser returned an error: %v", err)
			}

			err = p.Advance()
			if err != nil && tt.expectedJump != "" {
				t.Fatalf("Advance returned an unexpected error: %v", err)
			}

			if p.Jump() != tt.expectedJump {
				t.Errorf("Expected jump '%s', got '%s'", tt.expectedJump, p.Jump())
			}
		})
	}
}
