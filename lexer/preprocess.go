package lexer

import (
	"strings"
)

func Preprocess(source string) string {
	l := NewLexer(source)
	tokens := l.Tokenize()

	var preprocessedTokens []Token
	for _, token := range tokens {
		if token.Type != TOKEN_INCLUDE && token.Type != TOKEN_COMMENT {
			preprocessedTokens = append(preprocessedTokens, token)
		}
	}

	var preprocessedSource strings.Builder
	for _, token := range preprocessedTokens {
		preprocessedSource.WriteString(token.Lexeme)
		preprocessedSource.WriteString(" ") // Add a space to separate tokens
	}

	return preprocessedSource.String()
}
