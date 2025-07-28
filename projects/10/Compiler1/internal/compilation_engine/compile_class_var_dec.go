package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
)

func (ce *CompilationEngine) compileClassVarDec() error {
	// static or field
	token := ce.tokens[ce.index]
	if !token.IsClassVarDec() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	classVarDecComponent := component.New("classVarDec", "")
	parentComponent.Children = append(parentComponent.Children, classVarDecComponent)

	classVarDecComponent.Children = append(classVarDecComponent.Children, component.New("keyword", string(token.GetKeyword())))
	ce.index++

	// Type
	token = ce.tokens[ce.index]
	if !token.IsType() {
		return fmt.Errorf("index %d: expected type, got '%s'", ce.index, token.GetValue())
	}
	if token.IsIdentifier() {
		classVarDecComponent.Children = append(classVarDecComponent.Children, component.New("identifier", token.GetValue()))
		ce.index++
	} else {
		classVarDecComponent.Children = append(classVarDecComponent.Children, component.New("keyword", token.GetValue()))
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

		classVarDecComponent.Children = append(classVarDecComponent.Children, component.New("identifier", token.GetIdentifier()))
		ce.index++

		// Comma
		token = ce.tokens[ce.index]
		if !token.IsComma() {
			break
		}
		classVarDecComponent.Children = append(classVarDecComponent.Children, component.New("symbol", ","))
		ce.index++
	}

	// ';'
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	classVarDecComponent.Children = append(classVarDecComponent.Children, component.New("symbol", ";"))
	ce.index++

	return nil
}
