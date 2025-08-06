package token_patterns

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const demoCode = `class Main {
	function void main() {
		var Array a;
		var int length;
		var int i, sum;

	let length = Keyboard.readInt("HOW MANY NUMBERS? ");
	let a = Array.new(length);
	let i = 0;

	while (i < length) {
		let a[i] = Keyboard.readInt("ENTER THE NEXT NUMBER: ");
		let i = i + 1;
	}
	
	let i = 0;
	let sum = 0;
	
	while (i < length) {
		let sum = sum + a[i];
		let i = i + 1;
	}
	
	do Output.printString("THE AVERAGE IS: ");
	do Output.printInt(sum / length);
	do Output.println();
	
	return;
	}
}`

// TestTokenPatternOnMain verifies that tokenPattern correctly extracts the expected tokens
// from the provided Main.jack program after stripping comments.
func TestTokenPatternOnMain(t *testing.T) {
	tokens := TokenPattern.FindAllString(demoCode, -1)

	expected := []string{
		"class", "Main", "{",
		"function", "void", "main", "(", ")", "{",
		"var", "Array", "a", ";",
		"var", "int", "length", ";",
		"var", "int", "i", ",", "sum", ";",
		"let", "length", "=", "Keyboard", ".", "readInt", "(", "\"HOW MANY NUMBERS? \"", ")", ";",
		"let", "a", "=", "Array", ".", "new", "(", "length", ")", ";",
		"let", "i", "=", "0", ";",
		"while", "(", "i", "<", "length", ")", "{",
		"let", "a", "[", "i", "]", "=", "Keyboard", ".", "readInt", "(", "\"ENTER THE NEXT NUMBER: \"", ")", ";",
		"let", "i", "=", "i", "+", "1", ";",
		"}",
		"let", "i", "=", "0", ";",
		"let", "sum", "=", "0", ";",
		"while", "(", "i", "<", "length", ")", "{",
		"let", "sum", "=", "sum", "+", "a", "[", "i", "]", ";",
		"let", "i", "=", "i", "+", "1", ";",
		"}",
		"do", "Output", ".", "printString", "(", "\"THE AVERAGE IS: \"", ")", ";",
		"do", "Output", ".", "printInt", "(", "sum", "/", "length", ")", ";",
		"do", "Output", ".", "println", "(", ")", ";",
		"return", ";",
		"}", "}",
	}

	if len(tokens) != len(expected) {
		t.Fatalf("token count mismatch: got %d, expected %d", len(tokens), len(expected))
	}
	for i, tok := range tokens {
		if tok != expected[i] {
			t.Errorf("token mismatch at index %d: got %q, expected %q", i, tok, expected[i])
		}
	}
}

// TestKeywordPattern ensures that the keyword regular expression only matches valid Jack keywords.

func TestKeywordPattern(t *testing.T) {
	t.Run("should match valid keywords", func(t *testing.T) {
		samples := []string{"class", "constructor", "function", "method", "field", "static", "var", "int", "char", "boolean", "void", "true", "false", "null", "this", "let", "do", "if", "else", "while", "return"}
		for _, s := range samples {
			if !KeywordPattern.MatchString(s) {
				t.Errorf("keywordPattern should match %q", s)
			}
		}
	})
	t.Run("should not match invalid keywords", func(t *testing.T) {
		invalid := []string{"Class", "returnValue", "func", "intt", "whilee", "classtest", "double", ""}
		for _, s := range invalid {
			if KeywordPattern.MatchString(s) {
				t.Errorf("keywordPattern should not match %q", s)
			}
		}
	})
}

// TestSymbolPattern verifies that the symbol regular expression matches only Jack symbols.

func TestSymbolPattern(t *testing.T) {
	t.Run("should match valid symbols", func(t *testing.T) {
		symbols := []string{"{", "}", "(", ")", "[", "]", ".", ",", ";", "+", "-", "*", "/", "&", "|", "<", ">", "=", "~"}
		for _, sym := range symbols {
			if !SymbolPattern.MatchString(sym) {
				t.Errorf("symbolPattern should match %q", sym)
			}
		}
	})
	t.Run("should not match invalid symbols", func(t *testing.T) {
		invalid := []string{"a", "?", "1", "abc", "", " "}
		for _, s := range invalid {
			if SymbolPattern.MatchString(s) {
				t.Errorf("symbolPattern should not match %q", s)
			}
		}
	})
}

// TestIntegerConstantPattern checks the integer constant regular expression for valid and invalid ranges.
func TestIntegerConstantPattern(t *testing.T) {
	t.Run("should return correct token type for individual tokens11", func(t *testing.T) {
		tests := []struct {
			input    string
			expected int
		}{
			{"0", 0},
			{"9", 9},
			{"10", 10},
			{"3276", 3276},
			{"9999", 9999},
			{"10000", 10000},
			{"29999", 29999},
			{"32767", 32767},
		}
		for _, test := range tests {
			act := IntegerConstantPattern.FindString(test.input)
			val, err := strconv.Atoi(act)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, val)
		}
	})

	t.Run("should return correct token type for individual tokens22", func(t *testing.T) {
		tests := []struct {
			input string
		}{
			{"a0"},
			{"9a"},
			{"32768"},
			{"-1"},
		}
		for _, test := range tests {
			act := IntegerConstantPattern.MatchString(test.input)
			assert.False(t, act)
		}
	})
}

// TestStringConstantPattern tests matching of string constants and rejection of invalid strings.
func TestStringConstantPattern(t *testing.T) {
	t.Run("should match valid string constants", func(t *testing.T) {
		valid := []string{"\"\"", "\"hello\"", "\"Jack language\""}
		for _, v := range valid {
			if !StringConstantPattern.MatchString(v) {
				t.Errorf("stringConstantPattern should match %q", v)
			}
		}
	})
	t.Run("should not match invalid string constants", func(t *testing.T) {
		invalid := []string{"\"unterminated", "\"contains\nnewline\"", "hello", "\""}
		for _, v := range invalid {
			if StringConstantPattern.MatchString(v) {
				t.Errorf("stringConstantPattern should not match %q", v)
			}
		}
	})
}

// TestIdentifierPattern asserts that identifier regular expression matches valid identifiers and rejects invalid ones.

func TestIdentifierPattern(t *testing.T) {
	t.Run("should match valid identifiers", func(t *testing.T) {
		valid := []string{"x", "foo", "Bar", "_underscore", "a1b2"}
		for _, v := range valid {
			if !IdentifierPattern.MatchString(v) {
				t.Errorf("identifierPattern should match %q", v)
			}
		}
	})
	t.Run("should not match invalid identifiers", func(t *testing.T) {
		invalid := []string{"1abc", "9", "1_", ""}
		for _, v := range invalid {
			if IdentifierPattern.MatchString(v) {
				t.Errorf("identifierPattern should not match %q", v)
			}
		}
	})
}

func TestTokenSplitAndKeepDelimiters(t *testing.T) {

	t.Run("should split simple let statement with spaces and delimiters", func(t *testing.T) {
		input := `let x = 10;`
		expected := []string{"let", " ", "x", " ", "=", " ", "10", ";"}

		tokens := TokenSplitAndKeepDelimiters(input)
		if len(tokens) != len(expected) {
			t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
		}
		for i, token := range tokens {
			if token != expected[i] {
				t.Errorf("token mismatch at index %d: got %q, expected %q", i, token, expected[i])
			}
		}
	})

	t.Run("should split simple let statement with spaces and delimiters1", func(t *testing.T) {
		input := `Main.double(2)`
		expected := []string{"Main", ".", "double", "(", "2", ")"}

		tokens := TokenSplitAndKeepDelimiters(input)
		if len(tokens) != len(expected) {
			t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
		}
		for i, token := range tokens {
			if token != expected[i] {
				t.Errorf("token mismatch at index %d: got %q, expected %q", i, token, expected[i])
			}
		}
	})

	t.Run("should split demoCode with all delimiters and whitespace preserved", func(t *testing.T) {

		expected := []string{
			"class", " ", "Main", " ", "{",
			"\n\t", "function", " ", "void", " ", "main", "(", ")", " ", "{",
			"\n\t\t", "var", " ", "Array", " ", "a", ";",
			"\n\t\t", "var", " ", "int", " ", "length", ";",
			"\n\t\t", "var", " ", "int", " ", "i", ",", " ", "sum", ";",
			"\n\n\t", "let", " ", "length", " ", "=", " ", "Keyboard", ".", "readInt", "(", "\"HOW MANY NUMBERS? \"", ")", ";",
			"\n\t", "let", " ", "a", " ", "=", " ", "Array", ".", "new", "(", "length", ")", ";",
			"\n\t", "let", " ", "i", " ", "=", " ", "0", ";",
			"\n\n\t", "while", " ", "(", "i", " ", "<", " ", "length", ")", " ", "{",
			"\n\t\t", "let", " ", "a", "[", "i", "]", " ", "=", " ", "Keyboard", ".", "readInt", "(", "\"ENTER THE NEXT NUMBER: \"", ")", ";",
			"\n\t\t", "let", " ", "i", " ", "=", " ", "i", " ", "+", " ", "1", ";",
			"\n\t", "}",
			"\n\t\n\t", "let", " ", "i", " ", "=", " ", "0", ";",
			"\n\t", "let", " ", "sum", " ", "=", " ", "0", ";",
			"\n\t\n\t", "while", " ", "(", "i", " ", "<", " ", "length", ")", " ", "{",
			"\n\t\t", "let", " ", "sum", " ", "=", " ", "sum", " ", "+", " ", "a", "[", "i", "]", ";",
			"\n\t\t", "let", " ", "i", " ", "=", " ", "i", " ", "+", " ", "1", ";",
			"\n\t", "}",
			"\n\t\n\t", "do", " ", "Output", ".", "printString", "(", "\"THE AVERAGE IS: \"", ")", ";",
			"\n\t", "do", " ", "Output", ".", "printInt", "(", "sum", " ", "/", " ", "length", ")", ";",
			"\n\t", "do", " ", "Output", ".", "println", "(", ")", ";",
			"\n\t\n\t", "return", ";",
			"\n\t", "}",
			"\n", "}",
		}

		tokens := TokenSplitAndKeepDelimiters(demoCode)
		if len(tokens) != len(expected) {
			t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
		}
		for i, token := range tokens {
			if token != expected[i] {
				t.Errorf("token mismatch at index %d: got %q, expected %q", i, token, expected[i])
			}
		}
	})
}
