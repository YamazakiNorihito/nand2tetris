package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	"ny/nand2tetris/compiler/internal/token_patterns"
	vmwriter "ny/nand2tetris/compiler/internal/vm_writer"
)

func (ce *CompilationEngine) compileReturn() error {
	// return
	token := ce.tokens[ce.index]
	if !token.IsReturn() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	returnStatementComponent := component.New("returnStatement", "")
	parentComponent.Children = append(parentComponent.Children, returnStatementComponent)

	returnStatementComponent.Children = append(returnStatementComponent.Children, component.New("keyword", string(token_patterns.RETURN)))
	ce.index++

	// expression
	ce.componentStack.Push(returnStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	returnStatementComponent = ce.componentStack.Pop()

	// ;
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	returnStatementComponent.Children = append(returnStatementComponent.Children, component.New("symbol", ";"))
	ce.index++

	// Write VM return
	ce.writeVMReturn()

	return nil
}

func (ce *CompilationEngine) writeVMReturn() {
	indentLevel := ce.componentStack.Count() + 1
	if ce.subroutineInfo.returnType == string(token_patterns.VOID) {
		ce.vmWriter.WritePush(vmwriter.CONSTANT, 0, indentLevel)
	}
	ce.vmWriter.WriteReturn(indentLevel)
}
