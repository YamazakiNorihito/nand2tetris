package codewriter

import (
	"bytes"
	"errors"
	"testing"

	"ny/nand2tetris/basicvmtranslator/internal/config"

	"github.com/stretchr/testify/assert"
)

// Mock writer that can simulate write errors
type errorWriter struct{}

func (e *errorWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("forced write error")
}

func TestNewCodeWriter(t *testing.T) {
	t.Run("writes bootstrap code on initialization", func(t *testing.T) {
		var buf bytes.Buffer
		cw, err := NewCodeWriter(&buf)
		assert.NoError(t, err)
		assert.NotNil(t, cw)
	})
}

func TestSetFileName(t *testing.T) {
	var buf bytes.Buffer
	cw, _ := NewCodeWriter(&buf)
	cw.SetFileName("Foo")
	assert.Equal(t, "Foo", cw.(*codeWriter).filename)
}

func TestWriteArithmetic(t *testing.T) {
	var buf bytes.Buffer
	cw, _ := NewCodeWriter(&buf)

	t.Run("success with supported arithmetic commands", func(t *testing.T) {
		tests := []struct {
			cmd      string
			expected string
		}{
			{"add", `// add
@SP
M=M-1 // SP--
A=M   // D = *SP
D=M
@SP
M=M-1 // SP--
A=M
M=D+M
@SP
M=M+1 // SP++
`},
			{"sub", `// sub
@SP
M=M-1 // SP--
A=M   // D = *SP
D=M
@SP
M=M-1 // SP--
A=M
M=M-D
@SP
M=M+1 // SP++
`},
			{"neg", `// neg
@SP
M=M-1 // SP--
A=M
M=-M
@SP
M=M+1 // SP++
`},
			{"and", `// and
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M&D
@SP
M=M+1
`},
			{"or", `// or
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=D|M
@SP
M=M+1
`},
			{"not", `// not
@SP
M=M-1
A=M
M=!M
@SP
M=M+1
`},
		}

		for _, tt := range tests {
			buf.Reset()
			err := cw.WriteArithmetic(tt.cmd)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, buf.String())
		}
	})

	t.Run("failure with unsupported arithmetic command", func(t *testing.T) {
		buf.Reset()
		err := cw.WriteArithmetic("foo")
		assert.Error(t, err, "unsupported command should return error")
	})
}

func TestWritePushPop(t *testing.T) {
	t.Run("Push", func(t *testing.T) {
		var buf bytes.Buffer
		cw, _ := NewCodeWriter(&buf)
		cw.SetFileName("TestFile")

		t.Run("success with supported segments", func(t *testing.T) {
			segments := []struct {
				segment  string
				index    int
				desc     string
				expected string
			}{
				{"constant", 7, "push constant 7", `// push constant 7
@7
D=A
@SP
A=M
M=D
@SP
M=M+1
`},
				{"local", 2, "push local 2", `// push local 2
@LCL
D=M
@2
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`},
				{"argument", 1, "push argument 1", `// push argument 1
@ARG
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`},
				{"this", 0, "push this 0", `// push this 0
@THIS
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`},
				{"that", 3, "push that 3", `// push that 3
@THAT
D=M
@3
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`},
				{"temp", 4, "push temp 4", `// push temp 4
@5
D=A
@4
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`},
				{"pointer", 0, "push pointer 0", `// push pointer 0
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
`},
				{"pointer", 1, "push pointer 1", `// push pointer 1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
`},
				{"static", 5, "push static 5", `// push static 5
@TestFile.5 // アセンブラによってアドレス16から自動で割り当てられる
D=M
@SP
A=M
M=D
@SP
M=M+1
`},
			}

			for _, tt := range segments {
				buf.Reset()
				err := cw.WritePushPop(config.C_PUSH, tt.segment, tt.index)
				assert.NoError(t, err, tt.desc)
				assert.Equal(t, tt.expected, buf.String(), tt.desc)
			}
		})

		t.Run("failure with invalid pointer index", func(t *testing.T) {
			buf.Reset()
			err := cw.WritePushPop(config.C_PUSH, "pointer", 2)
			assert.Error(t, err)
		})

		t.Run("failure with missing filename for static", func(t *testing.T) {
			cw2, _ := NewCodeWriter(&buf)
			cw2.(*codeWriter).filename = ""
			err := cw2.WritePushPop(config.C_PUSH, "static", 1)
			assert.Error(t, err)
		})

		t.Run("failure with unsupported segment", func(t *testing.T) {
			err := cw.WritePushPop(config.C_PUSH, "foo", 0)
			assert.Error(t, err)
		})
	})

	t.Run("Pop", func(t *testing.T) {
		var buf bytes.Buffer
		cw, _ := NewCodeWriter(&buf)
		cw.SetFileName("TestFile")

		t.Run("success with supported segments", func(t *testing.T) {
			segments := []struct {
				segment  string
				index    int
				desc     string
				expected string
			}{
				{"local", 2, "pop local 2", `// pop local 2
@LCL
D=M
@2
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`},
				{"argument", 1, "pop argument 1", `// pop argument 1
@ARG
D=M
@1
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`},
				{"this", 0, "pop this 0", `// pop this 0
@THIS
D=M
@0
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`},
				{"that", 3, "pop that 3", `// pop that 3
@THAT
D=M
@3
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`},
				{"temp", 4, "pop temp 4", `// pop temp 4
@5
D=A
@4
D=D+A
@R13
M=D
@SP
M=M-1
A=M
D=M
@R13
A=M
M=D
`},
				{"pointer", 0, "pop pointer 0", `// pop pointer 0
@SP
M=M-1
A=M
D=M
@THIS
M=D
`},
				{"pointer", 1, "pop pointer 1", `// pop pointer 1
@SP
M=M-1
A=M
D=M
@THAT
M=D
`},
				{"static", 5, "pop static 5", `// pop static 5
@SP
M=M-1
A=M
D=M
@TestFile.5
M=D
`},
			}

			for _, tt := range segments {
				buf.Reset()
				err := cw.WritePushPop(config.C_POP, tt.segment, tt.index)
				assert.NoError(t, err, tt.desc)
				assert.Equal(t, tt.expected, buf.String(), tt.desc)
			}
		})

		t.Run("failure with invalid pointer index", func(t *testing.T) {
			buf.Reset()
			err := cw.WritePushPop(config.C_POP, "pointer", 2)
			assert.Error(t, err)
		})

		t.Run("failure with missing filename for static", func(t *testing.T) {
			cw2, _ := NewCodeWriter(&buf)
			cw2.(*codeWriter).filename = ""
			err := cw2.WritePushPop(config.C_POP, "static", 1)
			assert.Error(t, err)
		})

		t.Run("failure with unsupported segment", func(t *testing.T) {
			err := cw.WritePushPop(config.C_POP, "foo", 0)
			assert.Error(t, err)
		})
	})

	t.Run("failure with unsupported command type", func(t *testing.T) {
		var buf bytes.Buffer
		cw, _ := NewCodeWriter(&buf)
		err := cw.WritePushPop(config.C_UNKNOWN, "local", 0)
		assert.Error(t, err)
	})

	t.Run("failure with writer error", func(t *testing.T) {
		cw := &codeWriter{writer: &errorWriter{}}
		cw.SetFileName("TestFile")
		err := cw.WritePushPop(config.C_PUSH, "constant", 1)
		assert.Error(t, err)
	})
}
