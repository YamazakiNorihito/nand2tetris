package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
)

func (ce *CompilationEngine) compileExpression() error {
	token := ce.tokens[ce.index]
	if !token.IsTerm() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	expressionComponent := component.New("expression", "")
	parentComponent.Children = append(parentComponent.Children, expressionComponent)

	for {
		token = ce.tokens[ce.index]
		if !token.IsTerm() {
			return fmt.Errorf("index %d: expected term, got '%s'", ce.index, token.GetValue())
		}

		// Term
		ce.componentStack.Push(expressionComponent)
		if err := ce.compileTerm(); err != nil {
			return err
		}
		expressionComponent = ce.componentStack.Pop()

		// Operator
		token = ce.tokens[ce.index]
		if !token.IsOp() {
			break
		}
		expressionComponent.Children = append(expressionComponent.Children, component.New("symbol", token.GetValue()))
		ce.index++
	}
	return nil
}
