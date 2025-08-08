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
	WritePush(segment Segment, index int)
	WritePop(segment Segment, index int)
	WriteArithmetic(command ArithmeticCommand)
	WriteLabel(label string)
	WriteGoto(label string)
	WriteIf(label string)
	WriteCall(name string, nArgs int)
	WriteFunction(name string, nVars int)
	WriteReturn()
}

type VMWriter struct {
	writer io.Writer
}

func New(w io.Writer) IVMWriter {
	return &VMWriter{writer: w}
}

func (w *VMWriter) WritePush(segment Segment, index int) {
	w.writeLine(1, "push %s %d", segment, index)
}

func (w *VMWriter) WritePop(segment Segment, index int) {
	w.writeLine(1, "pop %s %d", segment, index)
}

func (w *VMWriter) WriteArithmetic(command ArithmeticCommand) {
	w.writeLine(1, "%s", command)
}

func (w *VMWriter) WriteLabel(label string) {
	w.writeLine(0, "label %s", label)
}

func (w *VMWriter) WriteGoto(label string) {
	w.writeLine(1, "goto %s", label)
}

func (w *VMWriter) WriteIf(label string) {
	w.writeLine(1, "if-goto %s", label)
}

func (w *VMWriter) WriteCall(name string, nArgs int) {
	w.writeLine(1, "call %s %d", name, nArgs)
}

func (w *VMWriter) WriteFunction(name string, nVars int) {
	w.writeLine(0, "function %s %d", name, nVars)
}

func (w *VMWriter) WriteReturn() {
	w.writeLine(1, "return")
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
