package expr

import "github.com/levi/holo/token"

type Interpreter interface {
	ToValue() interface{}
}

func (l Literal) ToValue() interface{} {
	return l.value
}

func (g Grouping) ToValue() interface{} {
	return evaluate(g.expression)
}

func (b Binary) ToValue() interface{} {
	left := evaluate(b.left)
	right := evaluate(b.right)

	switch b.operation.TokenType {
	case token.GreaterToken:
		return left.(float64) > right.(float64)
	case token.GreaterEqualToken:
		return left.(float64) >= right.(float64)
	case token.LessToken:
		return left.(float64) < right.(float64)
	case token.LessEqualToken:
		return left.(float64) <= right.(float64)
	case token.PlusToken:
		fLeft, lOk := left.(float64)
		fRight, rOk := right.(float64)

		if lOk && rOk {
			return fLeft + fRight
		}

		sLeft, lOk := left.(string)
		sRight, rOk := right.(string)

		if lOk && rOk {
			return sLeft + sRight
		}
	case token.MinusToken:
		return left.(float64) - right.(float64)
	case token.SlashToken:
		return left.(float64) / right.(float64)
	case token.StarToken:
		return left.(float64) * right.(float64)
	case token.BangEqualToken:
		return !isEqual(left, right)
	case token.EqualEqualToken:
		return isEqual(left, right)
	}

	return nil
}

func (u Unary) ToValue() interface{} {
	right := evaluate(u.right)

	switch u.operation.TokenType {
	case token.BangToken:
		return !isTruthy(right)
	case token.MinusToken:
		return -(right.(float64))
	}

	return nil
}

func evaluate(e Expr) interface{} {
	return e.(Interpreter).ToValue()
}

func isTruthy(i interface{}) bool {
	if i == nil {
		return false
	}
	if b, ok := i.(bool); ok {
		return b
	}
	return true
}

func isEqual(a, b interface{}) bool {
	return a == b
}
