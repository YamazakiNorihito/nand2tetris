package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
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
