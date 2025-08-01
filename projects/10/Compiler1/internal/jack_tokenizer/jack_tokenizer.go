package jack_tokenizer

import (
	"io"
	"ny/nand2tetris/compiler1/internal/token_patterns"
	"strconv"
	"strings"
)

type TokenType string

const (
	UNKNOWN      TokenType = "unknown"
	KEYWORD      TokenType = "keyword"
	SYMBOL       TokenType = "symbol"
	IDENTIFIER   TokenType = "identifier"
	INT_CONST    TokenType = "integerConstant"
	STRING_CONST TokenType = "stringConstant"
)

type IJackTokenizer interface {
	HasMoreTokens() bool
	Advance() error
	TokenType() TokenType
	Keyword() token_patterns.KeywordType
	Symbol() string
	Identifier() string
	IntVal() int
	StringVal() string

	getCodeLines() []string
}

type JackTokenizer struct {
	lines        []string
	currentIndex int

	tokens      []string
	token       string
	tokenType   TokenType
	keywordType token_patterns.KeywordType
}

func NewAnalysis(r io.Reader) (IJackTokenizer, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return &JackTokenizer{
			lines:        []string{},
			currentIndex: -1,
		}, nil
	}

	codelines := getCodeLines(string(content))
	return &JackTokenizer{
		lines:        codelines,
		currentIndex: -1,
	}, nil
}

func (j *JackTokenizer) HasMoreTokens() bool {
	return j.currentIndex+1 < len(j.lines) || len(j.tokens) > 0
}

func (j *JackTokenizer) Advance() error {
	if !j.HasMoreTokens() {
		return io.EOF
	}

	if len(j.tokens) > 0 {
		token := j.nextToken()
		j.setToken(token)
		return nil
	}

	j.currentIndex++
	line := j.lines[j.currentIndex]
	tokens := token_patterns.TokenSplitAndKeepDelimiters(line)

	// Remove empty values from tokens
	var nonEmptyTokens []string
	for _, token := range tokens {
		if strings.TrimSpace(token) != "" {
			nonEmptyTokens = append(nonEmptyTokens, token)
		}
	}
	j.tokens = nonEmptyTokens
	if len(j.tokens) == 0 {
		return io.EOF
	}
	token := j.nextToken()
	j.setToken(token)

	return nil
}

func (j *JackTokenizer) TokenType() TokenType {
	return j.tokenType
}
func (j *JackTokenizer) Keyword() token_patterns.KeywordType {
	if j.tokenType != KEYWORD {
		panic("Current token is not a keyword")
	}
	return j.keywordType
}
func (j *JackTokenizer) Symbol() string {
	if j.tokenType != SYMBOL {
		panic("Current token is not a symbol")
	}
	return j.token
}
func (j *JackTokenizer) Identifier() string {
	if j.tokenType != IDENTIFIER {
		panic("Current token is not an identifier")
	}
	return j.token
}
func (j *JackTokenizer) IntVal() int {
	if j.tokenType != INT_CONST {
		panic("Current token is not an integer constant")
	}
	val, err := strconv.Atoi(j.token)
	if err != nil {
		panic("Invalid integer constant: " + j.token)
	}
	return val
}
func (j *JackTokenizer) StringVal() string {
	if j.tokenType != STRING_CONST {
		panic("Current token is not a string constant")
	}
	return strings.Trim(j.token, "\"")
}

func (j *JackTokenizer) getCodeLines() []string {
	return j.lines
}

func getCodeLines(content string) []string {
	codeWithoutComments := removeCommentBlocks(content)
	cleanedCode := strings.ReplaceAll(codeWithoutComments, "\r\n", "\n")
	if cleanedCode == "" {
		return []string{}
	}
	return strings.Split(cleanedCode, "\n")
}

func removeCommentBlocks(content string) string {
	inString := false
	inBlockComment := false
	var result []rune

	for i := 0; i < len(content); {

		//  "/* comment */" の前後で文字列を認識するために、文字列内ではコメントを無視する
		if !inString && !inBlockComment && i+1 < len(content) && content[i] == '/' && content[i+1] == '*' {
			if i+2 < len(content) && content[i+2] == '*' {
				inBlockComment = true
				i += 3
				continue
			} else {
				inBlockComment = true
				i += 2
				continue
			}
		}
		if inBlockComment {
			if i+1 < len(content) && content[i] == '*' && content[i+1] == '/' {
				inBlockComment = false
				i += 2
			} else {
				i++
			}
			continue
		}

		//  "// comment" の前後で文字列を認識するために、文字列内ではコメントを無視する
		if !inString && i+1 < len(content) && content[i] == '/' && content[i+1] == '/' {
			// Skip to the end of line
			for i < len(content) && content[i] != '\n' {
				i++
			}
			continue
		}

		// "string" の文字列の処理をする。
		if content[i] == '"' {
			result = append(result, rune(content[i]))
			i++
			inString = !inString
			continue
		}

		result = append(result, rune(content[i]))
		i++
	}

	isEmpty := func(s string) bool {
		return strings.TrimSpace(s) == ""
	}
	lines := strings.Split(string(result), "\n")
	var cleaned []string
	for _, line := range lines {
		if isEmpty(line) {
			continue
		}
		cleaned = append(cleaned, line)
	}

	return strings.ReplaceAll(strings.Join(cleaned, "\n"), "\r\n", "\n")
}

func (j *JackTokenizer) setToken(token string) {
	j.clearToken()

	j.token = token

	switch {
	case token_patterns.KeywordPattern.MatchString(token):
		j.tokenType = KEYWORD
		j.keywordType = token_patterns.KeywordType(token)
	case token_patterns.SymbolPattern.MatchString(token):
		j.tokenType = SYMBOL
	case token_patterns.IntegerConstantPattern.MatchString(token):
		j.tokenType = INT_CONST
	case token_patterns.StringConstantPattern.MatchString(token):
		j.tokenType = STRING_CONST
	case token_patterns.IdentifierPattern.MatchString(token):
		j.tokenType = IDENTIFIER
	default:
		panic("Unknown token type: " + token)
	}
}

func (j *JackTokenizer) nextToken() string {
	if len(j.tokens) == 0 {
		panic("No tokens available in the current line")
	}
	first := j.tokens[0]
	j.tokens = j.tokens[1:]
	return first
}

func (j *JackTokenizer) clearToken() {
	j.token = ""
	j.tokenType = UNKNOWN
	j.keywordType = token_patterns.UNKNOWN
}
