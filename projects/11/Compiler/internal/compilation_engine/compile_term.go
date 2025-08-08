package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	symboltable "ny/nand2tetris/compiler/internal/symbol_table"
	"ny/nand2tetris/compiler/internal/token_patterns"
	Tokens "ny/nand2tetris/compiler/internal/tokens"
	vmwriter "ny/nand2tetris/compiler/internal/vm_writer"
)

func (ce *CompilationEngine) compileTerm() error {
	// term
	token := ce.tokens[ce.index]
	if !token.IsTerm() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	termComponent := component.New("term", "")
	parentComponent.Children = append(parentComponent.Children, termComponent)

	if token.IsIntConst() {
		termComponent.Children = append(termComponent.Children, component.New("integerConstant", token.GetValue()))

		ce.vmWriter.WritePush(vmwriter.CONSTANT, token.GetIntegerVal())

		ce.index++
		return nil
	}

	if token.IsStringConst() {
		termComponent.Children = append(termComponent.Children, component.New("stringConstant", token.GetValue()))

		length := len(token.GetStringVal())
		ce.vmWriter.WritePush(vmwriter.CONSTANT, length)
		ce.vmWriter.WriteCall("String.new", 1)
		for _, char := range token.GetStringVal() {
			ce.vmWriter.WritePush(vmwriter.CONSTANT, int(char))
			ce.vmWriter.WriteCall("String.appendChar", 2)
		}

		ce.index++
		return nil
	}

	if token.IsKeywordConstant() {
		termComponent.Children = append(termComponent.Children, component.New("keyword", token.GetValue()))

		switch token.GetKeyword() {
		case token_patterns.TRUE:
			ce.vmWriter.WritePush(vmwriter.CONSTANT, 0)
			ce.vmWriter.WriteArithmetic(vmwriter.NOT)
		case token_patterns.FALSE, token_patterns.NULL:
			ce.vmWriter.WritePush(vmwriter.CONSTANT, 0)
		case token_patterns.THIS:
			ce.vmWriter.WritePush(vmwriter.POINTER, 0)
		}

		ce.index++
		return nil
	}

	if token.IsUnaryOp() {
		// unary operation '-' or '~'
		unaryOperation := token.GetValue()
		termComponent.Children = append(termComponent.Children, component.New("symbol", unaryOperation))
		ce.index++

		// term
		ce.componentStack.Push(termComponent)
		if err := ce.compileTerm(); err != nil {
			return err
		}
		_ = ce.componentStack.Pop()

		switch unaryOperation {
		case "-":
			ce.vmWriter.WriteArithmetic(vmwriter.NEG)
		case "~":
			ce.vmWriter.WriteArithmetic(vmwriter.NOT)
		}

		return nil
	}

	if token.IsOpenParen() {
		// '('
		termComponent.Children = append(termComponent.Children, component.New("symbol", "("))
		ce.index++

		// expression
		ce.componentStack.Push(termComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		termComponent = ce.componentStack.Pop()

		// ')'
		token = ce.tokens[ce.index]
		if !token.IsCloseParen() {
			return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, ce.tokens[ce.index].GetValue())
		}
		termComponent.Children = append(termComponent.Children, component.New("symbol", ")"))
		ce.index++
		return nil
	}

	// Identifier
	var nextToken Tokens.IToken
	if len(ce.tokens) > ce.index+1 {
		nextToken = ce.tokens[ce.index+1]
	}
	if token.IsArrayItem(nextToken) {
		// array item
		if ce.symbolTable.KindOf(token.GetIdentifier()) == symboltable.NONE {
			return fmt.Errorf("index %d: identifier '%s' is not defined", ce.index, token.GetIdentifier())
		}
		termComponent.Children = append(termComponent.Children,
			component.NewVariableComponent("identifier",
				token.GetIdentifier(),
				component.Category(ce.symbolTable.KindOf(token.GetIdentifier())),
				ce.symbolTable.IndexOf(token.GetIdentifier()),
				component.USED))

		objectName := token.GetIdentifier()

		ce.index++

		// '['
		if !nextToken.IsOpenBracket() {
			return fmt.Errorf("index %d: expected '[', got '%s'", ce.index, nextToken.GetValue())
		}
		termComponent.Children = append(termComponent.Children, component.New("symbol", "["))
		ce.index++

		// expression
		ce.componentStack.Push(termComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		termComponent = ce.componentStack.Pop()

		// ']'
		nextToken = ce.tokens[ce.index]
		if !nextToken.IsCloseBracket() {
			return fmt.Errorf("index %d: expected ']', got '%s'", ce.index, nextToken.GetValue())
		}
		termComponent.Children = append(termComponent.Children, component.New("symbol", "]"))

		memorySegment := variableKindMemorySegmentMap[ce.symbolTable.KindOf(objectName)]
		index := ce.symbolTable.IndexOf(objectName)
		ce.vmWriter.WritePush(memorySegment, index)
		ce.vmWriter.WriteArithmetic(vmwriter.ADD)
		ce.vmWriter.WritePop(vmwriter.POINTER, 1)
		ce.vmWriter.WritePush(vmwriter.THAT, 0)
		ce.index++
	} else if token.IsSubroutineCall(nextToken) {
		// subroutine call
		ce.componentStack.Push(termComponent)
		if err := ce.compileSubroutineCall(); err != nil {
			return err
		}
		_ = ce.componentStack.Pop()
	} else {
		// identifier
		if ce.symbolTable.KindOf(token.GetIdentifier()) == symboltable.NONE {
			return fmt.Errorf("index %d: identifier '%s' is not defined", ce.index, token.GetIdentifier())
		}
		termComponent.Children = append(termComponent.Children,
			component.NewVariableComponent("identifier",
				token.GetIdentifier(),
				component.Category(ce.symbolTable.KindOf(token.GetIdentifier())),
				ce.symbolTable.IndexOf(token.GetIdentifier()),
				component.USED))

		memorySegment := variableKindMemorySegmentMap[ce.symbolTable.KindOf(token.GetIdentifier())]
		ce.vmWriter.WritePush(memorySegment,
			ce.symbolTable.IndexOf(token.GetIdentifier()))
		ce.index++
	}

	return nil
}
