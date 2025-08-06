package compilation_engine

import "ny/nand2tetris/compiler/internal/component"

func (ce *CompilationEngine) compileExpressionList() (int, error) {
	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	expressionListComponent := component.New("expressionList", "")
	parentComponent.Children = append(parentComponent.Children, expressionListComponent)

	token := ce.tokens[ce.index]
	if !token.IsTerm() {
		return 0, nil
	}

	count := 0
	for {
		// expression
		ce.componentStack.Push(expressionListComponent)
		if err := ce.compileExpression(); err != nil {
			return 0, err
		}
		expressionListComponent = ce.componentStack.Pop()
		count++

		// ','
		token = ce.tokens[ce.index]
		if !token.IsComma() {
			break
		}
		expressionListComponent.Children = append(expressionListComponent.Children, component.New("symbol", ","))
		ce.index++
	}

	return count, nil
}
