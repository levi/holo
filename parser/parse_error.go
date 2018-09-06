package parser

import "github.com/levi/holo/token"

type ParseError struct {
	Token   token.Token
	Message string
}

func NewParseError(token token.Token, message string) *ParseError {
	return &ParseError{
		token,
		message,
	}
}

func (e *ParseError) Error() string {
	return e.Message
}
