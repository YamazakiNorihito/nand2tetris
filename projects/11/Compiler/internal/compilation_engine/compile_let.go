package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	symboltable "ny/nand2tetris/compiler/internal/symbol_table"
	"ny/nand2tetris/compiler/internal/token_patterns"
	Tokens "ny/nand2tetris/compiler/internal/tokens"
	vmwriter "ny/nand2tetris/compiler/internal/vm_writer"
)

func (ce *CompilationEngine) compileLet() error {
	// let
	token := ce.tokens[ce.index]
	if !token.IsLet() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	letStatementComponent := component.New("letStatement", "")
	parentComponent.Children = append(parentComponent.Children, letStatementComponent)

	letStatementComponent.Children = append(letStatementComponent.Children, component.New("keyword", string(token_patterns.LET)))
	ce.index++

	// variable name
	token = ce.tokens[ce.index]

	if !token.IsIdentifier() {
		return fmt.Errorf("index %d: expected variable name, got '%s'", ce.index, token.GetValue())
	}

	if ce.symbolTable.KindOf(token.GetIdentifier()) == symboltable.NONE {
		return fmt.Errorf("index %d: variable '%s' is not defined", ce.index, token.GetIdentifier())
	}

	var nextToken Tokens.IToken
	if len(ce.tokens) > ce.index+1 {
		nextToken = ce.tokens[ce.index+1]
	}

	// '['
	isArray := token.IsArrayItem(nextToken)
	letName := token.GetIdentifier()
	if isArray {
		letStatementComponent.Children = append(letStatementComponent.Children,
			component.NewVariableComponent("identifier",
				token.GetIdentifier(),
				component.Category(ce.symbolTable.KindOf(token.GetIdentifier())),
				ce.symbolTable.IndexOf(token.GetIdentifier()),
				component.USED))

		objectName := token.GetIdentifier()
		ce.index++

		// '['
		letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "["))
		ce.index++

		// expression
		ce.componentStack.Push(letStatementComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		letStatementComponent = ce.componentStack.Pop()

		// ']'
		token = ce.tokens[ce.index]
		if !token.IsCloseBracket() {
			return fmt.Errorf("index %d: expected ']', got '%s'", ce.index, token.GetValue())
		}
		letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "]"))

		memorySegment := variableKindMemorySegmentMap[ce.symbolTable.KindOf(objectName)]
		index := ce.symbolTable.IndexOf(objectName)
		ce.vmWriter.WritePush(memorySegment, index) // push base address of array
		ce.vmWriter.WriteArithmetic(vmwriter.ADD)   // add index: a + i
	} else {
		letStatementComponent.Children = append(letStatementComponent.Children,
			component.NewVariableComponent("identifier",
				token.GetIdentifier(),
				component.Category(ce.symbolTable.KindOf(token.GetIdentifier())),
				ce.symbolTable.IndexOf(token.GetIdentifier()),
				component.USED))
	}

	ce.index++

	// '='
	token = ce.tokens[ce.index]
	if !token.IsEqual() {
		return fmt.Errorf("index %d: expected '=', got '%s'", ce.index, token.GetValue())
	}
	letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "="))
	ce.index++

	// expression
	ce.componentStack.Push(letStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	letStatementComponent = ce.componentStack.Pop()

	// ;
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", ";"))
	ce.index++

	memorySegmentType := variableKindMemorySegmentMap[ce.symbolTable.KindOf(letName)]
	if isArray {
		ce.vmWriter.WritePop(vmwriter.TEMP, 0)
		ce.vmWriter.WritePop(vmwriter.POINTER, 1)
		ce.vmWriter.WritePush(vmwriter.TEMP, 0)
		ce.vmWriter.WritePop(vmwriter.THAT, 0)
	} else {
		ce.vmWriter.WritePop(memorySegmentType, ce.symbolTable.IndexOf(letName))
	}

	return nil
}
