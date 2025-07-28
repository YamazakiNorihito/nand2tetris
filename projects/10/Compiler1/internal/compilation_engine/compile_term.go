package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
	Tokens "ny/nand2tetris/compiler1/internal/tokens"
)

func (ce *CompilationEngine) compileTerm() error {
	// term
	token := ce.tokens[ce.index]
	if !token.IsTerm() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	termComponent := component.New("term", "")
	parentComponent.Children = append(parentComponent.Children, termComponent)

	if token.IsIntConst() {
		termComponent.Children = append(termComponent.Children, component.New("integerConstant", token.GetValue()))
		ce.index++
		return nil
	}

	if token.IsStringConst() {
		termComponent.Children = append(termComponent.Children, component.New("stringConstant", token.GetValue()))
		ce.index++
		return nil
	}

	if token.IsKeywordConstant() {
		termComponent.Children = append(termComponent.Children, component.New("keyword", token.GetValue()))
		ce.index++
		return nil
	}

	if token.IsUnaryOp() {
		// unary operation '-' or '~'
		termComponent.Children = append(termComponent.Children, component.New("symbol", token.GetValue()))
		ce.index++

		// term
		ce.componentStack.Push(termComponent)
		if err := ce.compileTerm(); err != nil {
			return err
		}
		termComponent = ce.componentStack.Pop()
		return nil
	}

	if token.IsOpenParen() {
		// '('
		termComponent.Children = append(termComponent.Children, component.New("symbol", "("))
		ce.index++

		// expression
		ce.componentStack.Push(termComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		termComponent = ce.componentStack.Pop()

		// ')'
		token = ce.tokens[ce.index]
		if !token.IsCloseParen() {
			return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, ce.tokens[ce.index].GetValue())
		}
		termComponent.Children = append(termComponent.Children, component.New("symbol", ")"))
		ce.index++
		return nil
	}

	// Identifier
	var nextToken Tokens.IToken
	if len(ce.tokens) > ce.index+1 {
		nextToken = ce.tokens[ce.index+1]
	}
	if token.IsArrayItem(nextToken) {
		// array item
		termComponent.Children = append(termComponent.Children, component.New("identifier", token.GetIdentifier()))
		ce.index++

		// '['
		if !nextToken.IsOpenBracket() {
			return fmt.Errorf("index %d: expected '[', got '%s'", ce.index, nextToken.GetValue())
		}
		termComponent.Children = append(termComponent.Children, component.New("symbol", "["))
		ce.index++

		// expression
		ce.componentStack.Push(termComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		termComponent = ce.componentStack.Pop()

		// ']'
		nextToken = ce.tokens[ce.index]
		if !nextToken.IsCloseBracket() {
			return fmt.Errorf("index %d: expected ']', got '%s'", ce.index, nextToken.GetValue())
		}
		termComponent.Children = append(termComponent.Children, component.New("symbol", "]"))
		ce.index++
	} else if token.IsSubroutineCall(nextToken) {
		// subroutine call
		ce.componentStack.Push(termComponent)
		if err := ce.compileSubroutineCall(); err != nil {
			return err
		}
		termComponent = ce.componentStack.Pop()
	} else {
		// identifier
		termComponent.Children = append(termComponent.Children, component.New("identifier", token.GetIdentifier()))
		ce.index++
	}

	return nil
}
