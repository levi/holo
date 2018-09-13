package expr

import (
	"fmt"

	"github.com/levi/holo/token"
)

type Interpreter interface {
	ToValue() (interface{}, error)
}

func Interpret(statements []Stmt) error {
	for _, s := range statements {
		err := execute(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func execute(statement Stmt) error {
	_, err := statement.(Interpreter).ToValue()
	return err
}

func stringify(value interface{}) string {
	if value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", value)
}

func (e Expression) ToValue() (interface{}, error) {
	_, err := evaluate(e.expression)
	return nil, err
}

func (p Print) ToValue() (interface{}, error) {
	value, err := evaluate(p.expression)
	if err != nil {
		return nil, err
	}
	fmt.Println(stringify(value))
	return nil, nil
}

func (l Literal) ToValue() (interface{}, error) {
	return l.value, nil
}

func (g Grouping) ToValue() (interface{}, error) {
	return evaluate(g.expression)
}

func (b Binary) ToValue() (interface{}, error) {
	left, err := evaluate(b.left)
	right, err := evaluate(b.right)
	if err != nil {
		return nil, err
	}

	switch b.operation.TokenType {
	case token.GreaterToken:
		err := checkNumberOperands(b.operation, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) > right.(float64), nil
	case token.GreaterEqualToken:
		err := checkNumberOperands(b.operation, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) >= right.(float64), nil
	case token.LessToken:
		err := checkNumberOperands(b.operation, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) < right.(float64), nil
	case token.LessEqualToken:
		err := checkNumberOperands(b.operation, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) <= right.(float64), nil
	case token.PlusToken:
		fLeft, lOk := left.(float64)
		fRight, rOk := right.(float64)

		if lOk && rOk {
			return fLeft + fRight, nil
		}

		sLeft, lOk := left.(string)
		sRight, rOk := right.(string)

		if lOk && rOk {
			return sLeft + sRight, nil
		}

		return nil, NewRuntimeError(b.operation, "Operands must be two numbers or two strings.")
	case token.MinusToken:
		err := checkNumberOperands(b.operation, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) - right.(float64), nil
	case token.SlashToken:
		err := checkNumberOperands(b.operation, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) / right.(float64), nil
	case token.StarToken:
		err := checkNumberOperands(b.operation, left, right)
		if err != nil {
			return nil, err
		}
		return left.(float64) * right.(float64), nil
	case token.BangEqualToken:
		return !isEqual(left, right), nil
	case token.EqualEqualToken:
		return isEqual(left, right), nil
	}

	return nil, NewRuntimeError(b.operation, "Undefined operator.")
}

func (u Unary) ToValue() (interface{}, error) {
	right, err := evaluate(u.right)
	if err != nil {
		return nil, err
	}

	switch u.operation.TokenType {
	case token.BangToken:
		return !isTruthy(right), nil
	case token.MinusToken:
		err := checkNumberOperand(u.operation, right)
		if err != nil {
			return nil, err
		}
		return -(right.(float64)), nil
	}

	return nil, NewRuntimeError(u.operation, "Undefined operator.")
}

func evaluate(e Expr) (interface{}, error) {
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

// Runtime checks

func checkNumberOperand(operator token.Token, operand interface{}) error {
	if _, ok := operand.(float64); ok {
		return nil
	}

	return NewRuntimeError(operator, "Operand must be a number.")
}

func checkNumberOperands(operator token.Token, left, right interface{}) error {
	_, lOk := left.(float64)
	_, rOk := right.(float64)

	if lOk && rOk {
		return nil
	}

	return NewRuntimeError(operator, "Operands must be a number.")
}
