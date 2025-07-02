package parser

import (
	"errors"
	"ny/nand2tetris/basicvmtranslator/internal/config"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("forced read error")
}

func TestNewParser(t *testing.T) {

	t.Run("should return an error when reading from the input fails", func(t *testing.T) {
		// Arrange
		reader := &errorReader{}

		// Act
		p, err := NewParser(reader)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, p)
	})

	t.Run("should create a parser with no lines for empty input", func(t *testing.T) {
		// Arrange
		reader := strings.NewReader("")

		// Act
		p, err := NewParser(reader)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, p)
		// NewParserが正しく空の行スライスを生成したことを確認
		assert.Equal(t, []string{}, p.(*parser).lines)
	})

	t.Run("should correctly parse valid input, ignoring comments and empty lines", func(t *testing.T) {
		// Arrange
		input := `
// This is a full-line comment.
push constant 10 // This is an inline comment.

pop local 0      // Another inline comment.
`
		reader := strings.NewReader(input)
		expectedLines := []string{
			"push constant 10",
			"pop local 0",
		}

		// Act
		p, err := NewParser(reader)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, p)
		// コメントと空行が正しく除去されたことを確認
		assert.Equal(t, expectedLines, p.(*parser).lines)
	})
}

func TestParser_HasMoreLines(t *testing.T) {
	t.Run("should correctly report line availability", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant 1\npop local 2"))
		assert.NoError(t, err)

		// Act & Assert
		assert.True(t, p.HasMoreLines(), "should be true before first advance")
		err = p.Advance()
		assert.NoError(t, err)
		assert.True(t, p.HasMoreLines(), "should be true after first advance")
		err = p.Advance()
		assert.NoError(t, err)
		assert.False(t, p.HasMoreLines(), "should be false after second advance")
	})

	t.Run("should return false for empty input", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader(""))
		assert.NoError(t, err)

		// Act & Assert
		assert.False(t, p.HasMoreLines())
	})
}

func TestParser_Advance(t *testing.T) {
	t.Run("should parse arithmetic command", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("add"))

		// Act
		advanceErr := p.Advance()

		// Assert
		assert.NoError(t, err)

		assert.NoError(t, advanceErr)
		assert.Equal(t, 0, p.CurrentLineIndex())
	})

	t.Run("should parse push command", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant 7"))

		// Act
		advanceErr := p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.NoError(t, advanceErr)
		assert.Equal(t, 0, p.CurrentLineIndex())
	})

	t.Run("should return error for invalid command format", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant"))

		// Act
		advanceErr := p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Error(t, advanceErr)
	})

	t.Run("should return error for unknown command", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("foo bar baz"))

		// Act
		advanceErr := p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Error(t, advanceErr)
	})

	t.Run("should return error for invalid argument", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant notanumber"))

		// Act
		advanceErr := p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Error(t, advanceErr)
	})

	t.Run("should return error when no more lines", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("add"))

		// Act
		advanceErr1 := p.Advance()
		advanceErr2 := p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.NoError(t, advanceErr1)
		assert.Error(t, advanceErr2)
	})

	t.Run("should advance to first line (multi-line input)", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader(`
// This is a full-line comment.
push constant 10 // This is an inline comment.

pop local 0      // Another inline comment.
`))

		// Act
		advanceErr1 := p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.NoError(t, advanceErr1)
		assert.Equal(t, 0, p.CurrentLineIndex())
	})

	t.Run("should advance to second line and error on third advance (multi-line input)", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader(`
// This is a full-line comment.
push constant 10 // This is an inline comment.

pop local 0      // Another inline comment.
`))

		// Act
		advanceErr1 := p.Advance()
		advanceErr2 := p.Advance()
		advanceErr3 := p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.NoError(t, advanceErr1)
		assert.NoError(t, advanceErr2)
		assert.Error(t, advanceErr3)
		assert.Equal(t, 1, p.CurrentLineIndex())
	})
}

func TestParser_CommandType(t *testing.T) {
	t.Run("should return -1 before first advance", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant 7"))
		assert.NoError(t, err)

		// Act
		cmdType := p.CommandType()

		// Assert
		assert.Equal(t, config.C_UNKNOWN, cmdType)
	})

	t.Run("should return correct command type after advance", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant 7\nadd"))
		assert.NoError(t, err)

		// Act & Assert for C_PUSH
		err = p.Advance()
		assert.NoError(t, err)
		assert.Equal(t, config.C_PUSH, p.CommandType())

		// Act & Assert for C_ARITHMETIC
		err = p.Advance()
		assert.NoError(t, err)
		assert.Equal(t, config.C_ARITHMETIC, p.CommandType())
	})

	t.Run("should return -1 after failed advance", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("invalid command"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.Error(t, err)
		assert.Equal(t, config.C_UNKNOWN, p.CommandType())
	})
}

func TestParser_Arg1(t *testing.T) {
	t.Run("should return empty string before first advance", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant 7"))
		assert.NoError(t, err)

		// Act & Assert
		assert.Equal(t, "", p.Arg1())
	})

	t.Run("should return command for C_ARITHMETIC", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("add"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "add", p.Arg1())
	})

	t.Run("should return first argument for C_PUSH", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant 7"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "constant", p.Arg1())
	})

	t.Run("should return first argument for C_POP", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("pop local 0"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "local", p.Arg1())
	})

	t.Run("should return empty string after failed advance", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("invalid command"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "", p.Arg1())
	})
}

func TestParser_Arg2(t *testing.T) {
	t.Run("should return 0 before first advance", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant 7"))
		assert.NoError(t, err)

		// Act & Assert
		assert.Equal(t, 0, p.Arg2())
	})

	t.Run("should return second argument for C_PUSH", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant 7"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 7, p.Arg2())
	})

	t.Run("should return second argument for C_POP", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("pop local 0"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 0, p.Arg2())
	})

	t.Run("should return 0 for C_ARITHMETIC", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("add"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 0, p.Arg2())
	})

	t.Run("should return 0 after failed advance", func(t *testing.T) {
		// Arrange
		p, err := NewParser(strings.NewReader("push constant notanumber"))
		assert.NoError(t, err)

		// Act
		err = p.Advance()

		// Assert
		assert.Error(t, err)
		assert.Equal(t, 0, p.Arg2())
	})
}
