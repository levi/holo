package token

import (
	"fmt"
)

// Token is a scanned token
type Token struct {
	tokenType string
	lexeme    string
	literal   string
	line      int
}

// NewToken allocates a new token
func NewToken(tokenType string, lexeme string, literal string, line int) *Token {
	t := new(Token)
	t.tokenType = tokenType
	t.lexeme = lexeme
	t.literal = literal
	t.line = line
	return t
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.tokenType, t.lexeme, t.literal)
}
