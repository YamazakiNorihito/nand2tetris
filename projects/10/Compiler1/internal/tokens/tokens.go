package tokens

import (
	"fmt"
	"io"
	"ny/nand2tetris/compiler1/internal/jack_tokenizer"
	"ny/nand2tetris/compiler1/internal/token_patterns"
)

type Token struct {
	tokenType  jack_tokenizer.TokenType
	keyword    token_patterns.KeywordType
	symbol     string
	integerVal int
	stringVal  string
	identifier string
}

func (t *Token) GetTokenType() jack_tokenizer.TokenType {
	return t.tokenType
}

func (t *Token) GetKeyword() token_patterns.KeywordType {
	return t.keyword
}
func (t *Token) GetSymbol() string {
	return t.symbol
}
func (t *Token) GetIntegerVal() int {
	return t.integerVal
}
func (t *Token) GetStringVal() string {
	return t.stringVal
}
func (t *Token) GetIdentifier() string {
	return t.identifier
}
func (t *Token) GetName() string {
	return string(t.tokenType)
}
func (t *Token) GetValue() string {
	switch t.tokenType {
	case jack_tokenizer.KEYWORD:
		return string(t.keyword)
	case jack_tokenizer.SYMBOL:
		return t.symbol
	case jack_tokenizer.INT_CONST:
		return fmt.Sprintf("%d", t.integerVal)
	case jack_tokenizer.STRING_CONST:
		return t.stringVal
	case jack_tokenizer.IDENTIFIER:
		return t.identifier
	default:
		return ""
	}
}
func (t *Token) IsClass() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.CLASS
}
func (t *Token) IsMethod() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.METHOD
}
func (t *Token) IsFunction() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.FUNCTION
}
func (t *Token) IsConstructor() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.CONSTRUCTOR
}
func (t *Token) IsField() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.FIELD
}
func (t *Token) IsStatic() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.STATIC
}
func (t *Token) IsVar() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.VAR
}
func (t *Token) IsInt() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.INT
}
func (t *Token) IsChar() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.CHAR
}
func (t *Token) IsBoolean() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.BOOLEAN
}
func (t *Token) IsVoid() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.VOID
}
func (t *Token) IsReturn() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.RETURN
}
func (t *Token) IsIf() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.IF
}
func (t *Token) IsElse() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.ELSE
}
func (t *Token) IsWhile() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.WHILE
}
func (t *Token) IsDo() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.DO
}
func (t *Token) IsLet() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.LET
}
func (t *Token) IsTrue() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.TRUE
}
func (t *Token) IsFalse() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.FALSE
}
func (t *Token) IsNull() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.NULL
}
func (t *Token) IsThis() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && t.keyword == token_patterns.THIS
}
func (t *Token) IsSymbol() bool {
	return t.tokenType == jack_tokenizer.SYMBOL
}
func (t *Token) IsIdentifier() bool {
	return t.tokenType == jack_tokenizer.IDENTIFIER
}
func (t *Token) IsIntConst() bool {
	return t.tokenType == jack_tokenizer.INT_CONST
}
func (t *Token) IsStringConst() bool {
	return t.tokenType == jack_tokenizer.STRING_CONST
}

// term
func (t *Token) IsTerm() bool {
	return t.IsIntConst() || t.IsStringConst() || t.IsKeywordConstant() || t.IsIdentifier() || t.IsOpenParen() || t.IsUnaryOp()
}

// is KeywordConstant
func (t *Token) IsKeywordConstant() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && (t.keyword == token_patterns.TRUE || t.keyword == token_patterns.FALSE || t.keyword == token_patterns.NULL || t.keyword == token_patterns.THIS)
}

// is (
func (t *Token) IsOpenParen() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == "("
}

// is )
func (t *Token) IsCloseParen() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == ")"
}

// is {
func (t *Token) IsOpenBrace() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == "{"
}

// is }
func (t *Token) IsCloseBrace() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == "}"
}

// is [
func (t *Token) IsOpenBracket() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == "["
}

// is ]
func (t *Token) IsCloseBracket() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == "]"
}

// is classVarDec
func (t *Token) IsClassVarDec() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && (t.keyword == token_patterns.FIELD || t.keyword == token_patterns.STATIC)
}

// is subroutineDec
func (t *Token) IsSubroutineDec() bool {
	return t.tokenType == jack_tokenizer.KEYWORD && (t.keyword == token_patterns.CONSTRUCTOR || t.keyword == token_patterns.FUNCTION || t.keyword == token_patterns.METHOD)
}

// is type
func (t *Token) IsType() bool {
	return t.IsInt() || t.IsChar() || t.IsBoolean() || t.IsIdentifier()
}

// is ,
func (t *Token) IsComma() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == ","
}

// is ;
func (t *Token) IsSemicolon() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == ";"
}

// is statement
func (t *Token) IsStatement() bool {
	return t.IsLet() || t.IsIf() || t.IsWhile() || t.IsDo() || t.IsReturn()
}

// is =
func (t *Token) IsEqual() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == "="
}

// is .
func (t *Token) IsDot() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && t.symbol == "."
}

// is item of Array
func (t *Token) IsArrayItem(nextToken *Token) bool {
	return t.IsIdentifier() && nextToken != nil && nextToken.IsOpenBracket()
}

// is subroutine call
func (t *Token) IsSubroutineCall(nextToken *Token) bool {
	return t.IsIdentifier() && (nextToken != nil && (nextToken.IsOpenParen() || nextToken.IsDot()))
}

// is Op
func (t *Token) IsOp() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && (t.symbol == "+" || t.symbol == "-" || t.symbol == "*" || t.symbol == "/" || t.symbol == "&" || t.symbol == "|" || t.symbol == "<" || t.symbol == ">" || t.symbol == "=")
}

// is unaryOp
func (t *Token) IsUnaryOp() bool {
	return t.tokenType == jack_tokenizer.SYMBOL && (t.symbol == "-" || t.symbol == "~")
}

func Build(tokenizer jack_tokenizer.IJackTokenizer) (tokens []Token, err error) {
	if !tokenizer.HasMoreTokens() {
		return nil, nil
	}

	for tokenizer.HasMoreTokens() {
		err := tokenizer.Advance()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to advance Jack tokenizer: %w", err)
		}

		tokenType := tokenizer.TokenType()
		token := Token{tokenType: tokenType}
		switch tokenType {
		case jack_tokenizer.KEYWORD:
			token.keyword = tokenizer.Keyword()
		case jack_tokenizer.SYMBOL:
			token.symbol = tokenizer.Symbol()
		case jack_tokenizer.IDENTIFIER:
			token.identifier = tokenizer.Identifier()
		case jack_tokenizer.INT_CONST:
			token.integerVal = tokenizer.IntVal()
		case jack_tokenizer.STRING_CONST:
			token.stringVal = tokenizer.StringVal()
		default:
			return nil, fmt.Errorf("unknown token type: %s", tokenType)
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
