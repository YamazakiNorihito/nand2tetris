package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	"ny/nand2tetris/compiler/internal/token_patterns"
)

func (ce *CompilationEngine) compileWhile() error {
	// while
	token := ce.tokens[ce.index]
	if !token.IsWhile() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	whileStatementComponent := component.New("whileStatement", "")
	parentComponent.Children = append(parentComponent.Children, whileStatementComponent)

	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("keyword", string(token_patterns.WHILE)))
	ce.index++

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("symbol", "("))
	ce.index++

	// expression
	ce.componentStack.Push(whileStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	whileStatementComponent = ce.componentStack.Pop()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("symbol", ")"))
	ce.index++

	// '{'
	token = ce.tokens[ce.index]
	if !token.IsOpenBrace() {
		return fmt.Errorf("index %d: expected '{', got '%s'", ce.index, token.GetValue())
	}
	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("symbol", "{"))
	ce.index++

	// statements
	ce.componentStack.Push(whileStatementComponent)
	if err := ce.compileStatements(); err != nil {
		return err
	}
	whileStatementComponent = ce.componentStack.Pop()

	// '}'
	token = ce.tokens[ce.index]
	if !token.IsCloseBrace() {
		return fmt.Errorf("index %d: expected '}', got '%s'", ce.index, token.GetValue())
	}
	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("symbol", "}"))
	ce.index++

	return nil
}
