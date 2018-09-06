package expr

import (
	"github.com/levi/holo/token"
)

type Expr interface{}

type Binary struct {
	left      Expr
	operation token.Token
	right     Expr
}

func NewBinary(left Expr, operation token.Token, right Expr) Binary {
	return Binary{
		left,
		operation,
		right,
	}
}

type Grouping struct {
	expression Expr
}

func NewGrouping(expression Expr) Grouping {
	return Grouping{
		expression,
	}
}

type Literal struct {
	value interface{}
}

func NewLiteral(value interface{}) Literal {
	return Literal{
		value,
	}
}

type Unary struct {
	operation token.Token
	right     Expr
}

func NewUnary(operation token.Token, right Expr) Unary {
	return Unary{
		operation,
		right,
	}
}
