package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	"ny/nand2tetris/compiler/internal/token_patterns"
)

func (ce *CompilationEngine) compileSubroutine() error {
	// subroutine keyword (constructor, function, or method)
	token := ce.tokens[ce.index]
	if !token.IsSubroutineDec() {
		return nil
	}

	ce.symbolTable.Reset()

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	subroutineDecComponent := component.New("subroutineDec", "")
	parentComponent.Children = append(parentComponent.Children, subroutineDecComponent)

	// subroutine  (constructor, function, or method)
	subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("keyword", string(token.GetKeyword())))
	ce.index++

	// return type (void or type)
	token = ce.tokens[ce.index]
	if token.IsVoid() {
		subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("keyword", string(token_patterns.VOID)))
		ce.index++
	} else if token.IsIdentifier() {
		subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("identifier", token.GetValue()))
		ce.index++
	} else if token.IsType() {
		subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("keyword", token.GetValue()))
		ce.index++
	} else {
		return fmt.Errorf("index %d: expected return type (void or type), got '%s'", ce.index, token.GetValue())
	}

	// subroutine name
	token = ce.tokens[ce.index]
	if !token.IsIdentifier() {
		return fmt.Errorf("index %d: expected subroutine name, got '%s'", ce.index, token.GetValue())
	}
	subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.NewSubroutineComponent("identifier", token.GetIdentifier()))
	ce.index++

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("symbol", "("))
	ce.index++

	// parameterList
	ce.componentStack.Push(subroutineDecComponent)
	if err := ce.compileParameterList(); err != nil {
		return err
	}
	subroutineDecComponent = ce.componentStack.Pop()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("symbol", ")"))
	ce.index++

	// subroutineBody
	ce.componentStack.Push(subroutineDecComponent)
	if err := ce.compileSubroutineBody(); err != nil {
		return err
	}
	_ = ce.componentStack.Pop()

	return nil
}
