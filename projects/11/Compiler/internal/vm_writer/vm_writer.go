package vmwriter

import (
	"fmt"
	"io"
)

type Segment string

const (
	CONSTANT Segment = "constant"
	LOCAL    Segment = "local"
	ARGUMENT Segment = "argument"
	THIS     Segment = "this"
	THAT     Segment = "that"
	TEMP     Segment = "temp"
	POINTER  Segment = "pointer"
	STATIC   Segment = "static"
)

type ArithmeticCommand string

const (
	ADD ArithmeticCommand = "add"
	SUB ArithmeticCommand = "sub"
	NEG ArithmeticCommand = "neg"
	AND ArithmeticCommand = "and"
	OR  ArithmeticCommand = "or"
	NOT ArithmeticCommand = "not"
	EQ  ArithmeticCommand = "eq"
	GT  ArithmeticCommand = "gt"
	LT  ArithmeticCommand = "lt"
)

type IVMWriter interface {
	WritePush(segment Segment, index int, indentLevel int)
	WritePop(segment Segment, index int, indentLevel int)
	WriteArithmetic(command ArithmeticCommand, indentLevel int)
	WriteLabel(label string, indentLevel int)
	WriteGoto(label string, indentLevel int)
	WriteIf(label string, indentLevel int)
	WriteCall(name string, nArgs int, indentLevel int)
	WriteFunction(name string, nVars int)
	WriteReturn(indentLevel int)
}

type VMWriter struct {
	writer io.Writer
}

func New(w io.Writer) IVMWriter {
	return &VMWriter{writer: w}
}

func (w *VMWriter) WritePush(segment Segment, index int, indentLevel int) {
	w.writeLine(indentLevel, "push %s %d", segment, index)
}

func (w *VMWriter) WritePop(segment Segment, index int, indentLevel int) {
	w.writeLine(indentLevel, "pop %s %d", segment, index)
}

func (w *VMWriter) WriteArithmetic(command ArithmeticCommand, indentLevel int) {
	w.writeLine(indentLevel, "%s", command)
}

func (w *VMWriter) WriteLabel(label string, indentLevel int) {
	w.writeLine(indentLevel, "label %s", label)
}

func (w *VMWriter) WriteGoto(label string, indentLevel int) {
	w.writeLine(indentLevel, "goto %s", label)
}

func (w *VMWriter) WriteIf(label string, indentLevel int) {
	w.writeLine(indentLevel, "if-goto %s", label)
}

func (w *VMWriter) WriteCall(name string, nArgs int, indentLevel int) {
	w.writeLine(indentLevel, "call %s %d", name, nArgs)
}

func (w *VMWriter) WriteFunction(name string, nVars int) {
	w.writeLine(0, "function %s %d", name, nVars)
}

func (w *VMWriter) WriteReturn(indentLevel int) {
	w.writeLine(indentLevel, "return")
}

func (w *VMWriter) writeLine(indentLevel int, format string, a ...interface{}) {
	indent := ""
	for i := 0; i < indentLevel; i++ {
		indent += "  "
	}
	_, err := fmt.Fprintf(w.writer, "%s%s\n", indent, fmt.Sprintf(format, a...))
	if err != nil {
		panic(err)
	}
}
