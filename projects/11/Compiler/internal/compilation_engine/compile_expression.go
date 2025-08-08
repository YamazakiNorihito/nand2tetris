package compilation_engine

import (
	"fmt"
	"ny/nand2tetris/compiler/internal/component"
	vmwriter "ny/nand2tetris/compiler/internal/vm_writer"
)

var arithmeticCommandOperandMap = map[string]vmwriter.ArithmeticCommand{
	"+": vmwriter.ADD,
	"-": vmwriter.SUB,
	"&": vmwriter.AND,
	"|": vmwriter.OR,
	"<": vmwriter.LT,
	">": vmwriter.GT,
	"=": vmwriter.EQ,
}

// For '*', '/' operators, we need to call Math.multiply and Math.divide

func (ce *CompilationEngine) compileExpression() error {
	token := ce.tokens[ce.index]
	if !token.IsTerm() {
		return nil
	}

	parentComponent := ce.componentStack.Pop()
	defer ce.componentStack.Push(parentComponent)

	expressionComponent := component.New("expression", "")
	parentComponent.Children = append(parentComponent.Children, expressionComponent)

	var op string
	for {
		token = ce.tokens[ce.index]
		if !token.IsTerm() {
			return fmt.Errorf("index %d: expected term, got '%s'", ce.index, token.GetValue())
		}

		// Term
		ce.componentStack.Push(expressionComponent)
		if err := ce.compileTerm(); err != nil {
			return err
		}
		expressionComponent = ce.componentStack.Pop()

		command := arithmeticCommandOperandMap[op]
		if op == "*" {
			ce.vmWriter.WriteCall("Math.multiply", 2, ce.componentStack.Count()+1)
		} else if op == "/" {
			ce.vmWriter.WriteCall("Math.divide", 2, ce.componentStack.Count()+1)
		} else if command != "" {
			ce.vmWriter.WriteArithmetic(command, ce.componentStack.Count()+1)
		}

		// Operator
		token = ce.tokens[ce.index]
		if !token.IsOp() {
			break
		}

		op = token.GetValue()
		expressionComponent.Children = append(expressionComponent.Children, component.New("symbol", op))
		ce.index++
	}
	return nil
}
