package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
	"ny/nand2tetris/compiler1/internal/token_patterns"
)

func (ce *CompilationEngine) compileLet() error {
	// let
	token := ce.tokens[ce.index]
	if !token.IsLet() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	letStatementComponent := component.New("letStatement", "")
	parentComponent.Children = append(parentComponent.Children, letStatementComponent)

	letStatementComponent.Children = append(letStatementComponent.Children, component.New("keyword", string(token_patterns.LET)))
	ce.index++

	// variable name
	token = ce.tokens[ce.index]
	if !token.IsIdentifier() {
		return fmt.Errorf("index %d: expected variable name, got '%s'", ce.index, token.GetValue())
	}
	letStatementComponent.Children = append(letStatementComponent.Children, component.New("identifier", token.GetIdentifier()))
	ce.index++

	// '['
	token = ce.tokens[ce.index]
	if token.IsOpenBracket() {
		letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "["))
		ce.index++

		// expression
		ce.componentStack.Push(letStatementComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		letStatementComponent = ce.componentStack.Pop()

		// ']'
		token = ce.tokens[ce.index]
		if !token.IsCloseBracket() {
			return fmt.Errorf("index %d: expected ']', got '%s'", ce.index, token.GetValue())
		}
		letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "]"))
		ce.index++
	}

	// '='
	token = ce.tokens[ce.index]
	if !token.IsEqual() {
		return fmt.Errorf("index %d: expected '=', got '%s'", ce.index, token.GetValue())
	}
	letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "="))
	ce.index++

	// expression
	ce.componentStack.Push(letStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	letStatementComponent = ce.componentStack.Pop()

	// ;
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", ";"))
	ce.index++

	return nil
}
