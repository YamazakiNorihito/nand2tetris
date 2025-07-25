package compilation_engine

import (
	"encoding/xml"
	"fmt"
	"ny/nand2tetris/compiler1/internal/component"
	"ny/nand2tetris/compiler1/internal/token_patterns"
	Tokens "ny/nand2tetris/compiler1/internal/tokens"
)

type ICompilationEngine interface {
	CompileClass() error
}

type CompilationEngine struct {
	tokens          []Tokens.Token
	xmlEncoder      *xml.Encoder
	component       *component.Component
	index           int
	parentComponent *component.Component
	parentStack     []*component.Component
}

func New(tokens []Tokens.Token, xmlEncoder *xml.Encoder) (ICompilationEngine, error) {
	return &CompilationEngine{tokens: tokens, xmlEncoder: xmlEncoder, index: 0}, nil
}

/*
class ::= 'class' className '{' classVarDec* subroutineDec* '}'
*/
func (ce *CompilationEngine) CompileClass() error {
	classComponent := component.New("class", "")
	ce.component = classComponent

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
	classComponent.Children = append(classComponent.Children, component.New("identifier", token.GetIdentifier()))
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
		ce.pushParent(classComponent)
		if err := ce.compileClassVarDec(); err != nil {
			return err
		}
		token = ce.tokens[ce.index]
		ce.popParent()
	}

	// Subroutine
	token = ce.tokens[ce.index]
	for token.IsSubroutineDec() {
		ce.pushParent(classComponent)
		if err := ce.compileSubroutine(); err != nil {
			return err
		}
		token = ce.tokens[ce.index]
		ce.popParent()
	}

	// '}'
	if !ce.tokens[ce.index].IsCloseBrace() {
		return fmt.Errorf("index %d: expected '}', got '%s'", ce.index, ce.tokens[ce.index].GetValue())
	}
	classComponent.Children = append(classComponent.Children, component.New("symbol", "}"))
	ce.index++

	return ce.writeFlush()
}

func (ce *CompilationEngine) compileClassVarDec() error {
	// static or field
	token := ce.tokens[ce.index]
	if !token.IsClassVarDec() {
		return nil
	}
	classVarDecComponent := component.New("classVarDec", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, classVarDecComponent)

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

func (ce *CompilationEngine) compileSubroutine() error {
	// subroutine keyword (constructor, function, or method)
	token := ce.tokens[ce.index]
	if !token.IsSubroutineDec() {
		return nil
	}
	subroutineDecComponent := component.New("subroutineDec", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, subroutineDecComponent)

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
	subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("identifier", token.GetIdentifier()))
	ce.index++

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("symbol", "("))
	ce.index++

	// parameterList
	ce.pushParent(subroutineDecComponent)
	if err := ce.compileParameterList(); err != nil {
		return err
	}
	ce.popParent()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	subroutineDecComponent.Children = append(subroutineDecComponent.Children, component.New("symbol", ")"))
	ce.index++

	// subroutineBody
	ce.pushParent(subroutineDecComponent)
	if err := ce.compileSubroutineBody(); err != nil {
		return err
	}
	ce.popParent()

	return nil
}

func (ce *CompilationEngine) compileParameterList() error {
	paramListComponent := component.New("parameterList", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, paramListComponent)

	token := ce.tokens[ce.index]
	if !token.IsType() {
		return nil
	}

	for {
		token = ce.tokens[ce.index]
		if !token.IsType() {
			return fmt.Errorf("index %d: expected type after ',', got '%s'", ce.index, token.GetValue())
		}

		// Type token (e.g., int, char, boolean, className)
		paramListComponent.Children = append(paramListComponent.Children, component.New("keyword", token.GetValue()))
		ce.index++

		// Parameter name
		token = ce.tokens[ce.index]
		if !token.IsIdentifier() {
			return fmt.Errorf("index %d: expected parameter name, got '%s'", ce.index, token.GetValue())
		}
		paramListComponent.Children = append(paramListComponent.Children, component.New("identifier", token.GetIdentifier()))
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

func (ce *CompilationEngine) compileSubroutineBody() error {
	// '{'
	token := ce.tokens[ce.index]
	if !token.IsOpenBrace() {
		return fmt.Errorf("expected '{', got '%s'", token.GetValue())
	}

	subroutineBodyComponent := component.New("subroutineBody", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, subroutineBodyComponent)

	subroutineBodyComponent.Children = append(subroutineBodyComponent.Children, component.New("symbol", "{"))
	ce.index++

	// Compile variable declarations
	token = ce.tokens[ce.index]
	for token.IsVar() {
		parent := ce.parentComponent
		ce.parentComponent = subroutineBodyComponent

		if err := ce.compileVarDec(); err != nil {
			return err
		}
		token = ce.tokens[ce.index]
		ce.parentComponent = parent
	}

	// Compile statements
	ce.pushParent(subroutineBodyComponent)
	if err := ce.compileStatements(); err != nil {
		return err
	}
	ce.popParent()

	// '}'
	token = ce.tokens[ce.index]
	if !token.IsCloseBrace() {
		return fmt.Errorf("expected '}', got '%s'", token.GetValue())
	}
	subroutineBodyComponent.Children = append(subroutineBodyComponent.Children, component.New("symbol", "}"))
	ce.index++

	ce.parentComponent = nil
	return nil
}

func (ce *CompilationEngine) compileVarDec() error {
	// var
	token := ce.tokens[ce.index]
	if !token.IsVar() {
		return nil
	}

	varDec := component.New("varDec", "")
	if ce.parentComponent != nil {
		ce.parentComponent.Children = append(ce.parentComponent.Children, varDec)
	}

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

func (ce *CompilationEngine) compileStatements() error {
	statementsComponent := component.New("statements", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, statementsComponent)

	token := ce.tokens[ce.index]
	if !token.IsStatement() {
		return nil
	}

	for token.IsStatement() {
		parent := ce.parentComponent
		ce.parentComponent = statementsComponent

		switch {
		case token.IsLet():
			if err := ce.compileLet(); err != nil {
				return err
			}
		case token.IsIf():
			if err := ce.compileIf(); err != nil {
				return err
			}
		case token.IsWhile():
			if err := ce.compileWhile(); err != nil {
				return err
			}
		case token.IsDo():
			if err := ce.compileDo(); err != nil {
				return err
			}
		case token.IsReturn():
			if err := ce.compileReturn(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("index %d: unexpected statement type: %s", ce.index, token.GetValue())
		}
		token = ce.tokens[ce.index]
		ce.parentComponent = parent
	}
	return nil
}

func (ce *CompilationEngine) compileLet() error {
	// let
	token := ce.tokens[ce.index]
	if !token.IsLet() {
		return nil
	}

	letStatementComponent := component.New("letStatement", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, letStatementComponent)

	letStatementComponent.Children = append(letStatementComponent.Children, component.New("keyword", string(token_patterns.LET)))
	ce.index++

	// variable name
	token = ce.tokens[ce.index]
	if !token.IsIdentifier() {
		return fmt.Errorf("index %d: expected variable name, got '%s'", ce.index, token.GetValue())
	}
	letStatementComponent.Children = append(letStatementComponent.Children, component.New("identifier", token.GetIdentifier()))
	ce.index++

	// '['
	token = ce.tokens[ce.index]
	if token.IsOpenBracket() {
		letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "["))
		ce.index++

		// expression
		ce.pushParent(letStatementComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		ce.popParent()

		// ']'
		token = ce.tokens[ce.index]
		if !token.IsCloseBracket() {
			return fmt.Errorf("index %d: expected ']', got '%s'", ce.index, token.GetValue())
		}
		letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "]"))
		ce.index++
	}

	// '='
	token = ce.tokens[ce.index]
	if !token.IsEqual() {
		return fmt.Errorf("index %d: expected '=', got '%s'", ce.index, token.GetValue())
	}
	letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", "="))
	ce.index++

	// expression
	ce.pushParent(letStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	ce.popParent()

	// ;
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	letStatementComponent.Children = append(letStatementComponent.Children, component.New("symbol", ";"))
	ce.index++

	return nil
}

func (ce *CompilationEngine) compileIf() error {
	// if
	token := ce.tokens[ce.index]
	if !token.IsIf() {
		return nil
	}
	ifStatementComponent := component.New("ifStatement", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, ifStatementComponent)
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("keyword", string(token_patterns.IF)))
	ce.index++

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "("))
	ce.index++

	// expression
	ce.pushParent(ifStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	ce.popParent()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", ")"))
	ce.index++

	// '{'
	token = ce.tokens[ce.index]
	if !token.IsOpenBrace() {
		return fmt.Errorf("index %d: expected '{', got '%s'", ce.index, token.GetValue())
	}
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "{"))
	ce.index++

	// statements
	ce.pushParent(ifStatementComponent)
	if err := ce.compileStatements(); err != nil {
		return err
	}
	ce.popParent()

	// '}'
	token = ce.tokens[ce.index]
	if !token.IsCloseBrace() {
		return fmt.Errorf("index %d: expected '}', got '%s'", ce.index, token.GetValue())
	}
	ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "}"))
	ce.index++

	// else
	token = ce.tokens[ce.index]
	if token.IsElse() {
		ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("keyword", string(token_patterns.ELSE)))
		ce.index++

		// '{'
		token = ce.tokens[ce.index]
		if !token.IsOpenBrace() {
			return fmt.Errorf("index %d: expected '{', got '%s'", ce.index, token.GetValue())
		}
		ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "{"))
		ce.index++

		// statements
		ce.pushParent(ifStatementComponent)
		if err := ce.compileStatements(); err != nil {
			return err
		}
		ce.popParent()

		// '}'
		token = ce.tokens[ce.index]
		if !token.IsCloseBrace() {
			return fmt.Errorf("index %d: expected '}', got '%s'", ce.index, token.GetValue())
		}
		ifStatementComponent.Children = append(ifStatementComponent.Children, component.New("symbol", "}"))
		ce.index++
	}
	return nil
}

func (ce *CompilationEngine) compileWhile() error {
	// while
	token := ce.tokens[ce.index]
	if !token.IsWhile() {
		return nil
	}

	whileStatementComponent := component.New("whileStatement", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, whileStatementComponent)

	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("keyword", string(token_patterns.WHILE)))
	ce.index++

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("symbol", "("))
	ce.index++

	// expression
	ce.pushParent(whileStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	ce.popParent()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("symbol", ")"))
	ce.index++

	// '{'
	token = ce.tokens[ce.index]
	if !token.IsOpenBrace() {
		return fmt.Errorf("index %d: expected '{', got '%s'", ce.index, token.GetValue())
	}
	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("symbol", "{"))
	ce.index++

	// statements
	ce.pushParent(whileStatementComponent)
	if err := ce.compileStatements(); err != nil {
		return err
	}
	ce.popParent()

	// '}'
	token = ce.tokens[ce.index]
	if !token.IsCloseBrace() {
		return fmt.Errorf("index %d: expected '}', got '%s'", ce.index, token.GetValue())
	}
	whileStatementComponent.Children = append(whileStatementComponent.Children, component.New("symbol", "}"))
	ce.index++

	return nil
}

func (ce *CompilationEngine) compileDo() error {
	// do
	token := ce.tokens[ce.index]
	if !token.IsDo() {
		return nil
	}

	doStatementComponent := component.New("doStatement", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, doStatementComponent)

	doStatementComponent.Children = append(doStatementComponent.Children, component.New("keyword", string(token_patterns.DO)))
	ce.index++

	// subroutine call
	ce.pushParent(doStatementComponent)
	if err := ce.compileSubroutineCall(); err != nil {
		return err
	}
	ce.popParent()

	// ';'
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	doStatementComponent.Children = append(doStatementComponent.Children, component.New("symbol", ";"))
	ce.index++

	return nil
}

func (ce *CompilationEngine) compileReturn() error {
	// return
	token := ce.tokens[ce.index]
	if !token.IsReturn() {
		return nil
	}

	returnStatementComponent := component.New("returnStatement", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, returnStatementComponent)

	returnStatementComponent.Children = append(returnStatementComponent.Children, component.New("keyword", string(token_patterns.RETURN)))
	ce.index++

	// expression
	ce.pushParent(returnStatementComponent)
	if err := ce.compileExpression(); err != nil {
		return err
	}
	ce.popParent()

	// ;
	token = ce.tokens[ce.index]
	if !token.IsSemicolon() {
		return fmt.Errorf("index %d: expected ';', got '%s'", ce.index, token.GetValue())
	}
	returnStatementComponent.Children = append(returnStatementComponent.Children, component.New("symbol", ";"))
	ce.index++

	return nil
}

func (ce *CompilationEngine) compileExpression() error {
	token := ce.tokens[ce.index]
	if !token.IsTerm() {
		return nil
	}

	expressionComponent := component.New("expression", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, expressionComponent)

	for {
		token = ce.tokens[ce.index]
		if !token.IsTerm() {
			return fmt.Errorf("index %d: expected term, got '%s'", ce.index, token.GetValue())
		}

		// Term
		ce.pushParent(expressionComponent)
		if err := ce.compileTerm(); err != nil {
			return err
		}
		ce.popParent()

		// Operator
		token = ce.tokens[ce.index]
		if !token.IsOp() {
			break
		}
		expressionComponent.Children = append(expressionComponent.Children, component.New("symbol", token.GetValue()))
		ce.index++
	}
	return nil
}

func (ce *CompilationEngine) compileTerm() error {
	// term
	token := ce.tokens[ce.index]
	if !token.IsTerm() {
		return nil
	}

	termComponent := component.New("term", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, termComponent)

	if token.IsIntConst() {
		termComponent.Children = append(termComponent.Children, component.New("integerConstant", token.GetValue()))
		ce.index++
		return nil
	}

	if token.IsStringConst() {
		termComponent.Children = append(termComponent.Children, component.New("stringConstant", token.GetValue()))
		ce.index++
		return nil
	}

	if token.IsKeywordConstant() {
		termComponent.Children = append(termComponent.Children, component.New("keyword", token.GetValue()))
		ce.index++
		return nil
	}

	if token.IsUnaryOp() {
		// unary operation '-' or '~'
		termComponent.Children = append(termComponent.Children, component.New("symbol", token.GetValue()))
		ce.index++

		// term
		ce.pushParent(termComponent)
		if err := ce.compileTerm(); err != nil {
			return err
		}
		ce.popParent()
		return nil
	}

	if token.IsOpenParen() {
		// '('
		termComponent.Children = append(termComponent.Children, component.New("symbol", "("))
		ce.index++

		// expression
		ce.pushParent(termComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		ce.popParent()

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
	var nextToken Tokens.Token
	if len(ce.tokens) > ce.index+1 {
		nextToken = ce.tokens[ce.index+1]
	}
	if token.IsArrayItem(&nextToken) {
		// array item
		termComponent.Children = append(termComponent.Children, component.New("identifier", token.GetIdentifier()))
		ce.index++

		// '['
		if !nextToken.IsOpenBracket() {
			return fmt.Errorf("index %d: expected '[', got '%s'", ce.index, nextToken.GetValue())
		}
		termComponent.Children = append(termComponent.Children, component.New("symbol", "["))
		ce.index++

		// expression
		ce.pushParent(termComponent)
		if err := ce.compileExpression(); err != nil {
			return err
		}
		ce.popParent()

		// ']'
		nextToken = ce.tokens[ce.index]
		if !nextToken.IsCloseBracket() {
			return fmt.Errorf("index %d: expected ']', got '%s'", ce.index, nextToken.GetValue())
		}
		termComponent.Children = append(termComponent.Children, component.New("symbol", "]"))
		ce.index++
	} else if token.IsSubroutineCall(&nextToken) {
		// subroutine call
		ce.pushParent(termComponent)
		if err := ce.compileSubroutineCall(); err != nil {
			return err
		}
		ce.popParent()
	} else {
		// identifier
		termComponent.Children = append(termComponent.Children, component.New("identifier", token.GetIdentifier()))
		ce.index++
	}

	return nil
}

func (ce *CompilationEngine) compileExpressionList() (int, error) {
	expressionListComponent := component.New("expressionList", "")
	ce.parentComponent.Children = append(ce.parentComponent.Children, expressionListComponent)

	token := ce.tokens[ce.index]
	if !token.IsTerm() {
		return 0, nil
	}

	count := 0
	for {
		// expression
		ce.pushParent(expressionListComponent)
		if err := ce.compileExpression(); err != nil {
			return 0, err
		}
		ce.popParent()
		count++

		// ','
		token = ce.tokens[ce.index]
		if !token.IsComma() {
			break
		}
		expressionListComponent.Children = append(expressionListComponent.Children, component.New("symbol", ","))
		ce.index++
	}

	return count, nil
}

func (ce *CompilationEngine) compileSubroutineCall() error {
	token := ce.tokens[ce.index]

	var nextToken Tokens.Token
	if len(ce.tokens) > ce.index+1 {
		nextToken = ce.tokens[ce.index+1]
	}
	if !token.IsSubroutineCall(&nextToken) {
		return nil
	}

	ce.parentComponent.Children = append(ce.parentComponent.Children, component.New("identifier", token.GetValue()))
	ce.index++

	// '.'
	token = ce.tokens[ce.index]
	if token.IsDot() {
		ce.parentComponent.Children = append(ce.parentComponent.Children, component.New("symbol", "."))
		ce.index++

		token = ce.tokens[ce.index]
		if !token.IsIdentifier() {
			return fmt.Errorf("index %d: expected subroutine name after '.', got '%s'", ce.index, token.GetValue())
		}
		ce.parentComponent.Children = append(ce.parentComponent.Children, component.New("identifier", token.GetIdentifier()))
		ce.index++
	}

	// '('
	token = ce.tokens[ce.index]
	if !token.IsOpenParen() {
		return fmt.Errorf("index %d: expected '(', got '%s'", ce.index, token.GetValue())
	}
	ce.parentComponent.Children = append(ce.parentComponent.Children, component.New("symbol", "("))
	ce.index++

	// expressionList
	ce.pushParent(ce.parentComponent)
	_, err := ce.compileExpressionList()
	if err != nil {
		return err
	}
	ce.popParent()

	// ')'
	token = ce.tokens[ce.index]
	if !token.IsCloseParen() {
		return fmt.Errorf("index %d: expected ')', got '%s'", ce.index, token.GetValue())
	}
	ce.parentComponent.Children = append(ce.parentComponent.Children, component.New("symbol", ")"))
	ce.index++

	return nil
}

func (ce *CompilationEngine) pushParent(c *component.Component) {
	ce.parentStack = append(ce.parentStack, c)
	ce.parentComponent = c
}

func (ce *CompilationEngine) popParent() {
	if len(ce.parentStack) > 0 {
		ce.parentStack = ce.parentStack[:len(ce.parentStack)-1]
		if len(ce.parentStack) > 0 {
			ce.parentComponent = ce.parentStack[len(ce.parentStack)-1]
		} else {
			ce.parentComponent = nil
		}
	}
}

func (ce *CompilationEngine) writeFlush() error {
	ce.xmlEncoder.Indent("", "  ")
	if err := ce.xmlEncoder.Encode(ce.component); err != nil {
		return fmt.Errorf("index %d: failed to encode XML: %w", ce.index, err)
	}
	return nil
}
