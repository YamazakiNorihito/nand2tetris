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
	writer       io.Writer
	filename     string
	labelCounter int
}

func NewCodeWriter(w io.Writer) (ICodeWriter, error) {
	return &codeWriter{writer: w}, nil
}

func (cw *codeWriter) SetFileName(filename string) {
	cw.filename = filename
}

func (cw *codeWriter) WriteArithmetic(command string) error {
	var assembly string
	switch command {
	case "add":
		assembly = add
	case "sub":
		assembly = sub
	case "neg":
		assembly = neg
	case "and":
		assembly = and
	case "or":
		assembly = or
	case "not":
		assembly = not
	case "eq":
		assembly = fmt.Sprintf(eq, fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter))
		cw.labelCounter++
	case "gt":
		assembly = fmt.Sprintf(gt, fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter))
		cw.labelCounter++
	case "lt":
		assembly = fmt.Sprintf(lt, fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter), fmt.Sprint(cw.labelCounter))
		cw.labelCounter++
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
		assembly = fmt.Sprintf(pushConstant, index, index)
	case "local":
		assembly = fmt.Sprintf(pushLocal, index, index)
	case "argument":
		assembly = fmt.Sprintf(pushArgument, index, index)
	case "this":
		assembly = fmt.Sprintf(pushThis, index, index)
	case "that":
		assembly = fmt.Sprintf(pushThat, index, index)
	case "temp":
		assembly = fmt.Sprintf(pushTemp, index, index)
	case "pointer":
		if index != 0 && index != 1 {
			err = fmt.Errorf("invalid index for pointer segment: %d (must be 0 or 1)", index)
			break
		}
		symbol := "THIS"
		if index == 1 {
			symbol = "THAT"
		}
		assembly = fmt.Sprintf(pushPointer, fmt.Sprint(index), symbol)
	case "static":
		if cw.filename == "" {
			err = errors.New("filename not set for static variable")
			break
		}
		assembly = fmt.Sprintf(pushStatic, fmt.Sprint(index), cw.filename, fmt.Sprint(index))
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
	case "local":
		assembly = fmt.Sprintf(popLocal, index, index)
	case "argument":
		assembly = fmt.Sprintf(popArgument, index, index)
	case "this":
		assembly = fmt.Sprintf(popThis, index, index)
	case "that":
		assembly = fmt.Sprintf(popThat, index, index)
	case "temp":
		assembly = fmt.Sprintf(popTemp, index, index)
	case "pointer":
		if index != 0 && index != 1 {
			err = fmt.Errorf("invalid index for pointer segment: %d (must be 0 or 1)", index)
			break
		}
		symbol := "THIS"
		if index == 1 {
			symbol = "THAT"
		}
		assembly = fmt.Sprintf(popPointer, fmt.Sprint(index), symbol)
	case "static":
		if cw.filename == "" {
			err = errors.New("filename not set for static variable")
			break
		}
		assembly = fmt.Sprintf(popStatic, fmt.Sprint(index), cw.filename, fmt.Sprint(index))
	default:
		err = fmt.Errorf("unsupported segment for pop command: %s", segment)
	}

	if err != nil {
		return err
	}
	_, err = fmt.Fprint(cw.writer, assembly)
	return err
}
