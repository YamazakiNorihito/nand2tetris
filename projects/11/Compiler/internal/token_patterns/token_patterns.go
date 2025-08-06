package token_patterns

import (
	"regexp"
)

type KeywordType string

const (
	UNKNOWN     KeywordType = "unknown"
	CLASS       KeywordType = "class"
	METHOD      KeywordType = "method"
	FUNCTION    KeywordType = "function"
	CONSTRUCTOR KeywordType = "constructor"
	INT         KeywordType = "int"
	BOOLEAN     KeywordType = "boolean"
	CHAR        KeywordType = "char"
	VOID        KeywordType = "void"
	VAR         KeywordType = "var"
	STATIC      KeywordType = "static"
	FIELD       KeywordType = "field"
	LET         KeywordType = "let"
	DO          KeywordType = "do"
	IF          KeywordType = "if"
	ELSE        KeywordType = "else"
	WHILE       KeywordType = "while"
	RETURN      KeywordType = "return"
	TRUE        KeywordType = "true"
	FALSE       KeywordType = "false"
	NULL        KeywordType = "null"
	THIS        KeywordType = "this"
)

const (
	/*
		keyword: 'class' | 'constructor' | 'function' |
		'method' | 'field' | 'static' | 'var' | 'int' |
		'char' | 'boolean' | 'void' | 'true' | 'false' |
		'null' | 'this' | 'let' | 'do' | 'if' | 'else' |
		'while' | 'return’
	*/
	KeywordRegexExpression = `\b(class|constructor|function|method|field|static|var|int|char|boolean|void|true|false|null|this|let|do|if|else|while|return)\b`
	// symbol: '{' | '}' | '(' | ')' | '[' | ']' | '. ' | ', ' | '; ' | '+' | '-' | '*' | '/' | '&' | '|' | '<' | '>' | '=' | '~'
	SymbolRegexExpression = `[{}\(\)\[\]\.,;\+\-\*/&|<>=~]`
	// integerConstant: a decimal number in the range 0 ... 32767
	IntegerRegexExpression = `(` +
		`\b0\b|` + // 0
		`\b[1-9]\d{0,3}\b|` + // 1–9999
		`\b1\d{4}\b|` + // 10000–19999
		`\b2\d{4}\b|` + // 20000–29999
		`\b3[0-1][0-9]{3}\b|\b32[0-6][0-9]{2}\b|\b3276[0-7]\b` + // 30000–32767
		`)`
	// StringConstant: '"' a sequence of characters '"' (not including double quote or newline)
	StringRegexExpression = `"[^"\n]*"`
	// identifier: a sequence of letters, digits, and underscore ( '_' ) not starting with a digit.
	IdentifierRegexExpression = `[A-Za-z_][A-Za-z0-9_]*`
)

var TokenPattern = regexp.MustCompile(
	KeywordRegexExpression + "|" +
		SymbolRegexExpression + "|" +
		IntegerRegexExpression + "|" +
		StringRegexExpression + "|" +
		IdentifierRegexExpression,
)
var KeywordPattern = regexp.MustCompile("^" + KeywordRegexExpression + "$")
var SymbolPattern = regexp.MustCompile("^" + SymbolRegexExpression + "$")
var IntegerConstantPattern = regexp.MustCompile("^" + IntegerRegexExpression + "$")
var StringConstantPattern = regexp.MustCompile("^" + StringRegexExpression + "$")
var IdentifierPattern = regexp.MustCompile("^" + IdentifierRegexExpression + "$")

func TokenSplitAndKeepDelimiters(input string) []string {
	var tokens []string
	lastIndex := 0
	locs := TokenPattern.FindAllStringIndex(input, -1)

	for _, loc := range locs {
		start, end := loc[0], loc[1]
		if lastIndex < start {
			tokens = append(tokens, input[lastIndex:start])
		}
		tokens = append(tokens, input[start:end])
		lastIndex = end
	}

	if lastIndex < len(input) {
		tokens = append(tokens, input[lastIndex:])
	}
	return tokens
}
