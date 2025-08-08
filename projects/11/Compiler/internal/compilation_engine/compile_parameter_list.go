package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	symboltable "ny/nand2tetris/compiler/internal/symbol_table"
)

func (ce *CompilationEngine) compileParameterList() error {
	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	paramListComponent := component.New("parameterList", "")
	parentComponent.Children = append(parentComponent.Children, paramListComponent)

	token := ce.tokens[ce.index]
	if !token.IsType() {
		return nil
	}

	for {
		token = ce.tokens[ce.index]
		if !token.IsType() {
			return fmt.Errorf("index %d: expected type after ',', got '%s'", ce.index, token.GetValue())
		}

		classVarType := token.GetValue()
		// Type token (e.g., int, char, boolean, className)
		if token.IsIdentifier() {
			paramListComponent.Children = append(paramListComponent.Children, component.New("identifier", token.GetValue()))

		} else {
			paramListComponent.Children = append(paramListComponent.Children, component.New("keyword", token.GetValue()))
		}
		ce.index++

		// Parameter name
		token = ce.tokens[ce.index]
		if !token.IsIdentifier() {
			return fmt.Errorf("index %d: expected parameter name, got '%s'", ce.index, token.GetValue())
		}

		ce.symbolTable.Define(token.GetIdentifier(), classVarType, symboltable.ARG)

		paramListComponent.Children = append(paramListComponent.Children,
			component.NewVariableComponent(
				"identifier",
				token.GetIdentifier(),
				component.ARGMENT,
				ce.symbolTable.IndexOf(token.GetIdentifier()),
				component.DECLARED))
		ce.index++

		// If there is no comma, parameter list ends
		token = ce.tokens[ce.index]
		if !token.IsComma() {
			break
		}

		paramListComponent.Children = append(paramListComponent.Children, component.New("symbol", ","))
		ce.index++
	}

	return nil
}
