package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	symboltable "ny/nand2tetris/compiler/internal/symbol_table"
)

func (ce *CompilationEngine) compileClassVarDec() error {
	// static or field
	token := ce.tokens[ce.index]
	if !token.IsClassVarDec() {
		return nil
	}

	var variableKind symboltable.VariableKind = symboltable.FIELD
	if token.IsStatic() {
		variableKind = symboltable.STATIC
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

	classVarType := token.GetValue()
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

		ce.symbolTable.Define(token.GetIdentifier(), classVarType, variableKind)

		if variableKind == symboltable.STATIC {
			component := component.NewVariableComponent(
				"identifier", token.GetIdentifier(), component.STATIC, ce.symbolTable.IndexOf(token.GetIdentifier()), component.DECLARED)
			classVarDecComponent.Children = append(classVarDecComponent.Children, component)
		} else {
			component := component.NewVariableComponent(
				"identifier", token.GetIdentifier(), component.FIELD, ce.symbolTable.IndexOf(token.GetIdentifier()), component.DECLARED)
			classVarDecComponent.Children = append(classVarDecComponent.Children, component)
		}
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
