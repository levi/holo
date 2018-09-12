package expr

import "github.com/levi/holo/token"

type RuntimeError struct {
	Token   token.Token
	Message string
}

func NewRuntimeError(token token.Token, message string) *RuntimeError {
	return &RuntimeError{
		token,
		message,
	}
}

func (e *RuntimeError) Error() string {
	return e.Message
}
