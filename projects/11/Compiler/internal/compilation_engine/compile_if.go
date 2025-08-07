package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	"ny/nand2tetris/compiler/internal/token_patterns"
)

func (ce *CompilationEngine) compileIf() error {
	// if
	token := ce.tokens[ce.index]
	if !token.IsIf() {
		return nil
	}
	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	ifStatementComponent := component.New("ifStatement", "")
	parentComponent.Children = append(parentComponent.Children, ifStatementComponent)
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("keyword", string(token_patterns.IF)))
	ce.index++

	whileLabelCounter := ce.labelCounterIf
	ce.labelCounterIf++

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "("))
	ce.index++

	// expression
	ce.componentStack.Push(ifStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	ifStatementComponent = ce.componentStack.Pop()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", ")"))
	ce.index++

	// '{'
	token = ce.tokens[ce.index]
	if !token.IsOpenBrace() {
		return fmt.Errorf("index %d: expected '{', got '%s'", ce.index, token.GetValue())
	}
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "{"))
	ce.index++

	ce.vmWriter.WriteIf(fmt.Sprintf("IF_TRUE%d", whileLabelCounter), ce.componentStack.Count()+1)
	ce.vmWriter.WriteGoto(fmt.Sprintf("IF_FALSE%d", whileLabelCounter), ce.componentStack.Count()+1)

	// statements
	ce.vmWriter.WriteLabel(fmt.Sprintf("IF_TRUE%d", whileLabelCounter), ce.componentStack.Count()+1)
	ce.componentStack.Push(ifStatementComponent)
	if err := ce.compileStatements(); err != nil {
		return err
	}
	ifStatementComponent = ce.componentStack.Pop()

	// '}'
	token = ce.tokens[ce.index]
	if !token.IsCloseBrace() {
		return fmt.Errorf("index %d: expected '}', got '%s'", ce.index, token.GetValue())
	}
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "}"))
	ce.index++

	// else
	token = ce.tokens[ce.index]
	if token.IsElse() {
		ce.vmWriter.WriteGoto(fmt.Sprintf("IF_END%d", whileLabelCounter), ce.componentStack.Count()+1)

		ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("keyword", string(token_patterns.ELSE)))
		ce.index++

		// '{'
		token = ce.tokens[ce.index]
		if !token.IsOpenBrace() {
			return fmt.Errorf("index %d: expected '{', got '%s'", ce.index, token.GetValue())
		}
		ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "{"))
		ce.index++

		ce.vmWriter.WriteLabel(fmt.Sprintf("IF_FALSE%d", whileLabelCounter), ce.componentStack.Count()+1)

		// statements
		ce.componentStack.Push(ifStatementComponent)
		if err := ce.compileStatements(); err != nil {
			return err
		}
		ifStatementComponent = ce.componentStack.Pop()

		// '}'
		token = ce.tokens[ce.index]
		if !token.IsCloseBrace() {
			return fmt.Errorf("index %d: expected '}', got '%s'", ce.index, token.GetValue())
		}
		ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "}"))
		ce.index++
		ce.vmWriter.WriteLabel(fmt.Sprintf("IF_END%d", whileLabelCounter), ce.componentStack.Count()+1)
	} else {
		ce.vmWriter.WriteLabel(fmt.Sprintf("IF_FALSE%d", whileLabelCounter), ce.componentStack.Count()+1)
	}
	return nil
}
