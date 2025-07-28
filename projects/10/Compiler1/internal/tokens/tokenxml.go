package tokens

import (
	"encoding/xml"
	"fmt"
	"io"
	"ny/nand2tetris/compiler1/internal/jack_tokenizer"
	"ny/nand2tetris/compiler1/internal/token_patterns"
	"os"
	"strconv"
	"strings"
)

type node struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}
type root struct {
	XMLName xml.Name `xml:"tokens"`
	Tokens  []node   `xml:",any"`
}

func DumpTokensToXML(tokens []IToken, file *os.File) {
	xmlTokens := root{
		Tokens: make([]node, len(tokens)),
	}
	for i, token := range tokens {
		xmlTokens.Tokens[i] = node{
			XMLName: xml.Name{Local: token.GetName()},
			Value:   paddingFunc(token.GetValue(), " "),
		}
	}

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(xmlTokens); err != nil {
		panic(fmt.Errorf("failed to encode tokens to XML: %w", err))
	}
}

func paddingFunc(v string, padding string) string {
	return fmt.Sprintf("%s%s%s", padding, v, padding)
}

func LoadTokensFromXML(r io.Reader) (tokens []IToken) {
	decoder := xml.NewDecoder(r)
	var xmlTokens root
	if err := decoder.Decode(&xmlTokens); err != nil {
		panic(fmt.Errorf("failed to decode XML: %w", err))
	}

	tokens = make([]IToken, len(xmlTokens.Tokens))
	for i, xmlToken := range xmlTokens.Tokens {
		tokenType := stringToTokenType(xmlToken.XMLName.Local)

		value := strings.TrimPrefix(strings.TrimSuffix(xmlToken.Value, " "), " ")
		switch tokenType {
		case jack_tokenizer.KEYWORD:
			tokens[i] = &Token{
				tokenType: tokenType,
				keyword:   keywordTypeFromString(value),
			}
		case jack_tokenizer.SYMBOL:
			tokens[i] = &Token{
				tokenType: tokenType,
				symbol:    value,
			}
		case jack_tokenizer.IDENTIFIER:
			tokens[i] = &Token{
				tokenType:  tokenType,
				identifier: value,
			}
		case jack_tokenizer.INT_CONST:
			intVal, err := strconv.Atoi(value)
			if err != nil {
				panic(fmt.Errorf("invalid integer constant %s: %w", value, err))
			}
			tokens[i] = &Token{
				tokenType:  tokenType,
				integerVal: intVal,
			}
		case jack_tokenizer.STRING_CONST:
			tokens[i] = &Token{
				tokenType: tokenType,
				stringVal: value,
			}
		default:
			panic(fmt.Sprintf("unknown token type: %s", xmlToken.XMLName.Local))
		}

	}
	return tokens
}

func stringToTokenType(s string) jack_tokenizer.TokenType {
	switch s {
	case "keyword":
		return jack_tokenizer.KEYWORD
	case "symbol":
		return jack_tokenizer.SYMBOL
	case "identifier":
		return jack_tokenizer.IDENTIFIER
	case "integerConstant":
		return jack_tokenizer.INT_CONST
	case "stringConstant":
		return jack_tokenizer.STRING_CONST
	default:
		panic(fmt.Sprintf("unknown token type: %s", s))
	}
}

func keywordTypeFromString(s string) token_patterns.KeywordType {
	switch s {
	case "class":
		return token_patterns.CLASS
	case "method":
		return token_patterns.METHOD
	case "function":
		return token_patterns.FUNCTION
	case "constructor":
		return token_patterns.CONSTRUCTOR
	case "int":
		return token_patterns.INT
	case "boolean":
		return token_patterns.BOOLEAN
	case "char":
		return token_patterns.CHAR
	case "void":
		return token_patterns.VOID
	case "var":
		return token_patterns.VAR
	case "static":
		return token_patterns.STATIC
	case "field":
		return token_patterns.FIELD
	case "let":
		return token_patterns.LET
	case "do":
		return token_patterns.DO
	case "if":
		return token_patterns.IF
	case "else":
		return token_patterns.ELSE
	case "while":
		return token_patterns.WHILE
	case "return":
		return token_patterns.RETURN
	case "true":
		return token_patterns.TRUE
	case "false":
		return token_patterns.FALSE
	case "null":
		return token_patterns.NULL
	case "this":
		return token_patterns.THIS
	default:
		panic("unknown keyword type is not valid")
	}
}
