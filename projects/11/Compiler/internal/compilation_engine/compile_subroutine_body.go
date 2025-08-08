package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	symboltable "ny/nand2tetris/compiler/internal/symbol_table"
	vmwriter "ny/nand2tetris/compiler/internal/vm_writer"
)

func (ce *CompilationEngine) compileSubroutineBody() error {
	// '{'
	token := ce.tokens[ce.index]
	if !token.IsOpenBrace() {
		return fmt.Errorf("expected '{', got '%s'", token.GetValue())
	}
	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	subroutineBodyComponent := component.New("subroutineBody", "")
	parentComponent.Children = append(parentComponent.Children, subroutineBodyComponent)

	subroutineBodyComponent.Children = append(subroutineBodyComponent.Children, component.New("symbol", "{"))
	ce.index++

	// Compile variable declarations
	token = ce.tokens[ce.index]
	for token.IsVar() {
		ce.componentStack.Push(subroutineBodyComponent)

		if err := ce.compileVarDec(); err != nil {
			return err
		}
		token = ce.tokens[ce.index]
		subroutineBodyComponent = ce.componentStack.Pop()
	}

	// Write VM function declaration
	ce.writeVMFunctionDeclare()

	// Compile statements
	ce.componentStack.Push(subroutineBodyComponent)
	if err := ce.compileStatements(); err != nil {
		return err
	}
	subroutineBodyComponent = ce.componentStack.Pop()

	// '}'
	token = ce.tokens[ce.index]
	if !token.IsCloseBrace() {
		return fmt.Errorf("expected '}', got '%s'", token.GetValue())
	}
	subroutineBodyComponent.Children = append(subroutineBodyComponent.Children, component.New("symbol", "}"))
	ce.index++

	return nil
}

func (ce *CompilationEngine) writeVMFunctionDeclare() {
	indentLevel := ce.componentStack.Count() + 1
	ce.vmWriter.WriteFunction(ce.className+"."+ce.subroutineInfo.name, ce.symbolTable.VarCount(symboltable.VAR))

	switch ce.subroutineInfo.functionType {
	case CONSTRUCTOR:
		ce.vmWriter.WritePush(vmwriter.CONSTANT, ce.symbolTable.VarCount(symboltable.FIELD), indentLevel)
		ce.vmWriter.WriteCall("Memory.alloc", 1, indentLevel)
		ce.vmWriter.WritePop(vmwriter.POINTER, 0, indentLevel)
	case METHOD:
		ce.vmWriter.WritePush(vmwriter.ARGUMENT, 0, indentLevel)
		ce.vmWriter.WritePop(vmwriter.POINTER, 0, indentLevel)
	case FUNCTION:
		// No additional actions needed for functions
	default:
		panic("unknown subroutine type")
	}
}
