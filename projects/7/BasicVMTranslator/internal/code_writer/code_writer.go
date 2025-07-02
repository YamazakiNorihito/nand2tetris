package codewriter

import (
	"errors"
	"fmt"
	"io"
	"ny/nand2tetris/basicvmtranslator/internal/config"
)

var (
	segmentSymbolMap = map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
	}
)

type ICodeWriter interface {
	SetFileName(filename string)
	WriteArithmetic(command string) error
	WritePushPop(command config.CommandType, segment string, index int) error
}

type codeWriter struct {
	writer io.Writer
	// filename は現在処理中のVMファイル名で、static変数のシンボル生成に使われます。
	filename string
	// labelCounter は比較演算子でユニークなラベルを生成するために使用します。
	labelCounter int
}

func NewCodeWriter(w io.Writer) (ICodeWriter, error) {
	bootstrapAssembly := `// Bootstrap code
@256
D=A
@SP
M=D
`
	_, err := fmt.Fprint(w, bootstrapAssembly)

	if err != nil {
		return nil, err
	}

	return &codeWriter{writer: w}, nil
}

func (cw *codeWriter) SetFileName(filename string) {
	cw.filename = filename
}

func (cw *codeWriter) WriteArithmetic(command string) error {
	var assembly string
	switch command {
	case "add": // x + y
		assembly = `// add
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
`
	case "sub": // x - y
		assembly = `// sub
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
`
	case "neg": // -y
		assembly = `// neg
@SP
M=M-1 // SP--
A=M
M=-M
@SP
M=M+1 // SP++
`
	case "and": // x & y
		assembly = `// and
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
`
	case "or": // x | y
		assembly = `// or
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
`
	case "not": // !y
		assembly = `// not
@SP
M=M-1
A=M
M=!M
@SP
M=M+1
`
	default:
		return fmt.Errorf("unsupported arithmetic command: %s", command)
	}

	_, err := fmt.Fprint(cw.writer, assembly)
	return err
}

// 忘れちゃうのでメモ
// PushもPopもStackへ値をPush \Popする
// pushは今割り当てられている変数の値をstackのSPにPushする(変数のD Registoryの値をSPに設定する)
// popはstackのSPの値を変数にPopする(変数のD Registerに値を設定する)
func (cw *codeWriter) WritePushPop(command config.CommandType, segment string, index int) error {
	switch command {
	case config.C_PUSH:
		return cw.writePush(segment, index)
	case config.C_POP:
		return cw.writePop(segment, index)
	default:
		return fmt.Errorf("unsupported command type for WritePushPop: %v", command)
	}
}

func (cw *codeWriter) writePush(segment string, index int) error {
	var assembly string
	var err error

	switch segment {
	case "constant":
		assembly = fmt.Sprintf(`// push constant %d
@%d
D=A
@SP
A=M
M=D
@SP
M=M+1
`, index, index)

	case "local", "argument", "this", "that":
		symbol := segmentSymbolMap[segment]
		assembly = fmt.Sprintf(`// push %s %d
@%s
D=M
@%d
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`, segment, index, symbol, index)

	case "temp":
		assembly = fmt.Sprintf(`// push temp %d
@5
D=A
@%d
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`, index, index)

	case "pointer":
		if index != 0 && index != 1 {
			err = fmt.Errorf("invalid index for pointer segment: %d (must be 0 or 1)", index)
			break
		}
		symbol := "THIS"
		if index == 1 {
			symbol = "THAT"
		}
		assembly = fmt.Sprintf(`// push pointer %d
@%s
D=M
@SP
A=M
M=D
@SP
M=M+1
`, index, symbol)

	case "static":
		if cw.filename == "" {
			err = errors.New("filename not set for static variable")
			break
		}
		assembly = fmt.Sprintf(`// push static %d
@%s.%d // アセンブラによってアドレス16から自動で割り当てられる
D=M
@SP
A=M
M=D
@SP
M=M+1
`, index, cw.filename, index)

	default:
		err = fmt.Errorf("unsupported segment for push command: %s", segment)
	}

	if err != nil {
		return err
	}
	_, err = fmt.Fprint(cw.writer, assembly)
	return err
}

func (cw *codeWriter) writePop(segment string, index int) error {
	var assembly string
	var err error

	switch segment {
	case "local", "argument", "this", "that":
		symbol := segmentSymbolMap[segment]
		assembly = fmt.Sprintf(`// pop %s %d
@%s
D=M
@%d
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
`, segment, index, symbol, index)

	case "temp":
		assembly = fmt.Sprintf(`// pop temp %d
@5
D=A
@%d
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
`, index, index)

	case "pointer":
		if index != 0 && index != 1 {
			err = fmt.Errorf("invalid index for pointer segment: %d (must be 0 or 1)", index)
			break
		}
		symbol := "THIS"
		if index == 1 {
			symbol = "THAT"
		}
		assembly = fmt.Sprintf(`// pop pointer %d
@SP
M=M-1
A=M
D=M
@%s
M=D
`, index, symbol)

	case "static":
		if cw.filename == "" {
			err = errors.New("filename not set for static variable")
			break
		}
		assembly = fmt.Sprintf(`// pop static %d
@SP
M=M-1
A=M
D=M
@%s.%d
M=D
`, index, cw.filename, index)

	default:
		err = fmt.Errorf("unsupported segment for pop command: %s", segment)
	}

	if err != nil {
		return err
	}
	_, err = fmt.Fprint(cw.writer, assembly)
	return err
}
