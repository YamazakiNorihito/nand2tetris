package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	symboltable "ny/nand2tetris/compiler/internal/symbol_table"
	Tokens "ny/nand2tetris/compiler/internal/tokens"
)

func (ce *CompilationEngine) compileSubroutineCall() error {
	token := ce.tokens[ce.index]

	var nextToken Tokens.IToken
	if len(ce.tokens) > ce.index+1 {
		nextToken = ce.tokens[ce.index+1]
	}
	if !token.IsSubroutineCall(nextToken) {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	callFunctionName := ""
	instancename := ""

	// identifier
	var isInstanceCall = false
	if token.IsObjectMethodCall(nextToken) {
		if ce.symbolTable.KindOf(token.GetIdentifier()) == symboltable.NONE {
			parentComponent.Children = append(parentComponent.Children, component.New("identifier", token.GetValue()))
			callFunctionName = token.GetIdentifier()
		} else {
			parentComponent.Children = append(parentComponent.Children,
				component.NewVariableComponent("identifier",
					token.GetIdentifier(),
					component.Category(ce.symbolTable.KindOf(token.GetIdentifier())),
					ce.symbolTable.IndexOf(token.GetIdentifier()),
					component.USED))

			isInstanceCall = true
			instancename = token.GetIdentifier()
			callFunctionName = ce.symbolTable.TypeOf(token.GetIdentifier())
		}
		ce.index++
	} else {
		parentComponent.Children = append(parentComponent.Children, component.New("identifier", token.GetValue()))
		isInstanceCall = true
		callFunctionName = token.GetIdentifier()
		ce.index++
	}

	// '.'
	token = ce.tokens[ce.index]
	if token.IsDot() {
		parentComponent.Children = append(parentComponent.Children, component.New("symbol", "."))
		ce.index++

		token = ce.tokens[ce.index]
		if !token.IsIdentifier() {
			return fmt.Errorf("index %d: expected subroutine name after '.', got '%s'", ce.index, token.GetValue())
		}
		parentComponent.Children = append(parentComponent.Children, component.New("identifier", token.GetIdentifier()))

		callFunctionName += "." + token.GetIdentifier()
		ce.index++
	} else {
		callFunctionName = ce.className + "." + callFunctionName
	}

	if isInstanceCall {
		if instancename == "" {
			// this.method 呼び出し（例：do draw();）
			ce.vmWriter.WritePush("pointer", 0)
		} else {
			// 変数からのメソッド呼び出し（例：square.draw();）
			ce.vmWriter.WritePush(variableKindMemorySegmentMap[ce.symbolTable.KindOf(instancename)], ce.symbolTable.IndexOf(instancename))
		}
	}

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	parentComponent.Children = append(parentComponent.Children, component.New("symbol", "("))
	ce.index++

	// expressionList
	ce.componentStack.Push(parentComponent)
	nArg, err := ce.compileExpressionList()
	if err != nil {
		return err
	}
	parentComponent = ce.componentStack.Pop()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	parentComponent.Children = append(parentComponent.Children, component.New("symbol", ")"))
	ce.index++

	if isInstanceCall {
		ce.vmWriter.WriteCall(callFunctionName, nArg+1)
	} else {
		ce.vmWriter.WriteCall(callFunctionName, nArg)
	}

	return nil
}
