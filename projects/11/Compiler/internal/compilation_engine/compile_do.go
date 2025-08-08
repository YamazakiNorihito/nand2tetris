package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	"ny/nand2tetris/compiler/internal/token_patterns"
	vmwriter "ny/nand2tetris/compiler/internal/vm_writer"
)

func (ce *CompilationEngine) compileDo() error {
	// do
	token := ce.tokens[ce.index]
	if !token.IsDo() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	doStatementComponent := component.New("doStatement", "")
	parentComponent.Children = append(parentComponent.Children, doStatementComponent)

	doStatementComponent.Children = append(doStatementComponent.Children, component.New("keyword", string(token_patterns.DO)))
	ce.index++

	// subroutine call
	ce.componentStack.Push(doStatementComponent)
	if err := ce.compileSubroutineCall(); err != nil {
		return err
	}
	doStatementComponent = ce.componentStack.Pop()

	// ';'
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	doStatementComponent.Children = append(doStatementComponent.Children, component.New("symbol", ";"))
	ce.index++

	ce.vmWriter.WritePop(vmwriter.TEMP, 0)

	return nil
}
