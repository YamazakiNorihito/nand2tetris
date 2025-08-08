package jack_tokenizer

import (
	"errors"
	"io"
	"ny/nand2tetris/compiler/internal/token_patterns"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("forced read error")
}

func TestNewAnalysis(t *testing.T) {
	t.Run("should return an error when reader.Read returns an error", func(t *testing.T) {
		// Arrange
		reader := &errorReader{}

		// Act
		p, err := NewAnalysis(reader)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, p)
	})
	t.Run("should return no code lines for input with only single-line comments", func(t *testing.T) {
		// Arrange
		reader := strings.NewReader("// Empty input")

		// Act
		p, err := NewAnalysis(reader)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, p)

		assert.Empty(t, p.getCodeLines())
	})
	t.Run("should return no code lines for input with only block comments", func(t *testing.T) {
		// Arrange
		reader := strings.NewReader(`/* comment1
		block comment2
		block comment3 */`)

		// Act
		p, err := NewAnalysis(reader)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, p)

		assert.Empty(t, p.getCodeLines())
	})
	t.Run("should parse and return code lines correctly for input containing comments and class definition", func(t *testing.T) {
		// Arrange
		reader := strings.NewReader(`// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/10/ArrayTest/Main.jack

// (same as projects/9/Average/Main.jack)

/** Computes the average of a sequence of integers. */

class Main {
	function void main() { // main block
		/*
			process code
		*/
	return;
	}
}`)

		// Act
		p, err := NewAnalysis(reader)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, p)

		assert.Len(t, p.getCodeLines(), 5)
		assert.Equal(t, "class Main {", p.getCodeLines()[0])
		assert.Equal(t, "	function void main() { ", p.getCodeLines()[1])
		assert.Equal(t, "	return;", p.getCodeLines()[2])
		assert.Equal(t, "	}", p.getCodeLines()[3])
		assert.Equal(t, "}", p.getCodeLines()[4])
	})
	t.Run("should correctly ignore comments inside string constants and parse class definition", func(t *testing.T) {
		// Arrange
		reader := strings.NewReader(`
// single comment
/** This is a API Document block comment */
/**
		This is a API Document block comments
*/
/* This is a block comment */
/*
		This is a block comments
*/
class Main {
	function void main() { // main block
		do Output.printString("/* comments */");
		do Output.printString("/** comments */");
		do Output.printString("// comment ");
	return;
	}
}`)

		// Act
		p, err := NewAnalysis(reader)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, p)

		assert.Len(t, p.getCodeLines(), 8)
		assert.Equal(t, "class Main {", p.getCodeLines()[0])
		assert.Equal(t, "	function void main() { ", p.getCodeLines()[1])
		assert.Equal(t, "		do Output.printString(\"/* comments */\");", p.getCodeLines()[2])
		assert.Equal(t, "		do Output.printString(\"/** comments */\");", p.getCodeLines()[3])
		assert.Equal(t, "		do Output.printString(\"// comment \");", p.getCodeLines()[4])
		assert.Equal(t, "	return;", p.getCodeLines()[5])
		assert.Equal(t, "	}", p.getCodeLines()[6])
		assert.Equal(t, "}", p.getCodeLines()[7])
	})
}

func TestHasMoreTokens(t *testing.T) {
	t.Run("should return false for input containing only comments", func(t *testing.T) {
		reader := strings.NewReader("// only comment")
		p, err := NewAnalysis(reader)
		assert.NoError(t, err)
		assert.NotNil(t, p)
		assert.False(t, p.HasMoreTokens())
	})

	t.Run("should return true initially and false after advancing through all tokens", func(t *testing.T) {
		reader := strings.NewReader(`class`)
		p, err := NewAnalysis(reader)
		assert.NoError(t, err)
		assert.NotNil(t, p)

		assert.True(t, p.HasMoreTokens())
	})
}

func TestAdvance(t *testing.T) {
	t.Run("should advance through all tokens", func(t *testing.T) {
		const code = `class Main {
		function void main() {
			return;
		}
	}`

		p, err := NewAnalysis(strings.NewReader(code))
		assert.NoError(t, err)
		assert.NotNil(t, p)

		// class
		err = p.Advance()
		assert.NoError(t, err)
		// Main
		err = p.Advance()
		assert.NoError(t, err)
		// {
		err = p.Advance()
		assert.NoError(t, err)
		// function
		err = p.Advance()
		assert.NoError(t, err)
		// void
		err = p.Advance()
		assert.NoError(t, err)
		// main
		err = p.Advance()
		assert.NoError(t, err)
		// (
		err = p.Advance()
		assert.NoError(t, err)
		// )
		err = p.Advance()
		assert.NoError(t, err)
		// {
		err = p.Advance()
		assert.NoError(t, err)
		// return
		err = p.Advance()
		assert.NoError(t, err)
		// ;
		err = p.Advance()
		assert.NoError(t, err)
		// }
		err = p.Advance()
		assert.NoError(t, err)
		// }
		err = p.Advance()
		assert.NoError(t, err)

		err = p.Advance()
		assert.ErrorIs(t, err, io.EOF)
	})
	t.Run("should advance through all tokens1", func(t *testing.T) {
		const code = `"HOW MANY NUMBERS? "`

		p, err := NewAnalysis(strings.NewReader(code))
		assert.NoError(t, err)
		assert.NotNil(t, p)

		// class
		err = p.Advance()
		assert.NoError(t, err)
	})
}

func TestTokenType(t *testing.T) {
	t.Run("should return correct token type for individual tokens", func(t *testing.T) {
		cases := []struct {
			code         string
			expectedType TokenType
		}{
			{"class", KEYWORD},
			{"=", SYMBOL},
			{"myVar", IDENTIFIER},
			{"123", INT_CONST},
			{`"hello"`, STRING_CONST},
		}
		for _, c := range cases {
			p, err := NewAnalysis(strings.NewReader(c.code))
			assert.NoError(t, err)
			assert.NotNil(t, p)
			assert.True(t, p.HasMoreTokens())
			err = p.Advance()
			assert.NoError(t, err)
			assert.Equal(t, c.expectedType, p.TokenType())
		}
	})
}

func TestKeyword(t *testing.T) {
	t.Run("should correctly classify and return each keyword", func(t *testing.T) {
		// Arrange
		testCases := []struct {
			input    string
			expected token_patterns.KeywordType
		}{
			{"class", token_patterns.CLASS},
			{"constructor", token_patterns.CONSTRUCTOR},
			{"function", token_patterns.FUNCTION},
			{"method", token_patterns.METHOD},
			{"field", token_patterns.FIELD},
			{"static", token_patterns.STATIC},
			{"var", token_patterns.VAR},
			{"int", token_patterns.INT},
			{"char", token_patterns.CHAR},
			{"boolean", token_patterns.BOOLEAN},
			{"void", token_patterns.VOID},
			{"true", token_patterns.TRUE},
			{"false", token_patterns.FALSE},
			{"null", token_patterns.NULL},
			{"this", token_patterns.THIS},
			{"let", token_patterns.LET},
			{"do", token_patterns.DO},
			{"if", token_patterns.IF},
			{"else", token_patterns.ELSE},
			{"while", token_patterns.WHILE},
			{"return", token_patterns.RETURN},
		}
		for _, tc := range testCases {
			t.Run("should classify keyword: "+tc.input, func(t *testing.T) {
				p, err := NewAnalysis(strings.NewReader(tc.input))
				assert.NoError(t, err)
				assert.NotNil(t, p)
				assert.True(t, p.HasMoreTokens())
				err = p.Advance()
				assert.NoError(t, err)
				assert.Equal(t, KEYWORD, p.TokenType())
				assert.Equal(t, tc.expected, p.Keyword())
			})
		}
	})

	t.Run("should panic when Keyword() is called on non-keyword token", func(t *testing.T) {
		p, err := NewAnalysis(strings.NewReader("Main"))
		assert.NoError(t, err)
		assert.NotNil(t, p)
		p.Advance()
		assert.Panics(t, func() {
			_ = p.Keyword()
		})
	})
}

func TestSymbol(t *testing.T) {
	t.Run("should correctly return symbols and classify them", func(t *testing.T) {
		symbols := []string{"{", "}", "(", ")", "[", "]", ".", ",", ";", "+", "-", "*", "/", "&", "|", "<", ">", "=", "~"}
		code := strings.Join(symbols, " ")
		p, err := NewAnalysis(strings.NewReader(code))
		assert.NoError(t, err)
		assert.NotNil(t, p)

		for _, sym := range symbols {
			assert.True(t, p.HasMoreTokens())
			err := p.Advance()
			assert.NoError(t, err)
			assert.Equal(t, SYMBOL, p.TokenType())
			assert.Equal(t, sym, p.Symbol())
		}
		assert.False(t, p.HasMoreTokens())
	})

	t.Run("should panic when Symbol() is called on non-symbol token", func(t *testing.T) {
		p, err := NewAnalysis(strings.NewReader("class"))
		assert.NoError(t, err)
		assert.NotNil(t, p)
		p.Advance()
		assert.Panics(t, func() {
			_ = p.Symbol()
		})
	})
}

func TestIdentifier(t *testing.T) {
	t.Run("should correctly return identifiers", func(t *testing.T) {
		identifiers := []string{"foo", "bar", "baz"}
		code := strings.Join(identifiers, " ")
		p, err := NewAnalysis(strings.NewReader(code))
		assert.NoError(t, err)
		assert.NotNil(t, p)

		for _, ident := range identifiers {
			assert.True(t, p.HasMoreTokens())
			err := p.Advance()
			assert.NoError(t, err)
			assert.Equal(t, IDENTIFIER, p.TokenType())
			assert.Equal(t, ident, p.Identifier())
		}
		assert.False(t, p.HasMoreTokens())
	})

	t.Run("should panic when Identifier() is called on non-identifier token", func(t *testing.T) {
		p, err := NewAnalysis(strings.NewReader("class"))
		assert.NoError(t, err)
		assert.NotNil(t, p)
		p.Advance()
		assert.Panics(t, func() {
			_ = p.Identifier()
		})
	})
}

func TestIntVal(t *testing.T) {
	t.Run("should correctly parse integer constants", func(t *testing.T) {
		nums := []int{0, 1, 42, 32767}
		var parts []string
		for _, n := range nums {
			parts = append(parts, strconv.Itoa(n))
		}
		code := strings.Join(parts, " ")
		p, err := NewAnalysis(strings.NewReader(code))
		assert.NoError(t, err)
		assert.NotNil(t, p)

		for _, val := range nums {
			assert.True(t, p.HasMoreTokens())
			err := p.Advance()
			assert.NoError(t, err)
			assert.Equal(t, INT_CONST, p.TokenType())
			assert.Equal(t, val, p.IntVal())
		}
		assert.False(t, p.HasMoreTokens())
	})

	t.Run("should panic when IntVal() is called on non-integer token", func(t *testing.T) {
		p, err := NewAnalysis(strings.NewReader("class"))
		assert.NoError(t, err)
		assert.NotNil(t, p)
		p.Advance()
		assert.Panics(t, func() {
			_ = p.IntVal()
		})
	})
}
func TestStringVal(t *testing.T) {
	t.Run("should correctly parse string constants", func(t *testing.T) {
		stringsToTest := []string{`"a"`, `"hello world "`, `""`, `"Jack & Jill"`}
		code := strings.Join(stringsToTest, " ")
		p, err := NewAnalysis(strings.NewReader(code))
		assert.NoError(t, err)
		assert.NotNil(t, p)

		for _, raw := range stringsToTest {
			assert.True(t, p.HasMoreTokens())
			err := p.Advance()
			assert.NoError(t, err)
			assert.Equal(t, STRING_CONST, p.TokenType())
			expected := strings.Trim(raw, "\"")
			assert.Equal(t, expected, p.StringVal())
		}
		assert.False(t, p.HasMoreTokens())
	})

	t.Run("should panic when StringVal() is called on non-string token", func(t *testing.T) {
		p, err := NewAnalysis(strings.NewReader("123"))
		assert.NoError(t, err)
		assert.NotNil(t, p)
		p.Advance()
		assert.Panics(t, func() {
			_ = p.StringVal()
		})
	})
}
