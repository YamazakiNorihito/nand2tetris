package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
)

func (ce *CompilationEngine) compileVarDec() error {
	// var
	token := ce.tokens[ce.index]
	if !token.IsVar() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	varDec := component.New("varDec", "")
	parentComponent.Children = append(parentComponent.Children, varDec)

	varDec.Children = append(varDec.Children, component.New("keyword", "var"))
	ce.index++

	// Type
	token = ce.tokens[ce.index]
	if !token.IsType() {
		return fmt.Errorf("index %d: expected type, got '%s'", ce.index, token.GetValue())
	}
	if token.IsIdentifier() {
		varDec.Children = append(varDec.Children, component.New("identifier", token.GetValue()))
		ce.index++
	} else {
		varDec.Children = append(varDec.Children, component.New("keyword", token.GetValue()))
		ce.index++
	}

	// Variable name
	token = ce.tokens[ce.index]
	if !token.IsIdentifier() {
		return fmt.Errorf("index %d: expected variable name, got '%s'", ce.index, token.GetValue())
	}

	for {
		token = ce.tokens[ce.index]
		if !token.IsIdentifier() {
			return fmt.Errorf("index %d: expected variable name, got '%s'", ce.index, token.GetValue())
		}

		varDec.Children = append(varDec.Children, component.New("identifier", token.GetIdentifier()))
		ce.index++

		// Comma
		token = ce.tokens[ce.index]
		if !token.IsComma() {
			break
		}
		varDec.Children = append(varDec.Children, component.New("symbol", ","))
		ce.index++
	}

	// ';'
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	varDec.Children = append(varDec.Children, component.New("symbol", ";"))
	ce.index++

	return nil
}
