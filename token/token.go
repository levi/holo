package token

import (
	"fmt"
)

// Token is a scanned token
type Token struct {
	TokenType string
	Lexeme    string
	Literal   interface{}
	Line      int
}

// NewToken allocates a new token
func NewToken(tokenType string, lexeme string, literal interface{}, line int) *Token {
	t := new(Token)
	t.TokenType = tokenType
	t.Lexeme = lexeme
	t.Literal = literal
	t.Line = line
	return t
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %d", t.TokenType, t.Lexeme, t.Literal)
}
