package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
	Tokens "ny/nand2tetris/compiler1/internal/tokens"
)

func (ce *CompilationEngine) compileSubroutineCall() error {
	token := ce.tokens[ce.index]

	var nextToken Tokens.IToken
	if len(ce.tokens) > ce.index+1 {
		nextToken = ce.tokens[ce.index+1]
	}
	if !token.IsSubroutineCall(nextToken) {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	parentComponent.Children = append(parentComponent.Children, component.New("identifier", token.GetValue()))
	ce.index++

	// '.'
	token = ce.tokens[ce.index]
	if token.IsDot() {
		parentComponent.Children = append(parentComponent.Children, component.New("symbol", "."))
		ce.index++

		token = ce.tokens[ce.index]
		if !token.IsIdentifier() {
			return fmt.Errorf("index %d: expected subroutine name after '.', got '%s'", ce.index, token.GetValue())
		}
		parentComponent.Children = append(parentComponent.Children, component.New("identifier", token.GetIdentifier()))
		ce.index++
	}

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	parentComponent.Children = append(parentComponent.Children, component.New("symbol", "("))
	ce.index++

	// expressionList
	ce.componentStack.Push(parentComponent)
	_, err := ce.compileExpressionList()
	if err != nil {
		return err
	}
	parentComponent = ce.componentStack.Pop()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	parentComponent.Children = append(parentComponent.Children, component.New("symbol", ")"))
	ce.index++

	return nil
}
