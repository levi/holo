package main

import "fmt"

type AstPrinter interface {
	ToString() string
}

func (b *Binary) ToString() string {
	return parenthesize(b.operation.Lexeme, b.left, b.right)
}

func (g *Grouping) ToString() string {
	return parenthesize("group", g.expression)
}

func (l *Literal) ToString() string {
	if l.value == nil {
		return "null"
	}
	return fmt.Sprintf("%v", l.value)
}

func (u *Unary) ToString() string {
	return parenthesize(u.operation.Lexeme, u.right)
}

func parenthesize(name string, exprs ...Expr) string {
	out := "(" + name
	for _, expr := range exprs {
		out += " "
		out += expr.(AstPrinter).ToString()
	}
	out += ")"
	return out
}
