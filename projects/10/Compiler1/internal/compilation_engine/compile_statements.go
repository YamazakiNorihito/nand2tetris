package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
)

func (ce *CompilationEngine) compileStatements() error {
	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	statementsComponent := component.New("statements", "")
	parentComponent.Children = append(parentComponent.Children, statementsComponent)

	token := ce.tokens[ce.index]
	if !token.IsStatement() {
		return nil
	}

	for token.IsStatement() {
		ce.componentStack.Push(statementsComponent)

		switch {
		case token.IsLet():
			if err := ce.compileLet(); err != nil {
				return err
			}
		case token.IsIf():
			if err := ce.compileIf(); err != nil {
				return err
			}
		case token.IsWhile():
			if err := ce.compileWhile(); err != nil {
				return err
			}
		case token.IsDo():
			if err := ce.compileDo(); err != nil {
				return err
			}
		case token.IsReturn():
			if err := ce.compileReturn(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("index %d: unexpected statement type: %s", ce.index, token.GetValue())
		}
		token = ce.tokens[ce.index]
		statementsComponent = ce.componentStack.Pop()
	}
	return nil
}
