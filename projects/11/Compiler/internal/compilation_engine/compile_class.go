package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	"ny/nand2tetris/compiler/internal/token_patterns"
)

func (ce *CompilationEngine) CompileClass() error {
	classComponent := component.New("class", "")

	// 'class'
	token := ce.tokens[ce.index]
	if !token.IsClass() {
		return fmt.Errorf("index %d: expected 'class', got '%s'", ce.index, token.GetKeyword())
	}
	classComponent.Children = append(classComponent.Children, component.New("keyword", string(token_patterns.CLASS)))
	ce.index++

	// className
	token = ce.tokens[ce.index]
	if !token.IsIdentifier() {
		return fmt.Errorf("index %d: expected class name, got '%s'", ce.index, token.GetValue())
	}
	classComponent.Children = append(classComponent.Children,
		component.NewClassComponent("identifier", token.GetIdentifier()))
	ce.index++

	// '{'
	token = ce.tokens[ce.index]
	if !token.IsOpenBrace() {
		return fmt.Errorf("index %d: expected '{', got '%s'", ce.index, token.GetValue())
	}
	classComponent.Children = append(classComponent.Children, component.New("symbol", "{"))
	ce.index++

	// ClassVarDec
	token = ce.tokens[ce.index]
	for token.IsClassVarDec() {
		ce.componentStack.Push(classComponent)
		if err := ce.compileClassVarDec(); err != nil {
			return err
		}
		token = ce.tokens[ce.index]
		classComponent = ce.componentStack.Pop()
	}

	// Subroutine
	token = ce.tokens[ce.index]
	for token.IsSubroutineDec() {
		ce.componentStack.Push(classComponent)
		if err := ce.compileSubroutine(); err != nil {
			return err
		}
		token = ce.tokens[ce.index]
		classComponent = ce.componentStack.Pop()
	}

	// '}'
	if !ce.tokens[ce.index].IsCloseBrace() {
		return fmt.Errorf("index %d: expected '}', got '%s'", ce.index, ce.tokens[ce.index].GetValue())
	}
	classComponent.Children = append(classComponent.Children, component.New("symbol", "}"))
	ce.index++

	return writeFlush(ce.xmlEncoder, classComponent)
}
