package holo

import (
    "github.com/levi/holo/token"
)

type Expr interface {}

type Binary struct {
    left Expr{}
    operation token.Token
    right Expr{}
}

func NewBinary(left Expr{}, operation token.Token, right Expr{}) *Binary {
    n = new(Binary)
    n.left = left
    n.operation = operation
    n.right = right
    return n
}

type Grouping struct {
    expression Expr{}
}

func NewGrouping(expression Expr{}) *Grouping {
    n = new(Grouping)
    n.expression = expression
    return n
}

type Literal struct {
    value interface{}
}

func NewLiteral(value interface{}) *Literal {
    n = new(Literal)
    n.value = value
    return n
}

type Unary struct {
    operation token.Token
    right Expr{}
}

func NewUnary(operation token.Token, right Expr{}) *Unary {
    n = new(Unary)
    n.operation = operation
    n.right = right
    return n
}

